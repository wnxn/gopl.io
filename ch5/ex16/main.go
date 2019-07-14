package main

func main() {

}

func join1(strs ...string) string {
	res := ""
	for _, v := range strs {
		res += v
	}
	return res
}

func join2(strs ...string) string {
	res := []byte{}
	for _, v := range strs {
		res = append(res, []byte(v)...)
	}
	return string(res)
}

func join3(strs ...string) string {
	length := 0
	for _, v := range strs {
		length += len(v)
	}
	res := make([]byte, length)
	head := 0
	for _, v := range strs {
		copy(res[head:], []byte(v))
		head = head + len(v)
	}
	return string(res)
}
