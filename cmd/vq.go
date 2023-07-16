package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"os"

	"github.com/leorolland/vq/parser"
	"github.com/leorolland/vq/parsers"
)

const MAX_RECURSION_LEVEL = 100

func main() {
	scanner := bufio.NewScanner(os.Stdin)
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

		fmt.Println(string(jsonOutput))
	}

	if scanner.Err() != nil {
		panic(scanner.Err().Error())
	}
}
