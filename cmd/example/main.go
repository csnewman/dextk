package main

import (
	"fmt"
	"log"

	"github.com/csnewman/dextk"
	"golang.org/x/exp/mmap"
)

func main() {
	f, err := mmap.Open("classes.dex")
	if err != nil {
		log.Panicln(err)
	}

	defer f.Close()

	r, err := dextk.Read(f)
	if err != nil {
		log.Panicln(err)
	}

	ci := r.ClassIter()
	for ci.HasNext() {
		node, err := ci.Next()
		if err != nil {
			log.Panicln(err)
		}

		for _, method := range node.DirectMethods {
			fmt.Println(node.Name, ": ", method.Name)
			processMeth(r, method)
		}

		for _, method := range node.VirtualMethods {
			fmt.Println(node.Name, ": ", method.Name)
			processMeth(r, method)
		}
	}
}

func processMeth(r *dextk.Reader, m dextk.MethodNode) {
	if m.CodeOff == 0 {
		return
	}

	c, err := r.ReadCodeAndParse(m.CodeOff)
	if err != nil {
		log.Panic(err)
	}

	for _, o := range c.Ops {
		fmt.Println("  ", o)
	}
}
