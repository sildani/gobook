package mystring

import (
  "testing"
  "silvanolte.com/sildani/gobook/mystring"
)

func TestReverse(t *testing.T) {
  input := "hello"
  output := mystring.Reverse(input)
  expected := "olleh"
  if output != expected {
    t.Errorf("Reverse(%q) == %q but expected %q", input, output, expected)
  }
}

func TestReverseUsingTable(t *testing.T) {
  var tests = []struct {
    input, expected string
  }{
    {"Backward", "drawkcaB"},
    {"Hello, 汉字", "字汉 ,olleH"},
    {"", ""},
  }

  for _, test := range tests {
    output := mystring.Reverse(test.input)
    if output != test.expected {
      t.Errorf("Reverse(%q) == %q but expected %q", test.input, output, test.expected)
    }
  }
}