package aset

import (
	"encoding/json"
	"gitee.com/asktop_golib/util/acast"
	"gitee.com/asktop_golib/util/async"
	"reflect"
	"strings"
)

type OrderSet struct {
	mu  *async.RWMutex
	arr []interface{}
	m   map[interface{}]struct{}
}

// New create and returns a new set, which contains un-repeated items.
// The param <unsafe> used to specify whether using set in un-concurrent-safety,
// which is false in default.
func NewOrderSet(safe ...bool) *OrderSet {
	return &OrderSet{
		arr: []interface{}{},
		m:   make(map[interface{}]struct{}),
		mu:  async.New(safe...),
	}
}

// NewOrderSetFrom returns a new set from <items>.
// Parameter <items> can be either a variable of any type, or a slice.
func NewOrderSetFrom(items interface{}, safe ...bool) *OrderSet {
	m := make(map[interface{}]struct{})
	arr := []interface{}{}

	arrVal := reflect.ValueOf(items)
	if arrVal.Kind() == reflect.Slice {
		l := arrVal.Len()
		for i := 0; i < l; i++ {
			v := arrVal.Index(i).Interface()
			if _, ok := m[v]; !ok {
				m[v] = struct{}{}
				arr = append(arr, v)
			}
		}
	} else {
		m[items] = struct{}{}
		arr = append(arr, items)
	}

	return &OrderSet{
		arr: arr,
		m:   m,
		mu:  async.New(safe...),
	}
}

func (s *OrderSet) IsSafe() bool {
	return s.mu.IsSafe()
}

func (s *OrderSet) SetSafe(safe bool) {
	if safe == s.mu.IsSafe() {
		return
	}
	s.mu = async.New(safe)
}

func (s *OrderSet) Clone() *OrderSet {
	return NewOrderSetFrom(s.Slice(), s.mu.IsSafe())
}

// Clear deletes all items of the set.
func (s *OrderSet) Clear() *OrderSet {
	s.mu.Lock()
	s.arr = []interface{}{}
	s.m = make(map[interface{}]struct{})
	s.mu.Unlock()
	return s
}

// Slice returns the a of items of the set as slice.
func (s *OrderSet) Slice() []interface{} {
	return s.arr
}

func (s *OrderSet) SliceString() []string {
	arr := []string{}
	for _, v := range s.arr {
		arr = append(arr, acast.ToString(v))
	}
	return arr
}

func (s *OrderSet) SliceInt() []int {
	arr := []int{}
	for _, v := range s.arr {
		arr = append(arr, acast.ToInt(v))
	}
	return arr
}

func (s *OrderSet) SliceInt64() []int64 {
	arr := []int64{}
	for _, v := range s.arr {
		arr = append(arr, acast.ToInt64(v))
	}
	return arr
}

// Add adds one or multiple items to the set.
// 赋值（若值已存在，则替换新顺序）
func (s *OrderSet) Add(item ...interface{}) *OrderSet {
	if len(item) == 0 {
		return s
	}
	s.mu.Lock()
	newArr := []interface{}{}
	newMap := map[interface{}]struct{}{}
	itemMap := map[interface{}]struct{}{}
	for _, v := range item {
		itemMap[v] = struct{}{}
	}
	for _, v := range s.arr {
		if _, ok := itemMap[v]; !ok {
			newArr = append(newArr, v)
			newMap[v] = struct{}{}
		}
	}
	for _, v := range item {
		if _, ok := newMap[v]; !ok {
			newMap[v] = struct{}{}
			newArr = append(newArr, v)
		}
	}
	s.arr = newArr
	s.m = newMap
	s.mu.Unlock()
	return s
}

//赋值（若值已存在，则不替换）
func (s *OrderSet) Set(item ...interface{}) *OrderSet {
	if len(item) == 0 {
		return s
	}
	s.mu.Lock()
	for _, v := range item {
		if _, ok := s.m[v]; !ok {
			s.m[v] = struct{}{}
			s.arr = append(s.arr, v)
		}
	}
	s.mu.Unlock()
	return s
}

// Remove deletes <item> from set.
func (s *OrderSet) Remove(item ...interface{}) *OrderSet {
	if len(item) == 0 {
		return s
	}
	s.mu.Lock()
	newArr := []interface{}{}
	newMap := map[interface{}]struct{}{}
	itemMap := map[interface{}]struct{}{}
	for _, v := range item {
		itemMap[v] = struct{}{}
	}
	for _, v := range s.arr {
		if _, ok := itemMap[v]; !ok {
			newArr = append(newArr, v)
			newMap[v] = struct{}{}
		}
	}
	s.arr = newArr
	s.m = newMap
	s.mu.Unlock()
	return s
}

// Contains checks whether the set contains <item>.
func (s *OrderSet) Contains(item interface{}) bool {
	s.mu.RLock()
	_, exists := s.m[item]
	s.mu.RUnlock()
	return exists
}

// Size returns the size of the set.
func (s *OrderSet) Size() int {
	s.mu.RLock()
	l := len(s.arr)
	s.mu.RUnlock()
	return l
}

func (s *OrderSet) IsEmpty() bool {
	return s.Size() == 0
}

// Iterator iterates the set with given callback function <f>,
// if <f> returns true then continue iterating; or false to stop.
func (s *OrderSet) Iterator(f func(v interface{}) bool) *OrderSet {
	s.mu.RLock()
	defer s.mu.RUnlock()
	for _, v := range s.arr {
		if !f(v) {
			break
		}
	}
	return s
}

// LockFunc locks writing with callback function <f>.
func (s *OrderSet) LockFunc(f func(arr []interface{}, m map[interface{}]struct{})) {
	s.mu.Lock()
	defer s.mu.Unlock()
	f(s.arr, s.m)
}

// RLockFunc locks reading with callback function <f>.
func (s *OrderSet) RLockFunc(f func(arr []interface{}, m map[interface{}]struct{})) {
	s.mu.RLock()
	defer s.mu.RUnlock()
	f(s.arr, s.m)
}

// Equal checks whether the two sets equal.
func (s *OrderSet) Equal(other *OrderSet, force ...bool) bool {
	if s == other {
		return true
	}
	s.mu.RLock()
	defer s.mu.RUnlock()
	other.mu.RLock()
	defer other.mu.RUnlock()
	if len(s.arr) != len(other.arr) {
		return false
	}
	if len(force) > 0 && force[0] {
		for i := 0; i < len(s.arr); i++ {
			if s.arr[i] != other.arr[i] {
				return false
			}
		}
	} else {
		for key := range s.m {
			if _, ok := other.m[key]; !ok {
				return false
			}
		}
	}
	return true
}

// IsSubsetOf checks whether the current set is a sub-set of <other>.
func (s *OrderSet) IsSubsetOf(other *OrderSet) bool {
	if s == other {
		return true
	}
	s.mu.RLock()
	defer s.mu.RUnlock()
	other.mu.RLock()
	defer other.mu.RUnlock()
	for key := range s.m {
		if _, ok := other.m[key]; !ok {
			return false
		}
	}
	return true
}

// Merge adds items from <others> sets into <set>.
func (s *OrderSet) Merge(others ...*OrderSet) *OrderSet {
	s.mu.Lock()
	defer s.mu.Unlock()
	for _, other := range others {
		if s != other {
			other.mu.RLock()
		}
		for _, v := range other.arr {
			if _, ok := s.m[v]; !ok {
				s.m[v] = struct{}{}
				s.arr = append(s.arr, v)
			}
		}
		if s != other {
			other.mu.RUnlock()
		}
	}
	return s
}

// Union returns a new set which is the union of <set> and <others>.
// Which means, all the items in <newSet> are in <set> or in <others>.
func (s *OrderSet) Union(others ...*OrderSet) (newSet *OrderSet) {
	newSet = NewOrderSet(s.mu.IsSafe())
	s.mu.RLock()
	defer s.mu.RUnlock()
	for _, other := range others {
		if s != other {
			other.mu.RLock()
		}
		for _, v := range other.arr {
			newSet.m[v] = struct{}{}
			newSet.arr = append(newSet.arr, v)
		}
		if s != other {
			for _, v := range other.arr {
				if _, ok := newSet.m[v]; !ok {
					newSet.m[v] = struct{}{}
					newSet.arr = append(newSet.arr, v)
				}
			}
		}
		if s != other {
			other.mu.RUnlock()
		}
	}
	return newSet
}

// Diff returns a new set which is the difference set from <set> to <others>.
// Which means, all the items in <newSet> are in <set> but not in <others>.
func (s *OrderSet) Diff(others ...*OrderSet) (newSet *OrderSet) {
	newSet = NewOrderSet(s.mu.IsSafe())
	s.mu.RLock()
	defer s.mu.RUnlock()
	for _, other := range others {
		if s == other {
			continue
		}
		other.mu.RLock()
		for _, v := range s.arr {
			if _, ok := other.m[v]; !ok {
				newSet.m[v] = struct{}{}
				newSet.arr = append(newSet.arr, v)
			}
		}
		other.mu.RUnlock()
	}
	return newSet
}

// Intersect returns a new set which is the intersection from <set> to <others>.
// Which means, all the items in <newSet> are in <set> and also in <others>.
func (s *OrderSet) Intersect(others ...*OrderSet) (newSet *OrderSet) {
	newSet = NewOrderSet(s.mu.IsSafe())
	s.mu.RLock()
	defer s.mu.RUnlock()
	for _, other := range others {
		if s != other {
			other.mu.RLock()
		}
		for _, v := range s.arr {
			if _, ok := other.m[v]; ok {
				newSet.m[v] = struct{}{}
				newSet.arr = append(newSet.arr, v)
			}
		}
		if s != other {
			other.mu.RUnlock()
		}
	}
	return
}

// Complement returns a new set which is the complement from <set> to <full>.
// Which means, all the items in <newSet> are in <full> and not in <set>.
//
// It returns the difference between <full> and <set>
// if the given set <full> is not the full set of <set>.
func (s *OrderSet) Complement(full *OrderSet) (newSet *OrderSet) {
	newSet = NewOrderSet(s.mu.IsSafe())
	s.mu.RLock()
	defer s.mu.RUnlock()
	if s != full {
		full.mu.RLock()
		defer full.mu.RUnlock()
	}
	for _, v := range full.arr {
		if _, ok := s.m[v]; !ok {
			newSet.m[v] = struct{}{}
			newSet.arr = append(newSet.arr, v)
		}
	}
	return
}

// Sum sums items.
// Note: The items should be converted to int type,
// or you'd get a result that you unexpected.
func (s *OrderSet) Sum() (sum int) {
	s.mu.RLock()
	defer s.mu.RUnlock()
	for _, v := range s.arr {
		sum += acast.ToInt(v)
	}
	return
}

// Join joins items with a string <sep>.
func (s *OrderSet) Join(sep string) string {
	return strings.Join(acast.ToStringSlice(s.Slice()), sep)
}

func (s *OrderSet) MarshalJSON() ([]byte, error) {
	return json.Marshal(s.Slice())
}

func (s *OrderSet) UnmarshalJSON(b []byte) error {
	if s.mu == nil {
		s.mu = async.New()
	}
	var data []interface{}
	err := json.Unmarshal(b, &data)
	if err != nil {
		return err
	} else {
		s.Add(data...)
		return nil
	}
}

func (s *OrderSet) String() string {
	rs, _ := s.MarshalJSON()
	return string(rs)
}
