package main

import (
	"fmt"
	"regexp"
	"strings"
)

func main() {
	txt := "hussain ((((((up))))))) hello (((something)))"

	fmt.Println(ProcessRealCommands(txt))
}

func ProcessRealCommands(txt string) string {
	validCommands := map[string]bool{
		"up":  true,
		"low": true,
		"cap": true,
		"hex": true,
		"bin": true,
	}

	for {
		changed := false
		
		re := regexp.MustCompile(`\(+\s*([a-zA-Z]+)\s*\)+`)
		txt = re.ReplaceAllStringFunc(txt, func(match string) string {
			
			inner := strings.Trim(match, "() ")
			if validCommands[inner] {
				changed = true
				return "" 
			}
			return match 
		})

		if !changed {
			break
		}
	}

	return strings.TrimSpace(txt)
}
