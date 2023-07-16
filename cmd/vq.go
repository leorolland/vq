package main

import (
	"bufio"
	"fmt"
	"os"

	"github.com/leorolland/vq/parser"
	vq "github.com/leorolland/vq/pkg"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		things, err := parser.Parse(vq.Anythings(), scanner.Text())
		if err != nil {
			panic(err)
		}

		fmt.Println(things)
	}

	if scanner.Err() != nil {
		panic(scanner.Err().Error())
	}
}
