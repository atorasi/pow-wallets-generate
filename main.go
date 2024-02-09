package main

import (
	"fmt"
	"os"
	"strings"
	"unicode"

	sol "github.com/gagliardetto/solana-go"
)

func main() {
	for index := 0; ; index++ {
		x := sol.NewWallet()
		fmt.Printf("%v Address: %v | %v Private: %v | HUESOSI PLEASE GIVE ME SOL\n", index, x.PublicKey(), index, x.PrivateKey)

		if x.PublicKey().String()[:3] == "pow" && unicode.IsDigit(rune(x.PublicKey().String()[3])) {
			fileWrite(x.PublicKey().String()[:10], []byte(x.PrivateKey))
		}
	}
}

func fileWrite(filename string, data []byte) {
	file, _ := os.Create(filename + ".json")
	defer file.Close()

	str := makeStr(data)

	_, err := file.WriteString(str.String())
	if err != nil {
		fmt.Println("Ошибка при записи в файл:", err)
		return
	}

}

func makeStr(data []byte) strings.Builder {
	var str strings.Builder
	str.WriteString("[")
	for i, num := range data {
		if i != 0 {
			str.WriteString(",")
		}
		str.WriteString(fmt.Sprintf("%d", num))
	}
	str.WriteString("]")

	return str
}
