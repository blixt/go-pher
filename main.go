package main

import (
	"bytes"
	"flag"
	"fmt"
	"go/parser"
	"go/token"
	"io/ioutil"
	"regexp"
	"strings"
)

func main() {
	var input string
	flag.StringVar(&input, "i", "", "input file")
	flag.Parse()

	if len(input) == 0 {
		panic("no input file given")
	}

	var src string
	if data, err := ioutil.ReadFile(input); err != nil {
		panic(err)
	} else {
		src = string(data)
	}

	startTag := "<\\?=?\\s*"
	endTag := "\\?>(?:\\r\\n|\\r|\\n|)"
	untilTag := "((?:[^<]+|<[^?])*)"
	blocks := regexp.MustCompile("\\A" + untilTag + startTag + "|" + endTag + untilTag + startTag + "|" + endTag + "((?s).*)\\z")


	fs := token.NewFileSet()

	body := bytes.NewBufferString("")
	head := bytes.NewBufferString("")
	var i int

	matches := blocks.FindAllStringSubmatchIndex(src, -1)
	for _, indices := range matches {
		var match []int
		for i := 2; i < 8; i += 2 {
			if indices[i] > -1 {
				match = indices[i : i+2]
				break
			}
		}

		if i != indices[0] {
			code := strings.TrimSpace(src[i:indices[0]])
			if src[i - 1] == '=' {
				code = "fmt.Print(" + code + ")"
			}

			if _, err := parser.ParseExpr(code); err != nil {
				if _, err := parser.ParseFile(fs, "", "package p;" + code, 0); err != nil {
					panic(fmt.Sprintf("error in %s[%d]: %s", input, 0, err))
				} else {
					fmt.Fprintln(head, code)
				}
			} else {
				fmt.Fprintln(body, code)
			}
		}
		i = indices[1]

		if match[0] != match[1] {
			fmt.Fprintf(body, "fmt.Print(%q)\n", src[match[0]:match[1]])
		}
	}

	if i < len(src) - 1 {
		panic("parse failed")
	}

	fmt.Println("package main")
	fmt.Println("import \"fmt\"")
	fmt.Print(head)
	fmt.Println("func main() {")
	fmt.Print(body)
	fmt.Println("}")
}
