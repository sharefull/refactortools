package a

import "a/b"

func F() { G() }
func G() {
	b.B1()
	b.B3()
}