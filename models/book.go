package models

import "time"

type Book struct {
	Id int
	UserId int
	User User
	HotelId int
	Hotel Hotel
	StartDate time.Time
	EndDate time.Time
}