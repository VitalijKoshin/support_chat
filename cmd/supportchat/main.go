package main

import (
	"supportchat/internal/apperrors"
	"supportchat/internal/config"
	"supportchat/internal/controllers"
	"supportchat/internal/infrastructure/database"
	"supportchat/internal/infrastructure/logger"
	sessionrepository "supportchat/internal/repository/session_repository"
	repository "supportchat/internal/repository/user_repository"
	"supportchat/internal/routes"
	userusecase "supportchat/internal/usecase/user_usecase"

	"github.com/labstack/echo/v4"
)

var envPath = "configs/.env"

func main() {
	logger := logger.NewLogger()
	cfg, err := config.NewConfig(envPath)
	if err != nil {
		logger.Fatal(err)
	}

	db, err := database.NewMySQLDB(cfg.MySQL)
	if err != nil {
		logger.Fatal(err)
	}

	redis, err := database.NewRedisDB(cfg.Redis)
	if err != nil {
		logger.Fatal(err)
	}

	memory := database.NewMemoryDB()

	dbS := database.NewDatabase(redis, db, memory)
	userRepoMysql, err := repository.NewUserRepository(repository.UserRepositoryTypes.MySQL, dbS)
	if err != nil {
		logger.Fatal(err)
	}

	userRepoRedis, err := repository.NewUserRepository(repository.UserRepositoryTypes.Redis, dbS)
	if err != nil {
		logger.Fatal(err)
	}

	sessionRepoMysql, err := sessionrepository.NewSessionRepository(sessionrepository.SessionRepositoryTypes.MySQL, cfg, dbS)
	if err != nil {
		logger.Fatal(err)
	}

	sessionRepoRedis, err := sessionrepository.NewSessionRepository(sessionrepository.SessionRepositoryTypes.Redis, cfg, dbS)
	if err != nil {
		logger.Fatal(err)
	}

	e := echo.New()

	apiUseCase := userusecase.NewAuthUsecase(userRepoMysql, userRepoRedis, sessionRepoMysql, sessionRepoRedis)
	userUseCase := userusecase.NewUserUsecase(userRepoMysql, userRepoRedis)

	userController := controllers.NewUserController(userUseCase, cfg)
	apiController := controllers.NewApiController(userUseCase, apiUseCase, cfg)
	chatController := controllers.NewChatController(userUseCase, cfg)

	e = routes.NewRoutes(e, userController, apiController, chatController, cfg)

	err = e.Start(cfg.Port)
	if err != nil {
		logger.Fatal(apperrors.ServerStartError.AppendMessage(err))
	}
}
