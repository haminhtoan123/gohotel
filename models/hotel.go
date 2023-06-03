package models

type Hotel struct {
	Name      string
	Latitude  float64
	Longitude float64
	Address   string
}

func (h *Hotel) Update(name string, latitude float64, longitude float64, address string) {
	h.Name = name
	h.Latitude = latitude
	h.Longitude = longitude
	h.Address = address
}
