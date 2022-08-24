package Parser

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"strings"
)

func ReadFile(input string) []string {
	file, err := os.Open(input)

	if err != nil {
		log.Fatal("open file error: ", err)
	}
	defer file.Close()

	data, err := ioutil.ReadAll(file)
	if err != nil {
		log.Fatal("read file error: ", err)
	}
	dataSep := strings.Split(string(data), "\n")
	return dataSep
}

func WriteFile(data []string, output string) {
	if len(output) != 0 {
		file, err := os.Create(output)
		if err != nil {
			log.Fatal("create file error: ", err)
		}
		defer file.Close()

		for _, line := range data {
			if _, err = file.WriteString(line + "\n"); err != nil {
				log.Fatal("write file error: ", err)
			}
		}
	} else {
		fmt.Println(data)
	}
}
