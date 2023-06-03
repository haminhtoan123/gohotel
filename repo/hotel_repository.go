package repo

import (
	"myhotel/models/model"
)

type HotelRepository interface {
	Add(hotel *model.Hotel) error
	Update(hotel *model.Hotel) error
	Get(name string) (*model.Hotel, error)
}
