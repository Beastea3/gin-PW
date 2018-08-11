package models

import "gopkg.in/mgo.v2/bson"

const (
	// CollectionArticle holds the name of the articles collection
	CollectionJoke = "Jokes"
)


type Joke struct {
	Id        bson.ObjectId `json:"_id,omitempty" bson:"_id,omitempty"`
	ID		int		`json:"id" binding:"required"`
	Likes	int		`json:"likes"`
	Joke	string	`json:"joke" binding:"required"`
}

//func getJokes()  []Joke {
//
//}

