package go_utils

import (
	"bytes"
	"encoding/binary"
	"strconv"
)

// bytes to int 32
func BytesTo32Int(b []byte) int {
	buf := bytes.NewBuffer(b)
	var tmp uint32
	binary.Read(buf, binary.BigEndian, &tmp)
	return int(tmp)
}

// bytes to int 16
func BytesTo16Int(b []byte) int {
	buf := bytes.NewBuffer(b)
	var tmp uint16
	binary.Read(buf, binary.BigEndian, &tmp)
	return int(tmp)
}

// int to 4 bytes
func IntTo4Bytes(i int) []byte {
	buf := bytes.NewBuffer([]byte{})
	tmp := uint32(i)
	binary.Write(buf, binary.BigEndian, tmp)
	return buf.Bytes()
}

// int to 2 bytes
func IntTo2Bytes(i int) []byte {
	buf := bytes.NewBuffer([]byte{})
	tmp := uint16(i)
	binary.Write(buf, binary.BigEndian, tmp)
	return buf.Bytes()
}

// bytes to hex string
func BytesToHexString(b []byte) string {
	var buf bytes.Buffer
	for _, v := range b {
		t := strconv.FormatInt(int64(v), 16)
		if len(t) > 1 {
			buf.WriteString(t)
		} else {
			buf.WriteString("0" + t)
		}
	}
	return buf.String()
}

// hex string to bytes
func HexStringToBytes(s string) []byte {
	bs := make([]byte, 0)
	for i := 0; i < len(s); i = i + 2 {
		b, _ := strconv.ParseInt(s[i:i+2], 16, 16)
		bs = append(bs, byte(b))
	}
	return bs
}

// 二进制字符串转byte
// a, err := strconv.ParseInt("11111111", 2, 16)
