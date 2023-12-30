# dextk
Android dex file parser in Go

### Features

- String, classes and code parsing
- Low memory usage
- Mmap support
- Designed to support stream processing (No requirement to load entire dex file into memory)

### Example

```bash
go run ./cmd/example
```

```go
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
```

Output:

```
android/support/v4/app/INotificationSideChannel$_Parcel :  <init>
   invoke-direct method=java/lang/Object:<init>:([]):V args=[0]
   return-void value=-1
[...]
android/support/v4/app/INotificationSideChannel$_Parcel :  readTypedObject
   invoke-virtual method=android/os/Parcel:readInt:([]):I args=[1]
   move-result dst=0
   if-eqz a=0 b=-1 tgt=6
   invoke-interface method=android/os/Parcelable$Creator:createFromParcel:([Landroid/os/Parcel;]):Ljava/lang/Object; args=[2 1]
   move-result-object dst=0
   return-object value=0
   const/4 dst=0 value='0'
   return-object value=0
[...]
```