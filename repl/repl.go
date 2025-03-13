package repl

import (
	"bufio"
	"finger/evaluator"
	"finger/lexer"
	"finger/object"
	"finger/parser"
	"fmt"
	"io"
)

const PROMPT = ">> "

// 指纹语言模型
const FINGER_MODEL = `
	Hello, I'm the Finger programming language!
	
	Here are some commands:
`

func Start(in io.Reader, out io.Writer) {
	scanner := bufio.NewScanner(in)
	env := object.NewEnvironment()

	for {
		fmt.Fprintf(out, PROMPT)
		scanned := scanner.Scan()
		if !scanned {
			return
		}

		line := scanner.Text()
		l := lexer.New(line)
		p := parser.New(l)

		program := p.ParseProgram()

		if len(p.Errors()) != 0 {
			printParserErrors(out, p.Errors())
			continue
		}

		evaluated := evaluator.Eval(program, env)

		if evaluated != nil {
			io.WriteString(out, evaluated.Inspect())
			io.WriteString(out, "\n")
		}	
	}
}

func printParserErrors(out io.Writer, errors []string) {
	io.WriteString(out, FINGER_MODEL)
	io.WriteString(out, "Woops! We ran into some finger erros:\n")
	io.WriteString(out, " parser errors:\n")
	for _, msg := range errors {
		io.WriteString(out, "\t"+msg+"\n")
	}
}
