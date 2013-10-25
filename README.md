Pher
====

A very early/ugly implementation of the PHP of Go.

This program basically takes an input `.gopher` file (HTML mixed
with Go code in `<? … ?>` statements) and outputs a Go program.


Example
-------

To compile the example file, just run this on the command line:

    go run main.go -i example.gopher

The output will be printed to stdout.
