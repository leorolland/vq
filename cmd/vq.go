package main

import (
	"bufio"
	"fmt"
	"os"

	"github.com/leorolland/vq/parser"
	"github.com/leorolland/vq/parsers"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		things, err := parser.Parse(parsers.Anythings(), scanner.Text())
		if err != nil {
			panic(err)
		}

		fmt.Println(things)
	}

	if scanner.Err() != nil {
		panic(scanner.Err().Error())
	}
}
