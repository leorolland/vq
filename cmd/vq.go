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
	count := 0
	for scanner.Scan() {
		count += 1

		lineParser := vq.NewLineParser()

		line, err := parser.Parse(lineParser.LineParser, scanner.Text())
		if err != nil {
			panic(err)
		}

		fmt.Printf("%d: line: %s\n", count, line)
	}

	if scanner.Err() != nil {
		panic(scanner.Err().Error())
	}
}
