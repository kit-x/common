package stringx

import (
	"testing"
)

func TestRandString(t *testing.T) {
	str := RandomString(16)
	if len(str) != 16 {
		t.Error("random string with wrong len")
	}
}
