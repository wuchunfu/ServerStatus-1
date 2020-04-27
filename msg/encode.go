package msg

import (
    "bytes"
    "io"
)

func Write(typeMessage uint8, extra ...interface{}) []byte {
    buf := bytes.NewBuffer(make([]byte, 0))
    buf.WriteByte(typeMessage)
    for _, val := range extra {
        switch val.(type) {
        case string:
            buf.WriteString(val.(string))
        case rune:
            buf.WriteRune(val.(rune))
        case byte:
            buf.WriteByte(val.(byte))
        case int:
            buf.WriteByte(byte(val.(int)))
        case []byte:
            buf.Write(val.([]byte))
        case io.Writer:
            buf.WriteTo(val.(io.Writer))
        }
    }

    return buf.Bytes()
}
