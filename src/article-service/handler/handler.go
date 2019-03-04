package handler

import (
	pb "article-service/proto/article"
	"common/config"
	"common/config/respcode"
	"common/model"
	"context"
	"net/http"
	"time"
	userPb "user-service/proto/user"
)

type Handler struct {
	//Publisher micro.Publisher
}

func (h *Handler) GetOne(ctx context.Context, req *pb.Id, resp *pb.ArticleResponse) error {
	var (
		article model.Article
	)

	article.ID = req.Id
	if err := DB.First(&article).Error; err != nil {
		Log.Error(err)
		resp.Code = http.StatusInternalServerError
		resp.Msg = err.Error()
		return err
	}

	resp.Article = &pb.Article{
		Id:      article.ID,
		Title:   article.Title,
		Content: article.Content,
	}
	resp.Code = http.StatusOK

	return nil

}

func (h *Handler) GetMany(ctx context.Context, req *pb.PageNumPageSize, resp *pb.ArticleListResponse) error {

	var (
		articles []*model.Article
		count    int32
	)

	//sql := SqlArticle()

	err := DB.Order("create_time desc ").Offset((req.PageNum - 1) * req.PageSize).Limit(req.PageSize).Find(&articles).Error
	//err := DB.Offset((req.PageNum - 1) * req.PageSize).Limit(req.PageSize).Raw(sql).Scan(&articles).Error
	if err != nil {
		Log.Error(err)
		resp.Code = http.StatusInternalServerError
		resp.Msg = err.Error()
		return err
	}

	for _, value := range articles {
		resp.Articles = append(resp.Articles, &pb.Article{
			Id:         value.ID,
			Title:      value.Title,
			Content:    value.Content,
			UserId:     value.UserId,
			CreateTime: time.Time(value.CreateTime).Format(config.Time2),
			UpdateTime: time.Time(value.UpdateTime).Format(config.Time2),
		})
	}

	//var timestamp1 = timestamp.Timestamp{
	//	Seconds: 123123123,
	//	Nanos:   0,
	//}
	//fmt.Println("timestamp ", timestamp1.String())

	DB.Find(&articles).Count(&count)
	resp.Page = &pb.PageMap{
		PageNum:  req.PageNum,
		PageSize: req.PageSize,
		Count:    count,
	}

	resp.Code = http.StatusOK

	return nil
}

func (h *Handler) Post(ctx context.Context, req *pb.Article, resp *pb.IdResponse) error {
	var (
		article model.Article
	)

	//uuid1, err := uuid.NewV4()
	//if err != nil {
	//	Log.Error(err)
	//	resp.Code = http.StatusInternalServerError
	//	resp.Msg = err.Error()
	//	return err
	//}

	value := ctx.Value("userSession")

	user, ok := value.(*userPb.User)
	Log.Info(user, ok)
	if !ok || user.Id == 0 {
		resp.Code = respcode.Unauth.Code
		resp.Msg = respcode.Unauth.Msg
		return nil
	}

	//if user.Id == 0 {
	//	resp.Code = respcode.Unauth.Code
	//	resp.Msg = respcode.Unauth.Msg
	//	return nil
	//}

	article.ID = 0

	article.Title = req.Title
	article.Content = req.Content
	article.CreateTime = model.JSONTime(time.Now())
	article.UpdateTime = model.JSONTime(time.Now())
	article.UserId = user.Id

	err := DB.Create(&article).Error
	if err != nil {
		Log.Error(err)
		resp.Code = http.StatusInternalServerError
		resp.Msg = err.Error()
		return err
	}

	resp.Code = http.StatusOK
	resp.Id = article.ID

	return nil
}

func (h *Handler) Put(ctx context.Context, req *pb.Article, resp *pb.Response) error {
	var (
		err error
	)

	user := ctx.Value("userSession").(*userPb.User)
	if user.Id == 0 {
		resp.Code = respcode.Unauth.Code
		resp.Msg = respcode.Unauth.Msg
		return nil
	}

	savemap := map[string]interface{}{
		"Title":   req.Title,
		"Content": req.Content,
	}

	err = DB.Model(&model.Article{ID: req.Id}).Updates(savemap).Error
	if err != nil {
		Log.Error(err)
		resp.Code = http.StatusInternalServerError
		resp.Msg = err.Error()
		return err
	}

	resp.Code = http.StatusOK

	return nil
}

func (h *Handler) Delete(ctx context.Context, req *pb.Id, resp *pb.Response) error {

	user := ctx.Value("userSession").(*userPb.User)
	if user.Id == 0 {
		resp.Code = respcode.Unauth.Code
		resp.Msg = respcode.Unauth.Msg
		return nil
	}

	if req.Id == 0 {
		resp.Code = respcode.IdCanNot0.Code
		resp.Msg = respcode.IdCanNot0.Msg
		return nil
	}

	err := DB.Delete(&pb.Article{Id: req.Id}).Error
	if err != nil {
		Log.Error(err)
		resp.Code = http.StatusInternalServerError
		resp.Msg = err.Error()
		return err
	}

	resp.Code = http.StatusOK

	return nil

}
