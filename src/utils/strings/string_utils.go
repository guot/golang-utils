// StringUtils
package sutils

import (
	"bufio"
	"bytes"
	"fmt"
	"reflect"
	"strconv"
)

const endLine = "\r\n"

func ValueOf(v interface{}) (str string) {

	val := reflect.ValueOf(v)
	switch val.Kind() {
	case reflect.Bool:
		str = strconv.FormatBool(val.Bool())
	case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
		str = strconv.Itoa(int(val.Int()))
	case reflect.Float32, reflect.Float64:

		str = strconv.FormatFloat(val.Float(), 'f', -1, 32)
	case reflect.String:
		str = val.String()
	default:
		str = ""
		fmt.Println("unknown type!")

	}
	return str
}
func AppendToNewLine(buf *bytes.Buffer, val interface{}) {
	//	 strconv.AppendInt(buf,val.(int),10)
	wr := bufio.NewWriter(buf)
	fmt.Fprintf(wr, "%d\r\n", val.(int))
	wr.Flush()
	//	str = buf.String()
	//	buf.Write([]byte(ValueOf(val)))
	//	buf.WriteString(endLine)
}
