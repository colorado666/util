package example

import (
	"fmt"
	"gitee.com/asktop_golib/util/afile"
	"path"
	"path/filepath"
	"testing"
)

func TestGetNames(t *testing.T) {
	fp, _ := afile.GetWorkPath()
	fmt.Println(fp)
	fmt.Println(path.Split(fp))
	fmt.Println(filepath.Base(fp))
}
