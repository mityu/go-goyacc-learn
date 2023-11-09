package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	for {
		fmt.Print("> ")
		scanner.Scan()
		l := &Lexer{}
		l.Init(strings.NewReader(scanner.Text()))
		yyParse(l)
		v, err := l.result.Eval()
		if err != nil {
			log.Fatal(err)
		}
		fmt.Println(v.(ExprNum).literal)
	}
}
