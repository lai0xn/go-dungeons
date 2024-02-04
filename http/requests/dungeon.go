package requests

import uuid "github.com/satori/go.uuid"

type DungeonReq struct {
	Name        string    `json:"name"`
	Description string    `json:"description"`
	UserID      uuid.UUID `json:"id"`
}
