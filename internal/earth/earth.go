package earth

// Coordinate 绝对坐标
type Coordinate struct {
	year      uint16
	month     uint8
	day       uint8
	longitude float64 // 经度
	latitude  float64 // 维度
}

// NewCoordinate contribute a Coordinate instance(maybe)
func NewCoordinate(year uint16, month, day uint8, longitude, latitude float64) (cd *Coordinate) {
	return &Coordinate{year, month, day, longitude, latitude}
}

// GetDistance return the distance between the given coodrinate
func (c *Coordinate) GetDistance(rhs *Coordinate) float64 {
	dst := 0.0
	return dst
}

// GetSunriseTime return the sunrise time
func (c *Coordinate) GetSunriseTime() (time uint64) {
	return time
}

// GetSunsetTime return the sunset time
func (c *Coordinate) GetSunsetTime() (time uint64) {
	return time
}
