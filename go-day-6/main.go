package main

import (
	"bufio"
	"fmt"
	"strings"
)

func main() {
	s := "lorem hasfej jerhı ehjkh ek ıekkk ıle \nkkekhfaknıqüajfa şher öşafdjf owği malelu liaenro"
	scanner := bufio.NewScanner(strings.NewReader(s))
	scanner.Split(bufio.ScanRunes)
	for scanner.Scan() {
		line := scanner.Text()
		fmt.Println(line)
	}

}
