package chars

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

const staticStr = "stringx"

// "stringx"
var staticBytes = []uint8{0x73, 0x74, 0x72, 0x69, 0x6e, 0x67, 0x78,}

func TestBytesToString(t *testing.T) {
	assert.Equal(t, BytesToString(staticBytes), staticStr)
}

func TestStringToBytes(t *testing.T) {
	assert.Equal(t, StringToBytes(staticStr), staticBytes)
}

func BenchmarkBytesToString(b *testing.B) {
	for i := 0; i < b.N; i++ {
		BytesToString(staticBytes)
	}
}

func BenchmarkCommonBytesToString(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_ = string(staticBytes)
	}
}

func BenchmarkStringToBytes(b *testing.B) {
	StringToBytes(staticStr)
}

func BenchmarkCommonStringToBytes(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_ = []byte(staticStr)
	}
}
