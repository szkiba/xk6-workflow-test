package workflow_test

import "fmt"

func (*module) greeting(name string) string {
	if name == "" {
		name = "World"
	}

	return fmt.Sprintf("Hello, %s!", name)
}
