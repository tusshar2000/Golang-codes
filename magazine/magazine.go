package magazine

type Subscribers struct {
	Name   string
	Rate   float64
	Active bool
	Address
}

type Employee struct {
	Name   string
	salary float64
	Address
}

type Address struct {
	Building string
	Street   string
}
