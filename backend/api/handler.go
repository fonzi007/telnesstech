package api

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"telnesstech/db"
	"telnesstech/model"
)

type Handler struct {
	storage   db.Storage
	validator Validator
}

func NewHandler(storage db.Storage) (*Handler, error) {
	return &Handler{
		storage:   storage,
		validator: Validator{},
	}, nil
}

func (h *Handler) Search(ctx *gin.Context) {
	var request *model.PhoneNumberSearchRequest
	if err := ctx.BindJSON(&request); err != nil {
		ctx.AbortWithError(http.StatusBadRequest, err)
	}

	err := h.validator.validateFilterRequest(request)
	if err != nil {
		ctx.AbortWithError(http.StatusBadRequest, err)
	}

	phoneNumbers, err := h.storage.FindBy(*request)
	if err != nil {
		ctx.AbortWithError(http.StatusInternalServerError, err)
	}

	ctx.IndentedJSON(http.StatusOK, &model.PhoneNumberListResponse{
		PhoneNumbers: mapToResponse(phoneNumbers),
	})
}
