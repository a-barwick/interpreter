package repl

import (
	"bufio"
	"fmt"
	"io"

	"github.com/a-barwick/interpreter/lexer"
	"github.com/a-barwick/interpreter/token"
)

const PROMPT string = ">> "

func Start(in io.Reader, out io.Writer) {
	scanner := bufio.NewScanner(in)

	for {
		fmt.Fprint(out, PROMPT)
		scanned := scanner.Scan()
		if !scanned {
			return
		}
		line := scanner.Text()
		if line == "exit" || line == "quit" {
			return
		}
		l := lexer.New(line)
		for tok := l.NextToken(); tok.Type != token.EOF; tok = l.NextToken() {
			fmt.Fprintf(out, "%+v\n", tok)
		}
	}
}
