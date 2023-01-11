package repo

import (
	"github.com/kkakoz/ormx"
	"sync"
	"video-rpc/internal/user/model/entity"
)

type userRepo struct {
	ormx.IRepo[entity.User]
}

var userOnce sync.Once
var _user *userRepo

func User() *userRepo {
	userOnce.Do(func() {
		_user = &userRepo{IRepo: ormx.NewRepo[entity.User]()}
	})
	return _user
}
