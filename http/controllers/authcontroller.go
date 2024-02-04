package controllers

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/jn0x/reddigo/http/requests"
	service "github.com/jn0x/reddigo/services"
)

type AuthController interface {
	Login(w http.ResponseWriter, r *http.Request)
	Signup(w http.ResponseWriter, r *http.Request)
}

type authController struct {
	authService service.AuthService
}

func NewAuthController() AuthController {
	return AuthController(&authController{authService: service.NewAuthService()})
}

func (c *authController) Login(w http.ResponseWriter, r *http.Request) {
	var user requests.AuthReq
	json.NewDecoder(r.Body).Decode(&user)
	token, errs := c.authService.Login(user)
	resp := make(map[string]string)

	if errs != nil {
		for index, err := range errs {
			resp[fmt.Sprintf("error%d", index)] = err.Error()
		}
		json_res, err := json.Marshal(&resp)
		if err != nil {
			w.Write([]byte(err.Error()))
			return
		}
		w.Write(json_res)
		return
	}

	w.Header().Set("Content-Type", "application/json")

	resp["token"] = token
	json_res, err := json.Marshal(&resp)
	if err != nil {
		w.Write([]byte(err.Error()))
	}
	w.Write(json_res)
}

func (c *authController) Signup(w http.ResponseWriter, r *http.Request) {
	var user requests.AuthReq
	json.NewDecoder(r.Body).Decode(&user)
	errs := c.authService.Signup(user)

	resp := make(map[string]string)

	if errs != nil {
		for index, err := range errs {
			resp[fmt.Sprintf("error%d", index)] = err.Error()
		}
		json_res, err := json.Marshal(&resp)
		if err != nil {
			w.Write([]byte(err.Error()))
			return
		}
		w.Write(json_res)
		return
	}

	w.Header().Set("Content-Type", "application/json")

	resp["message"] = "user created"
	json_res, err := json.Marshal(&resp)
	if err != nil {
		w.Write([]byte(err.Error()))
	}
	w.Write(json_res)
}
