# gobook
A repo to keep the exercises I do for the Go Programming Language Addison-Wesley book (https://goo.gl/ObTYDn).

## How to
Clone this project into `$GOPATH/src/silvanolte.com/sildani` to play with the examples.

## Interesting programs and packages

1. The "silvanolte.com/sildani/gobook/mystring" package is an exmple library that implements some string manipulation functions. There are unit tests for these functions in the gobook/mystring_test directory. (I wanted to play around with putting unit tests separate from the actual source tests). Before running the test, make sure to `go install` the mystring package first.