package repo

import (
	"github.com/kkakoz/ormx"
	"github.com/kkakoz/video-rpc/internal/user/model/entity"
	"sync"
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
