Pher /fɜː(ɹ)/
=============

A very early/ugly implementation of the PHP of Go.

This program basically takes an input `.gopher` file (HTML mixed
with Go code in `<? … ?>` statements) and outputs a Go program.


Example
-------

To compile the example file, just run this on the command line:

    go run main.go -i example.gopher

This will turn the following *gopher* template:

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

Into this Go code (and output it on stdout):

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

Which in turn, when executed with `go run`, will output this:

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
