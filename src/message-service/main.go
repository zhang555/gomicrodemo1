package main

import (
	"common/config"
	"common/db"
	"log"
	"message-service/handler"
	pb "message-service/proto/message"
	"message-service/subscriber"
	"time"

	"github.com/micro/go-micro"
)

func main() {
	defer db.DB.Close()
	run2()
}

func run2() {
	//repo := &UserRepository{db}

	s := micro.NewService(
		micro.Name(config.SERVICE_MESSAGE),
		micro.Version("latest"),
		micro.RegisterInterval(time.Millisecond*1000),
		micro.RegisterTTL(time.Millisecond*1000),
	)

	s.Init()

	micro.RegisterSubscriber(config.TOPIC_USER_CREATED, s.Server(), new(subscriber.Subscriber))

	pb.RegisterMessageServiceHandler(s.Server(), &handler.Handler{})

	if err := s.Run(); err != nil {
		log.Fatalf("service : %s error: %v\n", config.SERVICE_MESSAGE, err)
	}

}
