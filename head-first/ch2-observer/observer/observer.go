package observer

type Weather struct {
	Temperature float32
	Humidity    float32
	Pressure    float32
}

type Observer interface {
	Update(weather Weather)
}
