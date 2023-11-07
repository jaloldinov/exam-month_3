package v1

import (
	"delever/api_gateway/config"
	"delever/api_gateway/genproto/user_service"
	"delever/api_gateway/pkg/helper"
	"delever/api_gateway/pkg/logger"
	"net/http"

	"github.com/gin-gonic/gin"
)

type SignInReq struct {
	Username string `json:"login"`
	Password string `json:"password"`
	Role     string `json:"role"`
}

type SignInResp struct {
	Token string `json:"token"`
}

func (h *handlerV1) SignIn(ctx *gin.Context) {
	var req SignInReq
	if err := ctx.ShouldBindJSON(&req); err != nil {
		h.log.Error("error while binding:", logger.Error(err))
		h.handlerResponse(ctx, "ShouldBindJSON()", http.StatusBadRequest, err.Error())
		return
	}

	if req.Role == "user" {
		user, err := h.services.User().GetByUsername(ctx, &user_service.UserGetByUsernameReq{
			Username: req.Username,
		})
		if err != nil {
			h.log.Error("error while getting user by username:", logger.Error(err))
			h.handlerResponse(ctx, "UserService().GetByUsername", http.StatusBadRequest, err.Error())
			return
		}

		if err = helper.ComparePasswords([]byte(user.Password), []byte(req.Password)); err != nil {
			h.log.Error("username or password incorrect")
			h.handlerResponse(ctx, "username or password incorrect", http.StatusBadRequest, "username or password incorrect")
			return
		}

		m := make(map[string]interface{})
		m["user_id"] = user.Id
		m["branch_id"] = user.BranchId
		m["username"] = user.Login

		token, err := helper.GenerateJWT(m, config.TokenExpireTime, config.JWTSecretKey)
		if err != nil {
			h.log.Error("error while generate jwt token", logger.Error(err))
			h.handlerResponse(ctx, "error while generate jwt token", http.StatusBadRequest, err.Error())
			return
		}

		ctx.JSON(http.StatusCreated, SignInResp{Token: token})
	} else {
		courier, err := h.services.Courier().GetByUsername(ctx, &user_service.CourierGetByUsernameReq{
			Username: req.Username,
		})
		if err != nil {
			h.log.Error("error while getting courier by username:", logger.Error(err))
			h.handlerResponse(ctx, "CourierService().GetByUsername", http.StatusBadRequest, err.Error())
			return
		}

		if err = helper.ComparePasswords([]byte(courier.Password), []byte(req.Password)); err != nil {
			h.log.Error("username or password incorrect")
			h.handlerResponse(ctx, "username or password incorrect", http.StatusBadRequest, err.Error())
			return
		}

		m := make(map[string]interface{})
		m["courier_id"] = courier.Id
		m["branch_id"] = courier.BranchId
		m["username"] = courier.Login

		token, err := helper.GenerateJWT(m, config.TokenExpireTime, config.JWTSecretKey)
		if err != nil {
			h.log.Error("error while generate jwt token", logger.Error(err))
			h.handlerResponse(ctx, "error while generate jwt token", http.StatusBadRequest, err.Error())
			return
		}

		ctx.JSON(http.StatusCreated, SignInResp{Token: token})
	}
}
