package models

type Movie struct {
	Id          int    `json:"id" bson:"id"`
	Title       string `json:"title" bson:"title"`
	Description string `json:"description" bson:"description"`
	Year        int    `json:"year" bson:"year"`
	IsAchieve   bool   `json:"is_achieve" bson:"is_archive"`
}
