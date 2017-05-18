package logrus

import (
	"strings"
)

//
// Fields define list of field in message attachments.
//
type Fields []Field

const (
	_emptyField = `{}`
	_sep        = `,`
)

//
// MarshalJSON will convert `field` into a valid JSON. We use manual
// convertion for gaining speed.
//
func (fields Fields) MarshalJSON() (out []byte, err error) {
	var fout []byte
	str := "["

	for _, field := range fields {
		fout, err = field.MarshalJSON()
		if err != nil {
			return
		}

		if string(fout) != _emptyField {
			str += string(fout) + _sep
		}
	}

	str = strings.TrimRight(str, _sep)

	str += "]"
	out = []byte(str)

	return
}
