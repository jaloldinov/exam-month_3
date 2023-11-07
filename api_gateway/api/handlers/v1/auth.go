package v1

import (
	"api_gateway/api/models"
	"api_gateway/config"
	"api_gateway/genproto/user_service"
	"api_gateway/pkg/helper"
	"api_gateway/pkg/logger"
	"fmt"

	"net/http"

	"github.com/gin-gonic/gin"
)

// @Router       /login [post]
// @Summary      create person
// @Description  api for create persons
// @Tags         persons
// @Accept       json
// @Produce      json
// @Param        person    body     models.LoginReq  true  "data of person"
// @Success      200  {object}  models.LoginRes
// @Failure      400  {object}  response.ErrorResp
// @Failure      404  {object}  response.ErrorResp
// @Failure      500  {object}  response.ErrorResp
func (h *handlerV1) SignIn(c *gin.Context) {

	var login models.LoginReq

	err := c.ShouldBindJSON(&login)

	if err != nil {
		h.log.Error("error while binding:", logger.Error(err))
		res := models.ErrorResp{Code: "BAD REQUEST", Message: "invalid fields in body"}
		c.JSON(http.StatusBadRequest, res)
		return
	}

	hashPass, err := helper.GeneratePasswordHash(login.Password)

	if err != nil {
		h.log.Error("error while binding:", logger.Error(err))
		res := models.ErrorResp{Code: "INVALID Password", Message: "invalid password"}
		c.JSON(http.StatusBadRequest, res)
		return
	}

	if login.Role == "courier" {
		resp, err := h.services.Courier().GetByLogin(c.Request.Context(), &user_service.IdRequest{
			Id: login.Login,
		})

		if err != nil {
			fmt.Println("error Staff GetByLoging:", err.Error())
			res := models.ErrorResp{Code: "INTERNAL ERROR", Message: "internal server error"}
			c.JSON(http.StatusInternalServerError, res)
			return
		}

		err = helper.ComparePasswords([]byte(hashPass), []byte(resp.Password))

		if err != nil {
			h.log.Error("error while binding:", logger.Error(err))
			res := models.ErrorResp{Code: "INVALID Password", Message: "invalid password"}
			c.JSON(http.StatusBadRequest, res)
			return
		}

		m := make(map[string]interface{})
		m["user_id"] = resp.Id
		token, err := helper.GenerateJWT(m, config.TokenExpireTime, config.JWTSecretKey)

		if err != nil {
			return
		}

		c.JSON(http.StatusCreated, models.LoginRes{Token: token})

		return

	} else if login.Role == "user" {

		resp, err := h.services.User().GetByLogin(c.Request.Context(), &user_service.IdRequest{
			Id: login.Login,
		})

		if err != nil {
			fmt.Println("error Staff GetByLoging:", err.Error())
			res := models.ErrorResp{Code: "INTERNAL ERROR", Message: "internal server error"}
			c.JSON(http.StatusInternalServerError, res)
			return
		}

		err = helper.ComparePasswords([]byte(hashPass), []byte(resp.Password))

		if err != nil {
			h.log.Error("error while binding:", logger.Error(err))
			res := models.ErrorResp{Code: "INVALID Password", Message: "invalid password"}
			c.JSON(http.StatusBadRequest, res)
			return
		}

		m := make(map[string]interface{})
		m["user_id"] = resp.Id
		token, err := helper.GenerateJWT(m, config.TokenExpireTime, config.JWTSecretKey)

		if err != nil {
			return
		}

		c.JSON(http.StatusCreated, models.LoginRes{Token: token})

		return

	} else {
		fmt.Println("Role undefined")
		res := models.ErrorResp{Code: "INVALID Role", Message: "invalid role"}
		c.JSON(http.StatusBadRequest, res)
		return
	}

}
