package stringx

import (
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

var (
	str   = strings.Repeat("s", 1024)
	bytes = []byte(str)
)

func TestBytesToString(t *testing.T) {
	if BytesToString(bytes) != str {
		t.Errorf("BytesToString should equal %s", str)
	}
}

func TestStringToBytes(t *testing.T) {
	assert.Equal(t, bytes, StringToBytes(str))
}
