package ginweb

import (
	"context"
	"errors"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"github.com/mkaiho/go-aws-sandbox/adapter"
	"github.com/mkaiho/go-aws-sandbox/usecase"
)

type (
	RegisterUserInput struct {
		Name  string `json:"name" binding:"required"`
		Email string `json:"email" binding:"required"`
	}
	RegisterUserOutput struct {
		ID    string `json:"id"`
		Name  string `json:"name"`
		Email string `json:"email"`
	}
)

type RegisterUserController struct {
	txm                    adapter.TxManager[RegisterUserOutput]
	registerUserInteractor usecase.RegisterUserInteractor
}

func NewRegisterUserController(
	txm *adapter.TxManager[RegisterUserOutput],
	registerUserInteractor usecase.RegisterUserInteractor,
) *RegisterUserController {
	return &RegisterUserController{
		txm:                    *txm,
		registerUserInteractor: registerUserInteractor,
	}
}

func (c *RegisterUserController) Handle(gctx *gin.Context) {
	var ctx context.Context
	var err error
	var resp RegisterUserOutput
	defer func() {
		code := http.StatusOK
		if err != nil {
			code = http.StatusInternalServerError
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
		gctx.PureJSON(http.StatusCreated, resp)
	}()

	var input RegisterUserInput
	err = gctx.ShouldBindJSON(&input)
	if err != nil {
		return
	}

	ctx, err = c.txm.ContextWithNewTx(context.Background())
	if err != nil {
		return
	}

	var regOutput *usecase.UserRegisterOutput
	regOutput, err = c.registerUserInteractor.ExecuteInTx(ctx, usecase.UserRegisterInput{
		Name:  input.Name,
		Email: input.Email,
	})
	if err != nil {
		return
	}
	resp = RegisterUserOutput{
		ID:    regOutput.RegisteredUserDetail.ID.String(),
		Name:  regOutput.RegisteredUserDetail.Name,
		Email: regOutput.RegisteredUserDetail.Email,
	}
}
