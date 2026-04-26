package user

import (
	"net/http"
)

type Signup struct{}

func NewSignup() Signup {
	return Signup{}
}

func (s *Signup) Handle(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(200)
	w.Write([]byte("OK"))
}
