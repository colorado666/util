package ajwt

import (
	"encoding/json"
	"github.com/dgrijalva/jwt-go"
)

// Claims type that uses the map[string]interface{} for JSON decoding
// This is the default claims type if you don't supply one
type MapClaims map[string]interface{}

// Compares the exp claim against cmp.
// If required is false, this method will return true if the value matches or is unset
func (m MapClaims) VerifyExpiresAt(cmp int64, req bool) bool {
	switch exp := m["exp"].(type) {
	case float64:
		return verifyExp(int64(exp), cmp, req)
	case json.Number:
		v, _ := exp.Int64()
		return verifyExp(v, cmp, req)
	}
	return req == false
}

// Compares the iat claim against cmp.
// If required is false, this method will return true if the value matches or is unset
func (m MapClaims) VerifyIssuedAt(cmp int64, req bool) bool {
	switch iat := m["iat"].(type) {
	case float64:
		return verifyIat(int64(iat), cmp, req)
	case json.Number:
		v, _ := iat.Int64()
		return verifyIat(v, cmp, req)
	}
	return req == false
}

// Compares the nbf claim against cmp.
// If required is false, this method will return true if the value matches or is unset
func (m MapClaims) VerifyNotBefore(cmp int64, req bool) bool {
	switch nbf := m["nbf"].(type) {
	case float64:
		return verifyNbf(int64(nbf), cmp, req)
	case json.Number:
		v, _ := nbf.Int64()
		return verifyNbf(v, cmp, req)
	}
	return req == false
}

// Validates time based claims "exp, iat, nbf".
// There is no accounting for clock skew.
// As well, if any of the above claims are not in the token, it will still
// be considered a valid claim.
func (m MapClaims) Valid() error {
	vErr := new(jwt.ValidationError)
	//now := jwt.TimeFunc().Unix()

	//if m.VerifyExpiresAt(now, false) == false {
	//	vErr.Inner = errors.New("Token is expired")
	//	vErr.Errors |= jwt.ValidationErrorExpired
	//}

	//if m.VerifyIssuedAt(now, false) == false {
	//	vErr.Inner = errors.New("Token used before issued")
	//	vErr.Errors |= jwt.ValidationErrorIssuedAt
	//}

	//if m.VerifyNotBefore(now, false) == false {
	//	vErr.Inner = errors.New("Token is not valid yet")
	//	vErr.Errors |= jwt.ValidationErrorNotValidYet
	//}

	if vErr.Errors == 0 {
		return nil
	}

	return vErr
}

func verifyExp(exp int64, now int64, required bool) bool {
	if exp == 0 {
		return !required
	}
	return now <= exp
}

func verifyIat(iat int64, now int64, required bool) bool {
	if iat == 0 {
		return !required
	}
	return now >= iat
}

func verifyNbf(nbf int64, now int64, required bool) bool {
	if nbf == 0 {
		return !required
	}
	return now >= nbf
}
