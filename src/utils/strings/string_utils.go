// StringUtils
package sutils

import (
	"bytes"
	"fmt"
	"strconv"
)

func ValueOf(val interface{}) (str string) {
	switch  val.(type) {
	case bool:
		str = strconv.FormatBool(val.(bool))
	case int:
		str = strconv.Itoa(val.(int))
	case float32:

		str = strconv.FormatFloat(float64(val.(float32)), 'f', -1, 32)
	case string:
		str = val.(string)
	default:
		str = ""
		fmt.Println("unknown type!")

	}
	return str
}
func AppendToNewLine(buf *bytes.Buffer, val interface{}) {
	buf.WriteString(ValueOf(val))
	buf.WriteString("\r\n")
}
