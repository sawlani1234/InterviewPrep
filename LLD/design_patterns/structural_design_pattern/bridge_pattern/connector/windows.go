package connector



type Windows struct {
	p Printer
}

func NewWindows(p Printer) Windows {
	return Windows{p}
}

func (m Windows) PrintDoc() {
	m.p.Print("Windows")
}