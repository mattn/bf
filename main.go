package main

import (
	"bytes"
	"fmt"
	"io"
	"io/ioutil"
	"os"
)

func brainfuck(src []rune) {
	slen := len(src)
	data := []byte{0}
	dlen := 0

	for pc := 0; pc < slen; pc++ {
		switch src[pc] {
		case '+':
			data[dlen] += 1
		case '-':
			data[dlen] -= 1
		case '>':
			dlen += 1
			if len(data) <= dlen {
				data = append(data, 0)
			}
		case '<':
			if dlen > 0 {
				dlen -= 1
			}
		case '.':
			fmt.Print(string(data[dlen]))
		case ',':
			b := make([]byte, 1)
			os.Stdin.Read(b)
			data[dlen] = b[0]
		case '[':
			if data[dlen] != 0 {
				break
			}
			for depth := 1; depth > 0; {
				pc++
				srcCharacter := src[pc]
				if srcCharacter == '[' {
					depth++
				} else if srcCharacter == ']' {
					depth--
				}
			}
		case ']':
			for depth := 1; depth > 0; {
				pc--
				srcCharacter := src[pc]
				if srcCharacter == '[' {
					depth--
				} else if srcCharacter == ']' {
					depth++
				}
			}
			pc--
		}
	}
}

func main() {
	var content bytes.Buffer

	for _, f := range os.Args[1:] {
		b, err := ioutil.ReadFile(f)
		if err != nil {
			fmt.Fprintln(os.Stderr, "read file:", err)
			os.Exit(1)
		}
		content.Write(b)
	}

	if content.Len() == 0 {
		if _, err := io.Copy(&content, os.Stdin); err != nil {
			fmt.Fprintln(os.Stderr, "read stdin:", err)
			os.Exit(1)
		}
	}

	brainfuck([]rune(content.String()))
}
