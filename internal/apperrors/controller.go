package apperrors

import "net/http"

var (
	ApiControllerSignupBind = AppError{
		Message:  "Error while binding signup request",
		Code:     "API_CONTROLLER_SIGNUP_BIND",
		HTTPCode: http.StatusBadRequest,
	}

	ApiControllerLoginBind = AppError{
		Message:  "Error while binding login request",
		Code:     "API_CONTROLLER_LOGIN_BIND",
		HTTPCode: http.StatusBadRequest,
	}

	ApiControllerLoginUserNotExist = AppError{
		Message:  "User does not exist",
		Code:     "API_CONTROLLER_LOGIN_USER_NOT_EXIST",
		HTTPCode: http.StatusNotFound,
	}

	UserControllerCreateUserBind = AppError{
		Message:  "Error while binding user creation request",
		Code:     "USER_CONTROLLER_CREATE_USER_BIND",
		HTTPCode: http.StatusBadRequest,
	}
	UserControllerCreateUserError = AppError{
		Message:  "Error while creating user",
		Code:     "USER_CONTROLLER_CREATE_USER_ERROR",
		HTTPCode: http.StatusInternalServerError,
	}

	UserControllerFetchJWTUser = AppError{
		Message:  "Error while fetching JWT user",
		Code:     "USER_CONTROLLER_FETCH_JWT_USER",
		HTTPCode: http.StatusBadRequest,
	}

	UserControllerFetchJWTUserUserContextNotValid = AppError{
		Message:  "User context is not valid",
		Code:     "USER_CONTROLLER_FETCH_JWT_USER_USER_CONTEXT_NOT_VALID",
		HTTPCode: http.StatusBadRequest,
	}

	UserControllerGetUserUuidParse = AppError{
		Message:  "Error while parsing user UUID",
		Code:     "USER_CONTROLLER_GET_USER_UUID_PARSE",
		HTTPCode: http.StatusBadRequest,
	}

	UserControllerGetUserUserNotExist = AppError{
		Message:  "User does not exist",
		Code:     "USER_CONTROLLER_GET_USER_USER_NOT_EXIST",
		HTTPCode: http.StatusNotFound,
	}

	UserControllerUpdateUserUuidParse = AppError{
		Message:  "Error while parsing user UUID",
		Code:     "USER_CONTROLLER_UPDATE_USER_UUID_PARSE",
		HTTPCode: http.StatusBadRequest,
	}

	UserControllerUpdateUserUserNotExist = AppError{
		Message:  "User does not exist",
		Code:     "USER_CONTROLLER_UPDATE_USER_USER_NOT_EXIST",
		HTTPCode: http.StatusNotFound,
	}

	UserControllerUpdateUserBind = AppError{
		Message:  "Error while binding user update request",
		Code:     "USER_CONTROLLER_UPDATE_USER_BIND",
		HTTPCode: http.StatusBadRequest,
	}

	UserControllerTryToSetAdmin = AppError{
		Message:  "Error while trying to set admin",
		Code:     "USER_CONTROLLER_TRY_TO_SET_ADMIN",
		HTTPCode: http.StatusForbidden,
	}

	UserControllerDeleteUserUuidParse = AppError{
		Message:  "Error while parsing user UUID",
		Code:     "USER_CONTROLLER_DELETE_USER_UUID_PARSE",
		HTTPCode: http.StatusBadRequest,
	}

	UserControllerGetUsersGetPaginationFromCtx = AppError{
		Message:  "Error while getting pagination from context",
		Code:     "USER_CONTROLLER_GET_USERS_GET_PAGINATION_FROM_CTX",
		HTTPCode: http.StatusBadRequest,
	}

	ChatControllerChatSupportUpgrade = AppError{
		Message:  "Error while upgrading connection",
		Code:     "CHAT_CONTROLLER_CHAT_SUPPORT_UPGRADE",
		HTTPCode: http.StatusInternalServerError,
	}

	ChatControllerChatSupportReadMessage = AppError{
		Message:  "Error while reading message",
		Code:     "CHAT_CONTROLLER_CHAT_SUPPORT_READ_MESSAGE",
		HTTPCode: http.StatusInternalServerError,
	}

	ChatControllerChatSupportWriteMessage = AppError{
		Message:  "Error while writing message",
		Code:     "CHAT_CONTROLLER_CHAT_SUPPORT_WRITE_MESSAGE",
		HTTPCode: http.StatusInternalServerError,
	}
)
