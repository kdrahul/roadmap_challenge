package user

import (
	"e-commerce/db"
	"net/http"
)

type Login struct {
	db db.DBInterface
}

func NewLogin(db db.DBInterface) Login {
	return Login{
		db: db,
	}
}

func (u *Login) Handle(w http.ResponseWriter, r *http.Request) {
	u.db.GetOne()
}
