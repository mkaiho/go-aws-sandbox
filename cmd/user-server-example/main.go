package main

import (
	"fmt"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/mkaiho/go-aws-sandbox/adapter"
	"github.com/mkaiho/go-aws-sandbox/adapter/id"
	"github.com/mkaiho/go-aws-sandbox/adapter/rdb"
	"github.com/mkaiho/go-aws-sandbox/controller/ginweb"
	"github.com/mkaiho/go-aws-sandbox/infrastructure"
	"github.com/mkaiho/go-aws-sandbox/usecase"
	"github.com/mkaiho/go-ecs-batch-sample/logging"
	"github.com/spf13/cobra"
)

var (
	initErr    error
	command    *cobra.Command
	db         rdb.DB
	controller ginweb.RegisterUserController
)

func init() {
	logging.InitLoggerWithZap()
	di()
	command = newCommand()
	command.Flags().IntP("port", "", 3000, "listening port")
	command.Flags().StringP("host", "", "", "host name")
}

func main() {
	if err := command.Execute(); err != nil {
		logging.GetLogger().Error(err, "error has occured")
		os.Exit(1)
	} else {
		logging.GetLogger().Info("completed")
	}
}

func newCommand() *cobra.Command {
	command := cobra.Command{
		Use:           "echo args...",
		Short:         "display args",
		Long:          "Display arguments on stdout.",
		RunE:          handle,
		SilenceUsage:  true,
		SilenceErrors: true,
	}

	return &command
}

func di() {
	// adapter
	var (
		idm           id.IDManager
		userIDManager *adapter.UserIDManager
		userGateway   *adapter.UserGateway
	)
	{
		var conf rdb.Config
		conf, initErr = infrastructure.LoadMySQLConfig()
		if initErr != nil {
			return
		}
		db, initErr = infrastructure.OpenRDB(conf)
		if initErr != nil {
			return
		}
		idm = infrastructure.NewULIDManager()
		userIDManager = adapter.NewUserIDManager(idm)
		userGateway = adapter.NewUserGateway(userIDManager)
	}
	// interactor
	var (
		userInteractor usecase.RegisterUserInteractor
	)
	{
		userInteractor = usecase.NewRegisterUserInteractor(
			adapter.NewTxManager[usecase.UserRegisterOutput](db),
			userIDManager,
			userGateway,
		)
	}
	// controller
	{
		controller = *ginweb.NewRegisterUserController(
			adapter.NewTxManager[ginweb.RegisterUserOutput](db),
			userInteractor,
		)
	}
}

func handle(cmd *cobra.Command, args []string) (err error) {
	var (
		host string
		port int
	)
	defer func() {
		if pErr := recover(); pErr != nil {
			err = pErr.(error)
			return
		}
	}()
	if initErr != nil {
		return initErr
	}

	host, err = cmd.Flags().GetString("host")
	if err != nil {
		return err
	}
	port, err = cmd.Flags().GetInt("port")
	if err != nil {
		return err
	}

	server := gin.Default()
	handler := controller.Handle
	server.POST("/users", handler)

	return server.Run(fmt.Sprintf("%s:%d", host, port))
}
