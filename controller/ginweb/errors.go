package ginweb

import (
	"errors"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"github.com/mkaiho/go-aws-sandbox/usecase"
)

func handleError(gctx *gin.Context, err error) {
	if err != nil {
		code := http.StatusInternalServerError
		if errors.Is(err, usecase.ErrNotFoundUser) {
			code = http.StatusNotFound
		}
		if errors.Is(err, usecase.ErrDuplicateUser) {
			code = http.StatusConflict
		}
		var vErr validator.ValidationErrors
		if errors.As(err, &vErr) {
			code = http.StatusBadRequest
		}
		payload := gin.H{
			"message": http.StatusText(code),
		}
		if vErr != nil {
			payload["detail"] = vErr.Error()
		}
		gctx.PureJSON(code, payload)
		return
	}
}
