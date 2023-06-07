package domains

import entities "mygram/domains/entity"

type UserTokenRepository interface {
	InsertToken(user *entities.User, token string)
	RemoveToken()
}
