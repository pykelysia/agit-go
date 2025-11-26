package decoder

import (
	"testing"

	"github.com/pykelysia/pyketools"
)

func TestDecoder(t *testing.T) {
	original := []byte{0xC4, 0xE3, 0xBA, 0xC3} // "你好" in GBK encoding
	decoded, err := NewDecoder(original)
	if err != nil {
		pyketools.Errorf("Decoding failed: %v", err)
		return
	}

	pyketools.Infof("decoded string: %s", decoded)
}
