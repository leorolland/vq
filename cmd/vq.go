package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"io"
	"os"

	"github.com/leorolland/vq/parser"
	"github.com/leorolland/vq/parsers"
)

const MAX_RECURSION_LEVEL = 100

func main() {
	end := make(chan bool)
	go parseAndWrite(os.Stdin, os.Stdout, end)
	<-end
}

func parseAndWrite(reader io.Reader, writer io.Writer, end chan bool) {
	scanner := bufio.NewScanner(reader)
	for scanner.Scan() {
		things, err := parser.Parse(
			parsers.Anythings(MAX_RECURSION_LEVEL),
			scanner.Text(),
		)
		if err != nil {
			panic(err)
		}

		jsonOutput, err := json.MarshalIndent(things, "", "  ")
		if err != nil {
			panic(err)
		}

		fmt.Fprint(writer, string(jsonOutput))
	}

	if scanner.Err() != nil {
		panic(scanner.Err().Error())
	}

	end <- true
}
