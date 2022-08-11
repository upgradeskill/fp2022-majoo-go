package cmd

import (
	"context"
	"mini-pos/cmd/util/gorm"
	"mini-pos/configs"
	outletService "mini-pos/core/outlet/service"
	userService "mini-pos/core/user/service"
	outletHandler "mini-pos/handlers/api/outlet"
	userHandler "mini-pos/handlers/api/user"
	utilHandler "mini-pos/handlers/util"
	outletRepository "mini-pos/repositories/outlet"
	outletUserRepository "mini-pos/repositories/outletUser"
	userRepository "mini-pos/repositories/user"
	"mini-pos/util/logger"
	"os"
	"os/signal"
	"time"

	"github.com/go-playground/validator"
	"github.com/spf13/cobra"

	"github.com/labstack/echo/v4"
	"github.com/sirupsen/logrus"
)

const LOG_IDENTIFIER = "APP_MAIN"

var log = logger.SetIdentifierField(LOG_IDENTIFIER)

func CustomValidator() *validator.Validate {
	customValidator := validator.New()
	return customValidator
}

func init() {
	rootCmd.AddCommand(apiCmd)
}

var apiCmd = &cobra.Command{
	Use:   "api",
	Short: "Start api service",
	Run: func(cmd *cobra.Command, args []string) {
		logrus.SetFormatter(&logrus.JSONFormatter{
			FieldMap: logrus.FieldMap{
				logrus.FieldKeyTime: "timestamp",
				logrus.FieldKeyMsg:  "message",
			},
		})
		logrus.SetLevel(logrus.InfoLevel)
		logrus.SetOutput(os.Stderr)
		dbCon, err := configs.ConnectDBMySql()
		if err != nil {
			log.Fatal(err)
			panic(err)
		}

		gormDb, err := gorm.InitGorm(dbCon)
		if err != nil {
			dbCon.Close()
			log.Fatal(err)
			panic(err)
		}

		user, err := userRepository.New(gormDb)
		if err != nil {
			log.Error(err)
		}
		userServices := userService.New(user)
		userHandlers := userHandler.New(userServices)

		outletUser, err := outletUserRepository.New(gormDb)
		if err != nil {
			log.Error(err)
		}

		outlet, err := outletRepository.New(gormDb)
		if err != nil {
			log.Error(err)
		}
		outletServices := outletService.New(outletUser, outlet)
		outletHandlers := outletHandler.New(outletServices)

		e := echo.New()
		e.HideBanner = true
		e.HidePort = true
		e.Validator = &utilHandler.BodyRequestValidator{Validator: CustomValidator()}

		userHandler.RegisterRouter(e, userHandlers)
		outletHandler.RegisterRouter(e, outletHandlers)

		go func() {
			if err := e.Start(":2022"); err != nil {
				dbCon.Close()
				log.Fatal(err)
				panic(err)
			}
		}()

		quit := make(chan os.Signal, 1)
		signal.Notify(quit, os.Interrupt)
		<-quit

		ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
		defer cancel()

		if err := e.Shutdown(ctx); err != nil {
			dbCon.Close()
			log.Fatal(err)
			panic(err)
		}
	},
}
