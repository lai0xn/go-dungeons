package controllers

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/golang-jwt/jwt/v5"
	"github.com/gorilla/mux"
	"github.com/jn0x/reddigo/http/requests"
	service "github.com/jn0x/reddigo/services"
	uuid "github.com/satori/go.uuid"
)

type DungeonController interface {
	GetDungeon(w http.ResponseWriter, r *http.Request)
	SearchDungeon(w http.ResponseWriter, r *http.Request)
	CreateDungeon(w http.ResponseWriter, r *http.Request)
	DeleteDungeon(w http.ResponseWriter, r *http.Request)
	JoinDungeon(w http.ResponseWriter, r *http.Request)
}

type dungeonController struct {
	serv service.DungeonService
}

func (c *dungeonController) CreateDungeon(w http.ResponseWriter, r *http.Request) {
	var reqbody requests.DungeonReq
	json.NewDecoder(r.Body).Decode(&reqbody)
	c.serv.CreateDungeon(reqbody)
}

func (c *dungeonController) GetDungeon(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	id := params["id"]
	uid, err := uuid.FromString(id)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("something went wrong"))
		return
	}
	dungen := c.serv.GetDungeonByID(uid)

	w.WriteHeader(http.StatusOK)
	response, err := json.Marshal(dungen)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("something went wrong"))
		return
	}
	w.Write(response)
}

func (c *dungeonController) SearchDungeon(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	query := params["name"]
	dungeons := c.serv.SearchDungeon(query)
	response, err := json.Marshal(dungeons)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	w.Write(response)
}

func (c *dungeonController) DeleteDungeon(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	uid := params["id"]
	id, err := uuid.FromString(uid)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("something wrong happend"))
	}
	c.serv.DeleteDungeon(id)
}

func (c *dungeonController) JoinDungeon(w http.ResponseWriter, r *http.Request) {
	userClaims := r.Context().Value("claims").(jwt.MapClaims)
	fmt.Println(userClaims["id"])
}

func NewDungeonController() DungeonController {
	return &dungeonController{serv: service.NewDungeonService()}
}
