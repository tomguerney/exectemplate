package exectemplate

import (
	"fmt"
	"strings"
	"text/template"

	"github.com/mattn/go-shellwords"
)

// Exec executes
func Exec(template string, values map[string]string) error {
	command, err := interpolate(template, values)
	if err != nil {
		return err
	}
	fmt.Println(command)
	return nil
}

func interpolate(tmpl string, values map[string]string) (string, error) {
	builder := strings.Builder{}
	parsed, err := template.New("stencil").Parse(tmpl)
	if err != nil {
		return "", err
	}
	err = parsed.Execute(&builder, values)
	if err != nil {
		return "", err
	}
	return builder.String(), nil
}

func sliceArgs(args string) ([]string, error) {
	slice, err := shellwords.Parse(args)
	if err != nil {
		return nil, err
	}
	return slice, nil
}
