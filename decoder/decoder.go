package decoder

import (
	"io"
	"strings"

	"golang.org/x/text/encoding/simplifiedchinese"
	"golang.org/x/text/transform"
)

func NewDecoder(p []byte) (string, error) {
	d := &decoder{
		p: p,
	}
	return d.decoder()
}

func (d *decoder) decoder() (string, error) {
	decoder := simplifiedchinese.GBK.NewDecoder()
	decodedP, err := io.ReadAll(transform.NewReader(
		strings.NewReader(string(d.p)),
		decoder,
	))
	if err != nil {
		return "", err
	}
	d.t = string(decodedP)
	return d.t, nil
}
