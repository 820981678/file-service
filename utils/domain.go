package utils

import (
	"fmt"
	"os"
	"bufio"
)

var ApplicationDomain []string

func Init_domain() {
	fmt.Println("load domain start......")
	
	ApplicationDomain = make([]string, 0, 10)
	file, err := os.Open("domain.properties")
	if err != nil {
		fmt.Println("file domain.properties error default png jpg gif jpeg bmp")
		return
	}
	defer file.Close()
	
	reader := bufio.NewReader(file)
	for {
		line, _, err := reader.ReadLine()
		if err != nil {
			break
		}
		ApplicationDomain = append(ApplicationDomain, string(line))
	}
	
	fmt.Printf("load domain complete ApplicationDomain: %v \n", ApplicationDomain)
}