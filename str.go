package go_utils

import "bytes"

func JoinStrings(slist ...string) string {
	//定义 一个字节缓冲，快速连接字符串
	var b bytes.Buffer
	for _, s := range slist {
		//将遍历的字节字符串写入字节数组
		b.WriteString(s)

	}
	//将连接好的字节数组转为字符串
	return b.String()
}
