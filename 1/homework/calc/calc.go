package main

import (
	"bufio"
	"bytes"
	"errors"
	"io"
	"os"
	"strconv"
	"strings"
	"unicode"
)

const none = ""

var operatorSet = map[string]struct {
	prec   int
	rAssoc bool
}{
	"^": {4, true},
	"*": {3, false},
	"/": {3, false},
	"+": {2, false},
	"-": {2, false},
	"=": {1, false},
}

type stack []string

func (s *stack) push(value string) {
	*s = append(*s, value)
}

func (s *stack) pop() (string, error) {
	l := len(*s)
	if l == 0 {
		return none, errors.New("Empty Stack")
	}

	res := (*s)[l-1]
	*s = (*s)[:l-1]
	return res, nil
}

func calc(out io.Writer, in io.Reader) error {
	stack := make(stack, 0)
	top := 0
	s := " "

	buff := bufio.NewReader(in)

	s, err := buff.ReadString('e')
	if err != nil {
		if err == io.EOF {

		}
	}
	s = parseInfix(s)

	for _, c := range strings.Fields(s) {

		switch c {
		case "=":
			y, err := stack.pop()
			if err != nil {
				return err
			}
			_, err = out.Write([]byte(y))

			return nil
		case "+":
			y, err := stack.pop()
			if err != nil {
				return err
			}
			z, err := stack.pop()
			if err != nil {
				return err
			}
			val0, _ := strconv.Atoi(z)
			val1, _ := strconv.Atoi(y)
			stack.push(strconv.Itoa(val0 + val1))
			top--
		case "-":
			y, err := stack.pop()
			if err != nil {
				return err
			}
			z, err := stack.pop()
			if err != nil {
				return err
			}
			val0, _ := strconv.Atoi(z)
			val1, _ := strconv.Atoi(y)
			stack.push(strconv.Itoa(val0 - val1))
			top--
		case "*":
			y, err := stack.pop()
			if err != nil {
				return err
			}
			z, err := stack.pop()
			if err != nil {
				return err
			}
			val0, _ := strconv.Atoi(z)
			val1, _ := strconv.Atoi(y)
			stack.push(strconv.Itoa(val0 * val1))
			top--
		case "/":
			y, err := stack.pop()
			if err != nil {
				return err
			}
			z, err := stack.pop()
			if err != nil {
				return err
			}
			if y == "0" {
				return errors.New("Division by zero")
			}
			val0, _ := strconv.Atoi(z)
			val1, _ := strconv.Atoi(y)
			stack.push(strconv.Itoa(val0 / val1))
			top--
		default:
			if c != " " {
				stack.push(c)
				top++
			}
		}
	}
	z, err := stack.pop()
	if err != nil {
		return err
	}
	_, err = out.Write([]byte(z))
	return nil
}

//Перевод в нужную нотацию: 2 + 5 -> 5 2 +
func parseInfix(expression string) (rpn string) {
	stack := make(stack, 0)
	buffer := bytes.Buffer{}
	for i, c := range expression {
		buffer.WriteRune(c)
		if i < len(expression)-1 && !(unicode.IsNumber(c) && unicode.IsNumber(rune(expression[i+1]))) {
			buffer.WriteRune(' ')
		}
	} //Вставляем пробелы для последующей итерации
	expression = buffer.String()

	for _, tok := range strings.Fields(expression) {
		switch tok {
		case "(":
			stack.push(tok)
		case ")":
			for {
				op, err := stack.pop()
				if err != nil {
					return none
				}
				if op == "(" {
					break
				}
				rpn += " " + op
			}
		default:
			if o1, isOp := operatorSet[tok]; isOp {
				for len(stack) > 0 {

					op := stack[len(stack)-1]
					if o2, isOp := operatorSet[op]; !isOp || o1.prec > o2.prec ||
						o1.prec == o2.prec && o1.rAssoc {
						break
					}

					stack = stack[:len(stack)-1]
					rpn += " " + op
				}
				stack = append(stack, tok)
			} else { // токен операнд
				if rpn > "" {
					rpn += " "
				}
				rpn += tok // кидаем операнд в конец
			}
		}

	}
	for len(stack) > 0 {
		rpn += " " + stack[len(stack)-1]
		stack = stack[:len(stack)-1]
	}
	return
}

func main() {
	calc(os.Stdout, os.Stdin)
}
