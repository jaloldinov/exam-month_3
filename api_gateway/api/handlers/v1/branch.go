package v1

import (
	"context"
	"errors"
	"net/http"

	user_service "api_gateway/genproto/user_service"
	"api_gateway/pkg/util"

	"github.com/gin-gonic/gin"
)

// Branch godoc
// @ID create-user
// @Router /v1/user/create [POST]
// @Summary create user
// @Description create user
// @Tags user
// @Accept json
// @Produce json
// @Param user body user_service.CreateBranchRequest true "user"
// @Success 200 {object} models.ResponseModel{data=user_service.Branch} "desc"
// @Response 400 {object} models.ResponseModel{error=string} "Bad Request"
// @Response 400 {object} models.ResponseModel{error=string} "Bad Request"
// @Failure 500 {object} models.ResponseModel{error=string} "Server Error"
func (h *handlerV1) CreateBranch(c *gin.Context) {
	var user user_service.CreateBranchRequest

	if err := c.BindJSON(&user); err != nil {
		h.handleErrorResponse(c, http.StatusBadRequest, "error while binding json", err)
		return
	}

	resp, err := h.services.Branch().Create(c.Request.Context(), &user)
	if err != nil {
		h.handleErrorResponse(c, http.StatusInternalServerError, "error while creating user", err)
		return
	}

	h.handleSuccessResponse(c, http.StatusCreated, "ok", resp)
}

// GetAllBranch godoc
// @ID get-user
// @Router /v1/user/list [GET]
// @Summary get user all
// @Description get user
// @Tags user
// @Accept json
// @Produce json
// @Param search query string false "search"
// @Param limit query integer false "limit"
// @Param offset query integer false "offset"
// @Success 200 {object} models.ResponseModel{data=user_service.ListBranchResponse} "desc"
// @Response 400 {object} models.ResponseModel{error=string} "Bad Request"
// @Response 400 {object} models.ResponseModel{error=string} "Bad Request"
// @Failure 500 {object} models.ResponseModel{error=string} "Server Error"
func (h *handlerV1) GetAllBranch(c *gin.Context) {
	limit, err := h.ParseQueryParam(c, "limit", "10")
	if err != nil {
		return
	}

	offset, err := h.ParseQueryParam(c, "offset", "0")
	if err != nil {
		return
	}

	resp, err := h.services.Branch().List(
		c.Request.Context(),
		&user_service.ListBranchRequest{
			Limit: int32(limit),
			Page:  int32(offset),
			Name:  c.Query("search"),
		},
	)

	if err != nil {
		h.handleErrorResponse(c, http.StatusInternalServerError, "error getting all users", err)
		return
	}

	h.handleSuccessResponse(c, http.StatusOK, "OK", resp)
}

// Get-BranchByID godoc
// @ID get-user-byID
// @Router /v1/user/get/{user_id} [GET]
// @Summary get user by ID
// @Description get user
// @Tags user
// @Accept json
// @Produce json
// @Param user_id path string true "user_id"
// @Success 200 {object} models.ResponseModel{data=user_service.GetBranchResponse} "desc"
// @Response 400 {object} models.ResponseModel{error=string} "Bad Request"
// @Response 400 {object} models.ResponseModel{error=string} "Bad Request"
// @Failure 500 {object} models.ResponseModel{error=string} "Server Error"
func (h *handlerV1) GetBranch(c *gin.Context) {

	user_id := c.Param("user_id")

	if !util.IsValidUUID(user_id) {
		h.handleErrorResponse(c, http.StatusBadRequest, "user id is not valid", errors.New("user id is not valid"))
		return
	}

	resp, err := h.services.Branch().Get(
		context.Background(),
		&user_service.IdRequest{
			Id: user_id,
		},
	)
	if !handleError(h.log, c, err, "error while getting user") {
		return
	}

	h.handleSuccessResponse(c, http.StatusOK, "ok", resp)

}

// Update Branch godoc
// @ID update_user
// @Router /v1/user/update/{user_id} [PUT]
// @Summary Update Branch
// @Description Update Branch by ID
// @Tags user
// @Accept json
// @Produce json
// @Param        user_id       path    string     true    "Branch ID to update"
// @Param user body user_service.CreateBranchRequest true "user"
// @Success 200 {object} models.ResponseModel{data=models.Status} "desc"
// @Response 400 {object} models.ResponseModel{error=string} "Bad Request"
// @Response 400 {object} models.ResponseModel{error=string} "Bad Request"
// @Failure 500 {object} models.ResponseModel{error=string} "Server Error"
func (h *handlerV1) UpdateBranch(c *gin.Context) {

	var user = user_service.UpdateBranchRequest{}

	// user.Id = c.Param("user_id")
	err := c.ShouldBindJSON(&user)
	if err != nil {
		h.handleErrorResponse(c, http.StatusBadRequest, "error while binding", err.Error())
		return
	}

	resp, err := h.services.Branch().Update(c.Request.Context(), &user)
	if !handleError(h.log, c, err, "error while getting user") {
		return
	}

	h.handleSuccessResponse(c, http.StatusOK, "ok", resp)
}

// / Delete Branch godoc
// @ID delete-user
// @Router /v1/user/delete/{user_id} [DELETE]
// @Summary delete user
// @Description Delete Branch
// @Tags user
// @Accept json
// @Produce json
// @Param user_id path string true "user_id"
// @Success 200 {object} models.ResponseModel{data=models.Status} "desc"
// @Response 400 {object} models.ResponseModel{error=string} "Bad Request"
// @Response 400 {object} models.ResponseModel{error=string} "Bad Request"
// @Failure 500 {object} models.ResponseModel{error=string} "Server Error"
func (h *handlerV1) DeleteBranch(c *gin.Context) {

	var id user_service.IdRequest
	id.Id = c.Param("user_id")

	if !util.IsValidUUID(id.Id) {
		h.handleErrorResponse(c, http.StatusBadRequest, "user id is not valid", errors.New("user id is not valid"))
		return
	}

	resp, err := h.services.Branch().Delete(c.Request.Context(), &id)
	if err != nil {
		h.handleErrorResponse(c, http.StatusBadRequest, "error while deleting user", err.Error())
		return
	}

	h.handleSuccessResponse(c, http.StatusOK, "ok", resp)
}
