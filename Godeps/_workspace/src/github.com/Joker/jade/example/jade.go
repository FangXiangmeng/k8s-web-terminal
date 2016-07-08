package main

import (
	"fmt"
	"github.com/beyondblog/k8s-web-terminal/Godeps/_workspace/src/github.com/Joker/jade"
	"io/ioutil"
)

func main() {
	dat, err := ioutil.ReadFile("template.jade")
	if err != nil {
		fmt.Printf("ReadFile error: %v", err)
		return
	}

	tmpl, err := jade.Parse("name_of_tpl", string(dat))
	if err != nil {
		fmt.Printf("Parse error: %v", err)
		return
	}

	fmt.Printf("\nOutput:\n\n%s", tmpl)
}
