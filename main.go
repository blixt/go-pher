package main

import (
	"bytes"
	"flag"
	"fmt"
	"go/parser"
	"go/token"
	"io/ioutil"
	"os"
	"regexp"
	"strings"
)

func main() {
	var input string
	flag.StringVar(&input, "i", "", "input file")
	flag.Parse()

	if len(input) == 0 {
		fmt.Fprintln(os.Stderr, "no input file given")
		os.Exit(1)
	}

	var src string
	if data, err := ioutil.ReadFile(input); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	} else {
		src = string(data)
	}

	const (
		startTag = "<\\?=?\\s*"
		endTag   = "\\?>(?:\\r\\n|\\r|\\n|)"
		untilTag = "((?:[^<]+|<[^?])*)"
	)
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
			if src[i-2] == '=' {
				code = "fmt.Print(" + code + ")"
			}

			if _, err := parser.ParseExpr(code); err != nil {
				if _, err := parser.ParseFile(fs, "", "package p;"+code, 0); err != nil {
					fmt.Fprintf(os.Stderr, "error in %s[%d]: %s\n", input, 0, err)
					os.Exit(1)
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

	if i < len(src)-1 {
		fmt.Fprintln(os.Stderr, "parse failed")
		os.Exit(1)
	}

	fmt.Println("package main")
	fmt.Println("import (")
	fmt.Println(`"fmt"`)
	fmt.Println(`"github.com/blixt/go-pher/pher"`)
	fmt.Println(")")
	fmt.Print(head)
	fmt.Println("func main() {")
	fmt.Println(`fmt.Print("Content-Type: text/html\r\n\r\n")`)
	fmt.Print(body)
	fmt.Println("}")
}
