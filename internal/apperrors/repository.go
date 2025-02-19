package apperrors

import "net/http"

var (
	NewSessionRepositoryTypeNotSupported = AppError{
		Message:  "Session repository type is not supported",
		Code:     "SESSION_REPOSITORY_TYPE_NOT_SUPPORTED",
		HTTPCode: http.StatusBadRequest,
	}

	NewUserRepositoryTypeNotSupported = AppError{
		Message:  "User repository type is not supported",
		Code:     "USER_REPOSITORY_TYPE_NOT_SUPPORTED",
		HTTPCode: http.StatusBadRequest,
	}

	SessionRepositoryRedisGetUserSessionBySessionNotFoundError = AppError{
		Message:  "User session not found",
		Code:     "SESSION_REPOSITORY_REDIS_GET_USER_SESSION_BY_SESSION_NOT_FOUND_ERROR",
		HTTPCode: http.StatusNotFound,
	}

	SessionRepositoryRedisGetUserSessionBySessionGetError = AppError{
		Message:  "Error while getting user session by session",
		Code:     "SESSION_REPOSITORY_REDIS_GET_USER_SESSION_BY_SESSION_GET_ERROR",
		HTTPCode: http.StatusInternalServerError,
	}

	SessionRepositoryRedisGetUserSessionBySessionParseError = AppError{
		Message:  "Error while parsing user session",
		Code:     "SESSION_REPOSITORY_REDIS_GET_USER_SESSION_BY_SESSION_PARSE_ERROR",
		HTTPCode: http.StatusInternalServerError,
	}

	SessionRepositoryMysqlGetUserSessionBySessionNotFoundError = AppError{
		Message:  "User session not found",
		Code:     "SESSION_REPOSITORY_MYSQL_GET_USER_SESSION_BY_SESSION_NOT_FOUND_ERROR",
		HTTPCode: http.StatusNotFound,
	}

	SessionRepositoryMysqlGetUserSessionBySessionGetError = AppError{
		Message:  "Error while getting user session by session",
		Code:     "SESSION_REPOSITORY_MYSQL_GET_USER_SESSION_BY_SESSION_GET_ERROR",
		HTTPCode: http.StatusInternalServerError,
	}

	SessionRepositoryMysqlGetUserSessionBySessionScanEmpty = AppError{
		Message:  "User session is empty",
		Code:     "SESSION_REPOSITORY_MYSQL_GET_USER_SESSION_BY_SESSION_SCAN_EMPTY",
		HTTPCode: http.StatusNotFound,
	}

	SessionRepositoryMysqlSetUserSessionBySessionSetError = AppError{
		Message:  "Error while setting user session by session",
		Code:     "SESSION_REPOSITORY_MYSQL_SET_USER_SESSION_BY_SESSION_SET_ERROR",
		HTTPCode: http.StatusInternalServerError,
	}

	SessionRepositoryRedisSetUserSessionBySessionMarshalError = AppError{
		Message:  "Error while marshaling user session",
		Code:     "SESSION_REPOSITORY_REDIS_SET_USER_SESSION_BY_SESSION_MARSHAL_ERROR",
		HTTPCode: http.StatusInternalServerError,
	}

	SessionRepositoryRedisSetUserSessionBySessionSetError = AppError{
		Message:  "Error while setting user session by session",
		Code:     "SESSION_REPOSITORY_REDIS_SET_USER_SESSION_BY_SESSION_SET_ERROR",
		HTTPCode: http.StatusInternalServerError,
	}

	SessionRepositoryRedisGetUserIDBySessionNotFoundError = AppError{
		Message:  "User id not found",
		Code:     "SESSION_REPOSITORY_REDIS_GET_USER_ID_BY_SESSION_NOT_FOUND_ERROR",
		HTTPCode: http.StatusNotFound,
	}

	SessionRepositoryRedisGetUserIDBySessionGetError = AppError{
		Message:  "Error while getting user id by session",
		Code:     "SESSION_REPOSITORY_REDIS_GET_USER_ID_BY_SESSION_GET_ERROR",
		HTTPCode: http.StatusInternalServerError,
	}

	SessionRepositoryRedisGetUserIDBySessionParseError = AppError{
		Message:  "Error while parsing user id",
		Code:     "SESSION_REPOSITORY_REDIS_GET_USER_ID_BY_SESSION_PARSE_ERROR",
		HTTPCode: http.StatusInternalServerError,
	}

	SessionRepositoryRedisBindingSessionIDWithUserIDSetError = AppError{
		Message:  "Error while setting session id with user id",
		Code:     "SESSION_REPOSITORY_REDIS_BINDING_SESSION_ID_WITH_USER_ID_SET_ERROR",
		HTTPCode: http.StatusInternalServerError,
	}

	SessionRepositoryRedisBindingSessionIDWithUserIDMarshalError = AppError{
		Message:  "Error while marshaling user id",
		Code:     "SESSION_REPOSITORY_REDIS_BINDING_SESSION_ID_WITH_USER_ID_MARSHAL_ERROR",
		HTTPCode: http.StatusInternalServerError,
	}

	SessionRepositoryRedisCreateUserIDBySessionSetError = AppError{
		Message:  "Error while setting user id by session",
		Code:     "SESSION_REPOSITORY_REDIS_CREATE_USER_ID_BY_SESSION_SET_ERROR",
		HTTPCode: http.StatusInternalServerError,
	}

	SessionRepositoryRedisDeleteSessionBySessionIDError = AppError{
		Message:  "Error while deleting session by session id",
		Code:     "SESSION_REPOSITORY_REDIS_DELETE_SESSION_BY_SESSION_ID_ERROR",
		HTTPCode: http.StatusInternalServerError,
	}

	SessionRepositoryMysqlGetUserIDBySessionGetError = AppError{
		Message:  "Error while getting user id by session",
		Code:     "SESSION_REPOSITORY_MYSQL_GET_USER_ID_BY_SESSION_GET_ERROR",
		HTTPCode: http.StatusInternalServerError,
	}

	SessionRepositoryMysqlGetUserIDBySessionGetEmpty = AppError{
		Message:  "User id is nil",
		Code:     "SESSION_REPOSITORY_MYSQL_GET_USER_ID_BY_SESSION_GET_EMPTY",
		HTTPCode: http.StatusNotFound,
	}

	SessionRepositoryMysqlGetUserIDBySessionScanError = AppError{
		Message:  "Error while scanning user id",
		Code:     "SESSION_REPOSITORY_MYSQL_GET_USER_ID_BY_SESSION_SCAN_ERROR",
		HTTPCode: http.StatusInternalServerError,
	}

	SessionRepositoryMysqlGetUserIDBySessionScanEmpty = AppError{
		Message:  "User id is empty",
		Code:     "SESSION_REPOSITORY_MYSQL_GET_USER_ID_BY_SESSION_SCAN_EMPTY",
		HTTPCode: http.StatusNotFound,
	}

	SessionRepositoryMysqlGetUserIDBySessionParseError = AppError{
		Message:  "Error while parsing user id",
		Code:     "SESSION_REPOSITORY_MYSQL_GET_USER_ID_BY_SESSION_PARSE_ERROR",
		HTTPCode: http.StatusInternalServerError,
	}

	SessionRepositoryMysqlBindingSessionIDWithUserIDSetError = AppError{
		Message:  "Error while setting session id with user id",
		Code:     "SESSION_REPOSITORY_MYSQL_BINDING_SESSION_ID_WITH_USER_ID_SET_ERROR",
		HTTPCode: http.StatusInternalServerError,
	}

	SessionRepositoryMysqlCreateUserIDBySessionSetError = AppError{
		Message:  "Error while setting user id by session",
		Code:     "SESSION_REPOSITORY_MYSQL_CREATE_USER_ID_BY_SESSION_SET_ERROR",
		HTTPCode: http.StatusInternalServerError,
	}

	SessionRepositoryMysqlDeleteSessionBySessionIDError = AppError{
		Message:  "Error while deleting session by session id",
		Code:     "SESSION_REPOSITORY_MYSQL_DELETE_SESSION_BY_SESSION_ID_ERROR",
		HTTPCode: http.StatusInternalServerError,
	}

	UserRepositoryMysqlCreateUserError = AppError{
		Message:  "Error while creating user",
		Code:     "USER_REPOSITORY_MYSQL_CREATE_USER_ERROR",
		HTTPCode: http.StatusInternalServerError,
	}

	UserRepositoryMysqlGetUserNotFoundError = AppError{
		Message:  "User not found",
		Code:     "USER_REPOSITORY_MYSQL_GET_USER_NOT_FOUND_ERROR",
		HTTPCode: http.StatusNotFound,
	}

	UserRepositoryMysqlGetUserError = AppError{
		Message:  "Error while getting user",
		Code:     "USER_REPOSITORY_MYSQL_GET_USER_ERROR",
		HTTPCode: http.StatusInternalServerError,
	}

	UserRepositoryMysqlGetUserScanEmpty = AppError{
		Message:  "User is empty",
		Code:     "USER_REPOSITORY_MYSQL_GET_USER_SCAN_EMPTY",
		HTTPCode: http.StatusNotFound,
	}

	UserRepositoryMysqlDeleteUserError = AppError{
		Message:  "Error while deleting user",
		Code:     "USER_REPOSITORY_MYSQL_DELETE_USER_ERROR",
		HTTPCode: http.StatusInternalServerError,
	}
)
