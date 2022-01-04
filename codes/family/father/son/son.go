package son

import "fmt"

func init() {
	fmt.Println("Father package initialized.")
}

type Son struct {
	Name string
}

func (f Son) Data(name string) string {
	f.Name = "Son: " + name
	return f.Name
}
