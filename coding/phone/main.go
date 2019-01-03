package main

import (
	"fmt"
	"strings"
)

func main() {
	phone := "13424300460"
	bytes := []byte(phone)
	retPhone := []byte{}
	for i := 0; i < len(bytes); i++ {
		if i >= 3 && i <= 6 {
			retPhone = append(retPhone, '*')
		} else {
			retPhone = append(retPhone, bytes[i])
		}
	}

	slice := strings.Split(phone, "")
	retPhone := strings.Join(slice[0:3], "") + "****" + strings.Join(slice[7:], "")

	fmt.Println(string(retPhone))
	//for i := 0; i < len(str); i++ {
	//
	//}
}
