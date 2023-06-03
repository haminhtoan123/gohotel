package repo

import model "github.com/haminhtoan123/gohotel/models/models"

type HotelRepository interface {
	Add(hotel *model.Hotel) error
	Update(hotel *model.Hotel) error
	Get(name string) (*model.Hotel, error)
}
