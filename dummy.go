package ms_forms

type Dummy struct {
}

func NewDummy() *Dummy {
	result := new(Dummy)
	return result
}

func (d *Dummy) isABC() bool {
	return true
}
