package exectemplate

import "fmt"

// Exec executes
func Exec(template string, values map[string]string) {
	fmt.Println("template is " + template)
}
