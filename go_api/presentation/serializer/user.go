package serializer

import (
	"github.com/google/uuid"
	"github.com/nozomi-iida/nozo_blog/entity"
)

type UserJson struct {
	UserID uuid.UUID `json:"userId"`
	UserName string `json:"userName"`
}

func User2Json(user entity.User) UserJson {
	return UserJson{
		UserID: user.UserId.ID,
		UserName: user.Username,
	}	
}
