package main

import (
	"gomicrodemo1/common/config"
	"gomicrodemo1/common/db"
	"gomicrodemo1/pb/pb"
	"gomicrodemo1/user-service/handler"
	"log"
	"time"

	"github.com/micro/go-micro"
)

func main() {
	defer db.DB.Close()
	run2()
}

func run2() {
	s := micro.NewService(
		micro.Name(config.SERVICE_USER),
		micro.Version("latest"),
		micro.RegisterInterval(time.Millisecond*1000),
		micro.RegisterTTL(time.Millisecond*1000),
	)

	s.Init()
	publisher := micro.NewPublisher(config.TOPIC_USER_CREATED, s.Client())

	pb.RegisterUserServiceHandler(s.Server(), &handler.Handler{publisher})

	if err := s.Run(); err != nil {
		log.Fatalf("user service error: %v\n", err)
	}

}
