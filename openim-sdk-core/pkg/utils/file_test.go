package utils

import "testing"

func Test_FileTmpPath(t *testing.T) {
	s := FileTmpPath("", "./")
	t.Log(s)

}
