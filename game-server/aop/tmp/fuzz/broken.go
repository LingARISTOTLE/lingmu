package fuzz

/*
BrokenMethod
@Description: 有一个错误 - 即使它的长度只有 3，它也会尝试读取数据的第 4 个索引。
@param Data
@return bool
*/
func BrokenMethod(Data string) bool {
	return len(Data) >= 3 &&
		Data[0] == 'F' &&
		Data[1] == 'U' &&
		Data[2] == 'Z' &&
		Data[3] == 'Z'
}

/*
Reverse
@Description: 反转字符串
@param s
@return string
*/
func Reverse(s string) string {
	bs := []byte(s)
	length := len(bs)
	for i := 0; i < length/2; i++ {
		bs[i], bs[length-i-1] = bs[length-i-1], bs[i]
	}
	return string(bs)
}
