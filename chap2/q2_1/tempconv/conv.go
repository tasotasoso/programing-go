package tempconv

func CToF(c Celsius) Fahrenheit     { return Fahrenheit(c*9/5 + 32) }
func FToC(f Fahrenheit) Celsius     { return Celsius((f - 32) * 5 / 9) }
func CToK(c Celsius) AbsoluteZeroC  { return AbsoluteZeroC(c + 273.15) }
func KToC(ab AbsoluteZeroC) Celsius { return Celsius(ab - 273.15) }
