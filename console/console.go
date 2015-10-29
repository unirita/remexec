// Pakcage console implements functions to create or display console messages.
package console

import (
	"fmt"
	"os"
)

var msgs = map[string]string{
	"REX001E": "INVALID ARGUMENT - %s",
}

// Display outputs formatted message associated with code to stdout.
func Display(code string, a ...interface{}) (int, error) {
	msg := GetMessage(code, a...)
	return fmt.Println(msg)
}

// DisplayError outputs formatted message associated with code to stderr.
func DisplayError(code string, a ...interface{}) (int, error) {
	msg := GetMessage(code, a...)
	return fmt.Fprintln(os.Stderr, msg)
}

// GetMessage creates formatted message associated with code.
func GetMessage(code string, a ...interface{}) string {
	return fmt.Sprintf("%s %s", code, fmt.Sprintf(msgs[code], a...))
}
