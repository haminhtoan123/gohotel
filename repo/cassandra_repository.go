package repo

import (
	"github.com/haminhtoan123/gohotel/models/model"
	"github.com/haminhtoan123/gohotel/models/repo"

	"github.com/gocql/gocql"
)

type CassandraHotelRepository struct {
	session *gocql.Session
}

func NewCassandraHotelRepository(session *gocql.Session) *repo.HotelRepository {
	return &CassandraHotelRepository{session: session}
}

func (r *CassandraHotelRepository) Add(hotel *model.Hotel) error {
	query := "INSERT INTO hotels (name, latitude, longitude, address) VALUES (?, ?, ?, ?)"
	return r.session.Query(query, hotel.Name, hotel.Latitude, hotel.Longitude, hotel.Address).Exec()
}

func (r *CassandraHotelRepository) Update(hotel *model.Hotel) error {
	query := "UPDATE hotels SET latitude = ?, longitude = ?, address = ? WHERE name = ?"
	return r.session.Query(query, hotel.Latitude, hotel.Longitude, hotel.Address, hotel.Name).Exec()
}

func (r *CassandraHotelRepository) Get(name string) (*model.Hotel, error) {
	query := "SELECT name, latitude, longitude, address FROM hotels WHERE name = ? LIMIT 1"
	var hotel model.Hotel
	err := r.session.Query(query, name).Scan(&hotel.Name, &hotel.Latitude, &hotel.Longitude, &hotel.Address)
	if err != nil {
		return nil, err
	}
	return &hotel, nil
}
