// 【坐在马桶上看算法】算法4：队列——解密QQ号
// http://blog.51cto.com/ahalei/1371613

package main

import (
	"fmt"
	"strings"
)

func main()  {

	encryptedMsg := "631758924"
	//fmt.Println("hello queue", encryptedMsg)

	fmt.Println(decrypt(encryptedMsg))

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