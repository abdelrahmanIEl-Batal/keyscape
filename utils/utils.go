package utils

import "fmt"

func Debug(s string) {
	fmt.Println(fmt.Sprintf("%s", s))
}

func DebugError(e error, operation string) {
	fmt.Println(fmt.Sprintf("error in operation: %s, error message: %s", operation, e.Error()))
}
