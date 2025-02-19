package apperrors

import (
	"fmt"
	"net/http"
)

type AppError struct {
	Message  string
	Code     string
	HTTPCode int
}

var (
	EnvConfigLoadError = AppError{
		Message:  "Failed to load env file",
		Code:     "ENV_INIT_ERR",
		HTTPCode: http.StatusInternalServerError,
	}

	EnvConfigParseError = AppError{
		Message:  "Failed to parse env file",
		Code:     "ENV_PARSE_ERR",
		HTTPCode: http.StatusInternalServerError,
	}

	EnvConfigMysqlParseError = AppError{
		Message:  "Failed to parse pastgres env file",
		Code:     "ENV_MYSQL_PARSE_ERR",
		HTTPCode: http.StatusInternalServerError,
	}

	EnvConfigRedisParseError = AppError{
		Message:  "Failed to parse redis env file",
		Code:     "ENV_REDIS_PARSE_ERR",
		HTTPCode: http.StatusInternalServerError,
	}

	EnvConfigJwtParseError = AppError{
		Message:  "Failed to parse jwt env file",
		Code:     "ENV_CONFIG_JWT_PARSE_ERROR",
		HTTPCode: http.StatusInternalServerError,
	}

	SqlOpenError = AppError{
		Message:  "Failed to connect database",
		Code:     "SQL_OPEN_ERR",
		HTTPCode: http.StatusInternalServerError,
	}

	PingDBError = AppError{
		Message:  "Failed ping to database",
		Code:     "PING_DB_ERR",
		HTTPCode: http.StatusInternalServerError,
	}

	PingRedisError = AppError{
		Message:  "Failed ping to redis",
		Code:     "PING_REDIS_ERR",
		HTTPCode: http.StatusInternalServerError,
	}

	ServerStartError = AppError{
		Message:  "Failed start app",
		Code:     "SERVER_START_ERR",
		HTTPCode: http.StatusInternalServerError,
	}
)

func (appError *AppError) Error() string {
	return appError.Code + ": " + appError.Message
}

func (appError *AppError) JsonError() map[string]interface{} {
	return map[string]interface{}{
		"code":    appError.Code,
		"message": appError.Message,
	}
}

func (appError *AppError) AppendMessage(anyErrs ...interface{}) *AppError {
	return &AppError{
		Message:  fmt.Sprintf("%v : %v", appError.Message, anyErrs),
		Code:     appError.Code,
		HTTPCode: appError.HTTPCode,
	}
}

func Is(err1 error, err2 *AppError) bool {
	err, ok := err1.(*AppError)
	if !ok {
		return false
	}

	return err.Code == err2.Code
}
