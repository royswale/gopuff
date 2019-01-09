// 【坐在马桶上看算法】算法4：队列——解密QQ号
// http://blog.51cto.com/ahalei/1371613

package main

import (
	"fmt"
	"strings"
)

type queue struct {
	data *[]string
	head int
	tail int
}

func main()  {

	encryptedMsg := "631758924"
	//fmt.Println("hello queue", encryptedMsg)

	//fmt.Println(decrypt(encryptedMsg))

	//fmt.Println(decryptUsePointer(encryptedMsg))

	decryptUseStruct(encryptedMsg)
}

// 浪费时间 不需要 s = s[1:] 这种操作
func decrypt(encryptedMsg string) string {

	//fmt.Println(encryptedMsg)

	s := make([]string, 0)
	for _, value := range encryptedMsg {
		s = append(s, string(value))
	}

	//fmt.Println(s)

	x := make([]string, 0)
	var shift, move string

	for {

		shift, s = s[0], s[1:]
		x = append(x, shift)

		//fmt.Println(shift)
		//fmt.Println(s)
		//fmt.Println(x)

		if len(s) == 0 {
			break
		}

		move, s = s[0], s[1:]
		s = append(s, move)

		//fmt.Println(move)
		//fmt.Println(s)

		//fmt.Println(strings.Join(s,""))
		fmt.Println(strings.Join(x,""))
	}

	//fmt.Println(s)
	//fmt.Println(strings.Join(s,""))

	//fmt.Println(x)
	//fmt.Println(strings.Join(x,""))

	return strings.Join(x, "")
}

func decryptUsePointer(encryptedMsg string) string {

	//x := make([]rune, 0)
	x := make([]string, 0)

	for _, value := range encryptedMsg {
		//x = append(x, value)
		x = append(x, string(value))
		//fmt.Println(value)
		//fmt.Println(string(value))
	}

	fmt.Println(x)

	var head = 0
	var tail = len(x)
	//fmt.Println(head, tail)

	s := make([]string, 0)

	for head < tail {

		s = append(s, x[head])
		//s = append(s, string(x[head]))
		fmt.Println(s)
		head++

		if head == tail {
			break
		}

		x = append(x, x[head])
		//x[tail] = x[head]
		head++
		tail++

		fmt.Println(x[head:tail])
	}

	//return encryptedMsg
	return strings.Join(s, "")
}

func decryptUseStruct(encryptedMsg string) string {

	x := make([]string, 0)
	for _, value := range encryptedMsg {
		x = append(x, string(value))
	}

	s := &queue{
		data: &x,
		head: 0,
		tail: len(x),
	}

	for s.head < s.tail {
		fmt.Print((*(s.data))[s.head])
		s.head++

		if s.head == s.tail {
			break
		}

		*(s.data) = append(*(s.data), (*(s.data))[s.head])
		s.head++
		s.tail++
	}

	//fmt.Println(s.data)

	return encryptedMsg
}