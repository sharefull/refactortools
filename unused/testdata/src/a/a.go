package a

import (
	"a/b"
	"fmt"
)

type Atype struct {
	AAA string
}

func F() { G() }
func G() {
	b.B1()
	b.B3()

	a := Atype{}
	fmt.Println(a.AAA)
}

func H() {
	b4 := b.New()
	b4.Get()

	all := b.All{
		A: b4,
	}

	all.A.Set()
}
