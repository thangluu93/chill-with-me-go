package models

type Movie struct {
	Id          string `json:"id" bson:"_id,omitempty"`
	Title       string `json:"title" bson:"title"`
	Description string `json:"description" bson:"description"`
	Year        int    `json:"year" bson:"year"`
	Genre       string `json:"genre" bson:"genre"`
	IsAchieve   bool   `json:"is_achieve" bson:"is_archive"`
}
