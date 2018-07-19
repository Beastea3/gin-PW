package models

type Joke struct {
	ID		int		`json:"id" binding:"required"`
	Likes	int		`joson:"likes"`
	Joke	string	`json:"joke" binding:"required"`
}

