package wrapper

import (
	"common/config"
	"context"
	"errors"
	"os"

	"github.com/micro/go-micro/client"
	"github.com/micro/go-micro/metadata"
	"github.com/micro/go-micro/server"

	"common/cerr"
	"common/log"
	userPb "user-service/proto/user"
)

func AuthWrapper(fn server.HandlerFunc) server.HandlerFunc {
	return func(ctx context.Context, req server.Request, resp interface{}) error {
		// consignment-service 独立测试时不进行认证
		if os.Getenv("DISABLE_AUTH") == "true" {
			return fn(ctx, req, resp)
		}
		meta, ok := metadata.FromContext(ctx)
		if !ok {
			return errors.New("no auth meta-data found in request")
		}

		// Note this is now uppercase (not entirely sure why this is...)
		token := meta["X-Csrf-Token"]

		// Auth here
		authClient := userPb.NewUserServiceClient(config.SERVICE_USER, client.DefaultClient)
		authResp, err := authClient.GetUserSessionByToken(context.Background(), &userPb.Token{
			Token: token,
		})
		if err != nil {
			log.Log.Error(err)
			return cerr.ErrorInternalServerError
		}

		//log.Log.Info(authResp)

		//if !authResp.Valid {
		//	context.WithValue(ctx, "userSession", model.User{})
		//}

		ctx = context.WithValue(ctx, "userSession", authResp.User)

		//log.Println("Auth Resp:", authResp)

		err = fn(ctx, req, resp)
		return err
	}
}

func Wrapper1(fn server.HandlerFunc) server.HandlerFunc {
	return func(ctx context.Context, req server.Request, resp interface{}) error {

		log.Log.Info("Wrapper1")

		context.WithValue(ctx, "userSession", "true")

		//http请求头 ，都在meta里面，
		meta, ok := metadata.FromContext(ctx)
		if ok {
			log.Log.Info(meta["X-Csrf-Token"])
			//log.Log.Info(meta["Local"])
			//log.Log.Info(meta["Origin"])

			//for key, value := range meta {
			//	log.Log.Info(key, value)
			//}

		} else {
			log.Log.Info("dont have metadata")

		}
		//log.Log.Info(meta, ok)
		//log.Log.Info(meta, ok)

		//pretty.Println(meta)

		return fn(ctx, req, resp)
	}
}
