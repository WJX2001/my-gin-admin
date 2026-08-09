package request

import (
	"github.com/WJX2001/gin-vue-admin-server/model/system"
	jwt "github.com/golang-jwt/jwt/v5"
	"github.com/google/uuid"
)

type CustomClaims struct {
	BaseClaims
	BufferTime    int64
	MustChangePwd bool `json:"mustChangePwd"`
	jwt.RegisteredClaims
}

type BaseClaims struct {
	UUID        uuid.UUID
	ID          uint
	Username    string
	Nickname    string
	AuthorityId uint
	UserType    system.UserType
}
