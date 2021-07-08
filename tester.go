package main

import (
	"fmt"
	"strings"
	"time"
)

type name struct {
	sirName    string
	actualName string
	prefix     string
}
type namingConvention interface {
	formatting() string
	renaming(newName string) string
}

func (n *name) formatting() string {
	return fmt.Sprintf("The full name of the person is %s.%s %s\n", n.prefix, n.actualName, n.sirName)
}
func (n *name) renaming(newName string) string {
	n.actualName = newName
	return fmt.Sprintf("The new full name of the person is %s.%s %s\n", n.prefix, n.actualName, n.sirName)
}

func main() {

	go func() {
		fmt.Println("this is inside go routine")
	}()

	time.Sleep(time.Millisecond * 100)
	var n namingConvention = &name{"Are", "manikanta", "Mr"}
	fmt.Println(n.formatting())
	fmt.Println(n.renaming("Manikanta vinod kumar"))
	fmt.Println(strings.ToUpper("goher"))

}
