package b

type app struct {
	AllModel All
}

type All struct {
	A AInterface
}

type AInterface interface {
	Get() bool
	Set()
	Put()
}

type b4Type struct {
}

func New() AInterface {
	return &b4Type{}
}

func NewApp() app {
	am := All{
		A: New(),
	}

	return app{
		AllModel: am,
	}
}

func (b4 *b4Type) Get() bool {
	return true
}
func (b4 *b4Type) Set() {}

func (b4 *b4Type) Put() {}

func B1() {
	a := NewApp()
	a.Send()
}
func B2() {}
func B3() {}
