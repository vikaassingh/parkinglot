package model

type Parkinglot struct {
	ID        uint
	Name      string
	Address   string
	Floors    []Floor
	EntryGate Gate
	ExitGate  Gate
}
