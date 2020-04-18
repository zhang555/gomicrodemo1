package handler

import (
	"context"
	"gomicrodemo1/common/model"
	"gomicrodemo1/pb/pb"
	"net/http"

	"github.com/micro/go-micro"

	"gomicrodemo1/common/config"
	"gomicrodemo1/common/config/respcode"
	"time"

	uuid "github.com/satori/go.uuid"
)

type Handler struct {
	Publisher micro.Publisher
}

func (h *Handler) GetOne(ctx context.Context, req *pb.UserId, resp *pb.UserResponse) error {

	u := &pb.User{Id: req.Id}
	//u.Id = req.Id
	if err := DB.First(&u).Error; err != nil {
		Log.Error(err)
		resp.Code = http.StatusInternalServerError
		resp.Msg = err.Error()
		return err
	}
	resp.User = u
	resp.Code = http.StatusOK

	return nil

}

func (h *Handler) GetMany(ctx context.Context, req *pb.PageNumPageSize, resp *pb.UserListResponse) error {

	var (
		users []*model.User
		count int32
	)

	err := DB.Order("create_time desc").Offset((req.PageNum - 1) * req.PageSize).Limit(req.PageSize).Find(&users).Error
	if err != nil {
		Log.Error(err)
		resp.Code = http.StatusInternalServerError
		resp.Msg = err.Error()
		return err
	}
	//resp.Result = users

	//Log.Info(users)

	//pbusers = &pb.Users{Users: users}

	//any, err := ptypes.MarshalAny(pbusers)
	//if err != nil {
	//	Log.Error(err)
	//	return err
	//}

	for _, value := range users {
		resp.Users = append(resp.Users, &pb.User{
			Id:         value.ID,
			Username:   value.Username,
			Password:   value.Password,
			Nickname:   value.Nickname,
			CreateTime: time.Time(value.CreateTime).Format(config.Time2),
			UpdateTime: time.Time(value.UpdateTime).Format(config.Time2),
		})
	}

	DB.Find(&users).Count(&count)
	resp.Page = &pb.PageMap{
		PageNum:  req.PageNum,
		PageSize: req.PageSize,
		Count:    count,
	}
	//resp.Resp.Code = 200
	//resp.Resp.Msg = ""

	resp.Code = http.StatusOK

	//resp.Code = 200

	//log.Println(DB)
	//log.Println(users)
	//log.Println(count)

	return nil
}

func (h *Handler) Post(ctx context.Context, req *pb.User, resp *pb.IdResponse) error {

	var (
		user model.User
	)

	user.ID = 0

	user.Username = req.Username
	user.Password = req.Password
	user.Nickname = req.Nickname
	user.CreateTime = model.JSONTime(time.Now())
	user.UpdateTime = model.JSONTime(time.Now())

	err := DB.Create(&user).Error
	if err != nil {
		Log.Error(err)
		resp.Code = http.StatusInternalServerError
		resp.Msg = err.Error()
		return err
	}

	req.Id = user.ID

	if err := h.Publisher.Publish(ctx, req); err != nil {
		Log.Error(err)
		resp.Code = http.StatusInternalServerError
		resp.Msg = err.Error()
		return err
	}

	resp.Code = http.StatusOK
	resp.Id = user.ID

	return nil
}

func (h *Handler) Put(ctx context.Context, req *pb.User, resp *pb.Response) error {

	var (
		err error
	)

	savemap := map[string]interface{}{
		"Username": req.Username,
		"Password": req.Password,
		"Nickname": req.Nickname,
	}

	err = DB.Model(&model.User{ID: req.Id}).Updates(savemap).Error
	if err != nil {
		Log.Error(err)
		resp.Code = http.StatusInternalServerError
		resp.Msg = err.Error()
		return err
	}

	resp.Code = http.StatusOK

	return nil
}

func (h *Handler) Delete(ctx context.Context, req *pb.UserId, resp *pb.Response) error {

	if req.Id == 0 {
		resp.Code = 200011
		resp.Msg = "id不能为零"
		return nil
	}

	err := DB.Delete(&pb.User{Id: req.Id}).Error
	if err != nil {
		Log.Error(err)
		resp.Code = http.StatusInternalServerError
		resp.Msg = err.Error()
		return err
	}

	resp.Code = http.StatusOK

	return nil
}

func (h *Handler) Auth(ctx context.Context, req *pb.User, resp *pb.Token) error {

	var (
		userLoginHistory model.UserLoginHistory
	)

	userId := Login(model.User{Username: req.Username, Password: req.Password})

	if userId == 0 {
		resp.Code = respcode.UsernamePasswordError.Code
		resp.Msg = respcode.UsernamePasswordError.Msg
		return nil
	}

	uuid1 := uuid.NewV4()
	resp.Token = uuid1.String()
	userLoginHistory.Token = uuid1.String()

	userLoginHistory.ID = 0
	userLoginHistory.Duration = 3600 * 24
	userLoginHistory.UserId = userId

	err := DB.Create(&userLoginHistory).Error
	if err != nil {
		Log.Error(err)
		resp.Code = http.StatusInternalServerError
		resp.Msg = err.Error()
		return err
	}

	resp.Code = http.StatusOK
	return nil
}

func (h *Handler) ValidateToken(ctx context.Context, req *pb.Token, resp *pb.Token) error {
	var (
		userLoginHistory model.UserLoginHistory
	)

	DB.Where(" token = ? ").First(&userLoginHistory)

	if userLoginHistory.ID == 0 {
		resp.Valid = false
		resp.Code = http.StatusOK
		return nil
	}

	if userLoginHistory.Logout != 0 {
		resp.Valid = false
		resp.Code = http.StatusOK
		return nil
	}

	time.Now().Add(time.Duration(userLoginHistory.Duration))

	if time.Time(userLoginHistory.CreateTime).Add(time.Duration(userLoginHistory.Duration)).After(time.Now()) {
		resp.Valid = false
		resp.Code = http.StatusOK
		return nil
	}

	resp.Valid = true
	resp.Code = http.StatusOK
	return nil
}

func (h *Handler) Logout(ctx context.Context, req *pb.Token, resp *pb.Response) error {
	var (
		userLoginHistory model.UserLoginHistory
	)
	userLoginHistory.Logout = 1

	err := DB.Model(&userLoginHistory).Where(" token = ? ", req.Token).Updates(&userLoginHistory).Error
	if err != nil {
		Log.Error(err)
		resp.Code = http.StatusInternalServerError
		resp.Msg = err.Error()
		return err
	}

	resp.Code = http.StatusOK
	return nil
}

func (h *Handler) GetUserSessionByToken(ctx context.Context, req *pb.Token, resp *pb.UserResponse) error {
	var (
	//userLoginHistory model.UserLoginHistory
	//user model.User
	)
	//userLoginHistory.Logout = 1

	user, err := GetUserSessionByToken(req.Token)
	if err != nil {
		Log.Error(err)
		resp.Code = http.StatusInternalServerError
		resp.Msg = err.Error()
		return err
	}

	Log.Info(req.Token)
	Log.Info(user)

	//err := DB.Model(&user).Where(" token = ? ", req.Token).First(&user).Error
	//if err != nil {
	//	Log.Error(err)
	//	resp.Code = http.StatusInternalServerError
	//	resp.Msg = err.Error()
	//	return err
	//}

	resp.Code = http.StatusOK
	resp.User = &pb.User{
		Id:       user.ID,
		Username: user.Username,
		Nickname: user.Nickname,
		//CreateTime: user.CreateTime,
	}
	return nil
}

func (h *Handler) GetNicknamesByIds(ctx context.Context, req *pb.Ids, resp *pb.UserListResponse) error {
	var ()

	users, err := GetUsernamesByIds(req.Ids)
	if err != nil {
		Log.Error(err)
		resp.Code = http.StatusInternalServerError
		resp.Msg = err.Error()
		return err
	}

	for _, value := range users {
		resp.Users = append(resp.Users, &pb.User{
			Id:       value.ID,
			Nickname: value.Nickname,
		})
	}

	resp.Code = http.StatusOK

	return nil
}
