package logic

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"github.com/Shopify/sarama"
	"github.com/go-sql-driver/mysql"
	"github.com/jinzhu/copier"
	"github.com/kkakoz/ormx"
	"github.com/kkakoz/ormx/opt"
	"github.com/kkakoz/pkg/cryption"
	"github.com/kkakoz/pkg/errno"
	"github.com/kkakoz/pkg/redisx"
	"github.com/kkakoz/video-rpc/internal/pkg/async/producer"
	"github.com/kkakoz/video-rpc/internal/pkg/emailx"
	"github.com/kkakoz/video-rpc/internal/pkg/errs"
	"github.com/kkakoz/video-rpc/internal/pkg/keys"
	"github.com/kkakoz/video-rpc/internal/user/logic/internal/repo"
	"github.com/kkakoz/video-rpc/internal/user/model/entity"
	"github.com/kkakoz/video-rpc/pb/common"
	userpb "github.com/kkakoz/video-rpc/pb/user"
	uuid "github.com/satori/go.uuid"
	"github.com/spf13/viper"
	"google.golang.org/protobuf/proto"
	"google.golang.org/protobuf/types/known/emptypb"
	"strings"
	"time"
)

type user struct {
}

func User() *user {
	return &user{}
}

func (u user) Login(ctx context.Context, req *userpb.LoginReq) (*userpb.LoginRes, error) {
	options := opt.NewOpts()
	if strings.Contains(req.Name, "@") {
		options = options.Where("email = ?", req.Name)
	} else {
		options = options.Where("phone = ?", req.Name)
	}
	user, err := repo.User().Get(ctx, options...)
	if user == nil {
		return nil, errno.New400("账号不存在")
	}
	if err != nil {
		return nil, err
	}

	if user.State == entity.UserStateRegister {
		return nil, errno.New400("账号未激活")
	}

	security, err := repo.UserSecurity().Get(ctx, opt.Where("user_id = ?", user.ID))
	if err != nil {
		return nil, err
	}
	if security == nil {
		return nil, errno.New400("账号不存在")
	}
	if security.Password != cryption.Md5Str(req.Password+security.Salt) {
		return nil, errno.New400("密码错误")
	}
	token := cryption.UUID()
	target := &entity.User{}
	err = copier.Copy(target, user)
	if err != nil {
		return nil, err
	}
	data, err := json.Marshal(target)
	if err != nil {
		return nil, err
	}
	err = redisx.Client().Set(keys.TokenKey(token), data, time.Hour*24*3).Err()
	if err != nil {
		return nil, err
	}
	return &userpb.LoginRes{Token: token}, nil
}

func (u user) Register(ctx context.Context, req *userpb.RegisterReq) (*emptypb.Empty, error) {
	err := ormx.Transaction(ctx, func(ctx context.Context) error {

		salt := cryption.UUID()
		user := &entity.User{
			Name:   req.Name,
			Email:  req.Email,
			State:  1,
			Avatar: "https://kkako-blog-bucket.oss-cn-beijing.aliyuncs.com/avatar/default_avatar.gif",
		}
		err := repo.User().Add(ctx, user)
		if err != nil {
			e := &mysql.MySQLError{}
			if errors.As(err, &e) {
				// 唯一index冲突
				if e.Number == 1062 {
					return errs.New(400, "", "该邮箱已经注册")
				}
			}
			return err
		}
		security := &entity.UserSecurity{
			UserId:   user.ID,
			Password: cryption.Md5Str(req.Password + salt),
			Salt:     salt,
		}
		err = repo.UserSecurity().Add(ctx, security)
		if err != nil {
			return err
		}

		// 发送激活邮件
		code := uuid.NewV4().String()
		_, err = redisx.Client().Set(keys.UserActive(user.ID), code, time.Hour*24*3).Result()
		if err != nil {
			return err
		}
		err = emailx.Send(req.Email, "感谢注册", fmt.Sprintf(html, viper.GetString("app.addr")+fmt.Sprintf("/user/active?code=%s&user_id=%d", code, user.ID)))
		if err != nil {
			return err
		}

		event := &common.Event{
			EventType:     common.EventType_UserRegister,
			TargetId:      user.ID,
			TargetType:    0,
			ActorId:       user.ID,
			TargetOwnerId: user.ID,
		}
		eventData, err := proto.Marshal(event)
		if err != nil {
			return err
		}

		// 发送注册事件 初始化
		return producer.Send(&sarama.ProducerMessage{
			Value: sarama.ByteEncoder(eventData),
			Topic: "video-rpc-event",
		})
	})
	if err != nil {
		return nil, err
	}
	return &emptypb.Empty{}, nil
}

var html = `
<!DOCTYPE html>
<html lang="en">

<head>
    <meta charset="UTF-8"/>
</head>

<body>
    <p>感谢注册,点击<a href="%s">此处</a>激活账户</p>
</body>

</html>
`

func (u user) UserInfo(ctx context.Context, id *userpb.ID) (*userpb.UserInfoRes, error) {
	//TODO implement me
	panic("implement me")
}
