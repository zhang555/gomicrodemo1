package handler

import (
	"common/config"
	"common/model"
	"context"
	pb "message-service/proto/message"
	"net/http"
	"time"
)

type Handler struct {
	//Publisher micro.Publisher
}

func (h *Handler) GetMany(ctx context.Context, req *pb.PageNumPageSize, resp *pb.MessageListResponse) error {

	var (
		beans []*model.Message
		count int32
	)

	err := DB.Order("create_time desc").Offset((req.PageNum - 1) * req.PageSize).Limit(req.PageSize).Find(&beans).Error
	if err != nil {
		Log.Error(err)
		resp.Code = http.StatusInternalServerError
		resp.Msg = err.Error()
		return err
	}

	for _, value := range beans {
		resp.Messages = append(resp.Messages, &pb.Message{
			Id:         value.ID,
			UserId:     value.UserId,
			Message:    value.Message,
			CreateTime: time.Time(value.CreateTime).Format(config.Time2),
			UpdateTime: time.Time(value.UpdateTime).Format(config.Time2),
		})
	}

	DB.Find(&beans).Count(&count)
	resp.Page = &pb.PageMap{
		PageNum:  req.PageNum,
		PageSize: req.PageSize,
		Count:    count,
	}

	resp.Code = http.StatusOK

	return nil
}
