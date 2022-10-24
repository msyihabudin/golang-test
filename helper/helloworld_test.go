package helper

import "testing"

func TestHelloWorld_success(t *testing.T) {
	result := HelloWorld("Syihab")
	if result != "Hello Syihab" {
		panic("Result is not Hello Syihab")
	}
}