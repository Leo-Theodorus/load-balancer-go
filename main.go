package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	sc := bufio.NewScanner(os.Stdin)
	sc.Buffer(make([]byte, 1024*1024), 1024*1024)
	type Command string

	const (
		PICK Command =  "PICK"
		POOL Command = "POOL"
		RESET Command =  "RESET"
	)

	var servers []string

	currentIndex := 0

	for sc.Scan() {
		line := sc.Text()
		if line == "" { continue }
		args := strings.Split(line, " ")

		argCount := len(args)
		if argCount == 0 { continue }

		command := Command(args[0])
		switch command {
		case PICK :
			serverCount := len(servers)
			if serverCount == 0 {
				fmt.Println("EMPTY")
				break
			}

			if currentIndex >= serverCount {
				panic("INDEX OOB")
			}
			fmt.Println(servers[currentIndex])

			currentIndex += 1
			if currentIndex == serverCount {
				currentIndex = 0
			}

		case POOL:
			servers = nil
			for index, command := range args {
				if index == 0 { continue }
				servers = append(servers, command)
			}
			fmt.Println("OK")
			
		case RESET:
			fmt.Println("OK")
			currentIndex = 0
		}
		
	}
}
