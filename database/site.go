package database

type Tariff struct {
	ID             int
	Type           int
	Price          float64
	Name           string
	Description    string
	Speed          int
	DigitalChannel int
	AnalogChannel  int
	Image          string
	ColorType      int
}
