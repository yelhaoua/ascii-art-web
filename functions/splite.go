package asciiweb

import (
	"fmt"
	"os"
	"strings"
)

func Spite(input, style string) string {
	out := ""
	data, err := os.ReadFile("../styles/"+style+".txt")
	if err != nil {
		return "error type 01"
	}
	file := string(data)
	var final [][]string
	file = strings.ReplaceAll(file, "\r", "")
	Start := strings.Split(file, "\n")
	str := strings.ReplaceAll(input, "\\n", "\n")
	if strings.Trim(str, "\n") == "" {
		for i := 0; i < len(input); i++ {
			if input[i] == '\\' {
				fmt.Println()
			}
		}
		return  ""
	}
	splited := strings.Split(str, "\n")

	for _, va := range splited {
		for _, v := range va {
			index := int(((v - 32) * 9) + 1)
			final = append(final, Start[index:index+8])
		}
		if len(va) != 0 {
			
			for i := 0; i < 8; i++ {
				for a := range final {
					out += final[a][i]
				}
				if i != 7 {
					out += "\n"
				}
			}
			return  out
		} else {
			return ""
		}
	}
	return  out
}
