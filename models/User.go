package models

import "time"

type User struct {
	ID             string    `json:"id"`
	Username       string    `json:"username"`
	Bio            string    `json:"bio"`
	Birthday       string    `json:"birthday"`
	ProfilePicture string    `json:"profile_picture"`
	CreationDate   time.Time `json:"creation_date"`
}
