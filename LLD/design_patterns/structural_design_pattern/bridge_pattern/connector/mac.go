package connector


type Mac struct {
	p Printer
}

func NewMac(p Printer) Mac {
	return Mac{p}
}

func (m Mac) PrintDoc() {
	m.p.Print("Mac")
}