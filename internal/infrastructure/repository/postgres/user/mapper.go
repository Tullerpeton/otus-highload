package user

import (
	"github.com/Tullerpeton/otus-highload/internal/usecase/model"
)

type userRow struct {
	Id         string `db:"id"`
	FirstName  string `db:"first_name"`
	SecondName string `db:"second_name"`
	Birthdate  string `db:"birthdate"`
	Biography  string `db:"biography"`
	City       string `db:"city"`
}

func mapToRaw(user model.User) userRow {
	return userRow{
		Id:         user.Id,
		FirstName:  user.FirstName,
		SecondName: user.SecondName,
		Birthdate:  user.Birthdate,
		Biography:  user.Biography,
		City:       user.City,
	}
}

func mapToModel(raw userRow) model.User {
	return model.User{
		Id:         raw.Id,
		FirstName:  raw.FirstName,
		SecondName: raw.SecondName,
		Birthdate:  raw.Birthdate,
		Biography:  raw.Biography,
		City:       raw.City,
	}
}
