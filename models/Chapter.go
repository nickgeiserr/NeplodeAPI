package models

import "time"

type Chapter struct {
	ChapterID    string    `json:"chapter_id"`
	Name         string    `json:"name"`
	Description  string    `json:"description"`
	OwnerUserID  string    `json:"owner_user_id"`
	CreationDate time.Time `json:"creation_date"`
	ModifiedDate time.Time `json:"modified_date"`
	LogoURL      string    `json:"logo_url"`
	BannerURL    string    `json:"banner_url"`
}
