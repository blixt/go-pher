Pher /fə(ɹ)/
=============

A very early/ugly implementation of the PHP of Go.

This program basically takes an input `.gopher` file (HTML mixed
with Go code in `<? … ?>` statements) and outputs a Go program.


Example
-------

To run the example:

```bash
# Convert the .gopher file to a .go file
go run main.go -i example.gopher > example.go
# Run the .go file to get the HTML
go run example.go > example.html
# Open the HTML file!
open example.html
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
	<h1><?= greet("blixt") ?></h1>
</body>
</html>
```

### example.go

```go
package main
import "fmt"
func greet(name string) string{
  return "what up, " + name
}
func main() {
fmt.Print("<!DOCTYPE html>\n<html>\n<head>\n\t<title>Test</title>\n</head>\n<body>\n")
fmt.Print("\t<h1>")
fmt.Print(greet("blixt"))
fmt.Print("</h1>\n</body>\n</html>\n")
}
```

### example.html

```html
<!DOCTYPE html>
<html>
<head>
	<title>Test</title>
</head>
<body>
	<h1>what up, blixt</h1>
</body>
</html>
```
