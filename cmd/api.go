package cmd

import (
	"context"
	"mini-pos/configs"
	userService "mini-pos/core/user/service"
	userHandler "mini-pos/handlers/api/user"
	utilHandler "mini-pos/handlers/util"
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

		configs.ConnectDB()

		// gormDb, err := gorm.InitGorm(dbCon)
		// if err != nil {
		// 	dbCon.Close()
		// 	log.Fatal(err)
		// 	panic(err)
		// }

		user, err := userRepository.New(gormDb)
		if err != nil {
			log.Error(err)
		}
		userServices := userService.New(user)
		userHandlers := userHandler.New(userServices)

		e := echo.New()
		e.HideBanner = true
		e.HidePort = true
		e.Validator = &utilHandler.BodyRequestValidator{Validator: CustomValidator()}

		userHandler.RegisterRouter(e, userHandlers)

		go func() {
			if err := e.Start(":2022"); err != nil {
				// dbCon.Close()
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
			// dbCon.Close()
			log.Fatal(err)
			panic(err)
		}
	},
}
