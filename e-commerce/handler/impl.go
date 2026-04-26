package handler

import (
	"context"
	"e-commerce/api"
	"e-commerce/db"
	"net/http"
)

type Server struct {
	db *db.DBInterface
}

func NewServer(db db.DBInterface) *Server {
	return &Server{
		&db,
	}
}

func (a *Server) PostAuthLogin(ctx context.Context, request api.PostAuthLoginRequestObject) (api.PostAuthLoginResponseObject, error) {
	email := "abc@mail.com"
	id := "123"

	return api.PostAuthLogin200JSONResponse{
		User: &api.User{
			Email: &email,
			Id:    &id,
		},
	}, nil
}
func (a *Server) PostAuthLogout(ctx context.Context, request api.PostAuthLogoutRequestObject) (api.PostAuthLogoutResponseObject, error) {
	return nil, nil
}
func (a *Server) GetAuthMe(ctx context.Context, request api.GetAuthMeRequestObject) (api.GetAuthMeResponseObject, error) {
	email := "rahul@abc.com"
	id := "rah123"
	if false {
		return api.GetAuthMe401Response{}, nil
	}

	if true {
		return api.GetAuthMe500JSONResponse{
			Error:   "",
			Message: "",
		}, nil
	}
	return api.GetAuthMe200JSONResponse{
		Email: &email,
		Id:    &id,
	}, nil
}
func (a *Server) PostAuthSignup(ctx context.Context, request api.PostAuthSignupRequestObject) (api.PostAuthSignupResponseObject, error) {
	return nil, nil
}

// func (a *Server) PostUserId(ctx context.Context, request api.PostUserIdRequestObject) (api.PostUserIdResponseObject, error) {
// 	return nil, nil
// }
