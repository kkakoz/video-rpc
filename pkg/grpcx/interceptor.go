package grpcx

import (
	"errors"
	"fmt"
	grpc_recovery "github.com/grpc-ecosystem/go-grpc-middleware/recovery"
)

func RecoveryInterceptor() grpc_recovery.Option {
	return grpc_recovery.WithRecoveryHandler(func(p interface{}) (err error) {
		err, ok := p.(error)
		if ok {
			return errors.New(err.Error())
		}
		return errors.New(fmt.Sprintf("%+v", p))
	})
}
