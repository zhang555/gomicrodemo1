package main

import (
	"gomicrodemo1/article-service/handler"
	"gomicrodemo1/common/config"
	"gomicrodemo1/common/db"
	"gomicrodemo1/pb/pb"
	"log"

	"github.com/micro/go-micro"
)

func main() {

	defer db.DB.Close()

	s := micro.NewService(
		micro.Name(config.SERVICE_ARTICLE),
		micro.Version("latest"),
		//micro.WrapHandler(wrapper.Wrapper1),
		//micro.WrapHandler(wrapper.AuthWrapper),
	)

	s.Init()

	pb.RegisterArticleServiceHandler(s.Server(), &handler.Handler{})

	if err := s.Run(); err != nil {
		log.Fatalf("service error: %v\n", err)
	}

}
