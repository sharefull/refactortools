package a

func f(hogeID int) { // want "hogeID's type should be user defined type but int"
	_ = hogeID // want "hogeID's type should be user defined type but int"
}
