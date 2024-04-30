package main

import (
	"log"
	"os"
	"path"
	"strings"

	"github.com/csnewman/dextk"
	"github.com/csnewman/dextk/internal/decompiler"
	"golang.org/x/exp/mmap"
)

func main() {
	f, err := mmap.Open("classes4.dex")
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

		// TODO: Add better checking
		if strings.Contains(node.Name, ".") {
			panic("invalid name: " + node.Name)
		}

		out := decompiler.ProcessClass(r, node)

		dst := "./out/" + node.Name + ".java"

		dir, _ := path.Split(dst)

		if err := os.MkdirAll(dir, 0o750); err != nil {
			panic(err)
		}

		if err := os.WriteFile(dst, []byte(out), 0o666); err != nil {
			panic(err)
		}
	}

}
