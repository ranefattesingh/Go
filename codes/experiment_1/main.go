/* experimental file*/

package main

import "fmt"

func main() {
	for i, v := range "hello" {
		fmt.Println(i, string(v))
	}
}
