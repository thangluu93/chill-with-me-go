package models

type User struct {
	Id          string `json:"id" bson:"id"`
	Email       string `json:"email" bson:"email"`
	Password    string `json:"password" bson:"password"`
	PhoneNumber string `json:"phoneNumber" bson:"phoneNumber"`
	IsActive    bool   `json:"isActive" bson:"isActive"`
	FirstName   string `json:"firstName" bson:"firstName"`
	LastName    string `json:"lastName" bson:"lastName"`
	UpdateAt    int    `json:"updateAt" bson:"updateAt"`
	CreatedAt   int    `json:"createdAt" bson:"createdAt"`
}
