package ginweb

import (
	"context"
	"net/http"

	"github.com/gin-gonic/gin"
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
		if err != nil {
			handleError(gctx, err)
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
