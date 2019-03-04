package subscriber

import (
	"common/db"
	"common/log"
	"common/model"
	"context"
	userPb "user-service/proto/user"
)

type Subscriber struct{}

func (sub *Subscriber) Process(ctx context.Context, user *userPb.User) error {
	//log.Println("[Picked up a new message]")
	//log.Println("[Sending email to]:", user.Nickname)

	bean := model.Message{
		ID:      0,
		UserId:  user.Id,
		Message: "You have successfully created an account",
	}

	err := db.DB.Create(&bean).Error
	if err != nil {
		log.Log.Error(err)
		return err
	}

	return nil
}

func (sub *Subscriber) Process2(ctx context.Context, user *userPb.User) error {
	log.Log.Info("process2")
	return nil
}
func (sub *Subscriber) Process3(ctx context.Context, user *userPb.User) error {
	log.Log.Info("process3")
	return nil
}
