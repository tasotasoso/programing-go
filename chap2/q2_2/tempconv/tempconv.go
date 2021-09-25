package tempconv

import "fmt"

type Celsius float64
type Fahrenheit float64
type AbsoluteZeroC float64

const (
	FreezingC Celsius = 0
	BoilingC  Celsius = 100
)

func (c Celsius) String() string        { return fmt.Sprintf("%g°C", c) }
func (f Fahrenheit) String() string     { return fmt.Sprintf("%g°F", f) }
func (ab AbsoluteZeroC) String() string { return fmt.Sprintf("%g°K", ab) }
