package coreFunctions

import "fmt"

func Hello(name string)(s string, e error){
	s=fmt.Sprintf("Hello %s",name)
	return
}

