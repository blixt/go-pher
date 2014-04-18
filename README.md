Pher /fə(ɹ)/
=============

A very early/ugly implementation of the PHP of Go.

This program basically takes an input `.gopher` file (HTML mixed
with Go code in `<? … ?>` statements) and outputs a Go program.


Set up
------

To set it up, just install this package:

```bash
go get github.com/blixt/go-pher
```


Example
-------

To run the example:

```bash
# Convert the .gopher file to a .go file
go run main.go -i example.gopher > example.go
# Build the .go file into a cgi-bin directory
mkdir cgi-bin
go build -o cgi-bin/example example.go
# Start up a CGI server (Python makes it easy)
python -m CGIHTTPServer
# Open the page!
open http://localhost:8000/cgi-bin/example?name=world
```

See below for the resulting contents of each file.

### example.gopher

```html
<!DOCTYPE html>
<html>
<head>
	<title>Test</title>
</head>
<body>
<?
func greet(name string) string{
  return "what up, " + name
}
?>
	<h1><?= greet(pher.Get("name")) ?></h1>
</body>
</html>
```

### example.go

```go
package main
import (
"fmt"
"github.com/blixt/go-pher/pher"
)
func greet(name string) string{
  return "what up, " + name
}
func main() {
fmt.Print("Content-Type: text/html\r\n\r\n")
fmt.Print("<!DOCTYPE html>\n<html>\n<head>\n\t<title>Test</title>\n</head>\n<body>\n")
fmt.Print("\t<h1>")
fmt.Print(greet(pher.Get("name")))
fmt.Print("</h1>\n</body>\n</html>\n")
}
```

### Output HTML

```html
<!DOCTYPE html>
<html>
<head>
	<title>Test</title>
</head>
<body>
	<h1>what up, world</h1>
</body>
</html>
```
