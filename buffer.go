package main

import (
	"bytes"
	"fmt"
)

func main() {
	var b = bytes.NewBuffer(make([]byte, 26))
	var texts = []string{
		`This fork has had a few improvements by ourselves as well as several
PR's merged`,


		`Further PR's (with tests) are welcome,
but please maintain backwards compatibility.`,

		`Detailed documentation of the API is available at GoDoc.`,

		`A sub-package that implements the BSON specification is also
included,`,
		`and may be used independently of the driver,`,
		`from the original mgo repo that are currently awaiting review.
Changes are mostly geared towards performance improvements and bug fixes,
though a few new features have been added.`,
	}
	for i := range texts {
		b.Reset()
		b.WriteString(texts[i])
		fmt.Println("Length:", b.Len(), "\tCapacity:", b.Cap())
	}
}
