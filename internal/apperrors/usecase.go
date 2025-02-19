package apperrors

import "net/http"

var (
	UsecaseGetUserByIDError = AppError{
		Message:  "The user usecase get user by id operation has been failed",
		Code:     "USECASE_GET_USER_BY_ID",
		HTTPCode: http.StatusBadRequest,
	}

	UsecaseGetUserByNicknameError = AppError{
		Message:  "The user usecase get user by nickname operation has been failed",
		Code:     "USECASE_GET_USER_BY_NICKNAME",
		HTTPCode: http.StatusBadRequest,
	}

	UsecaseCreateUserError = AppError{
		Message:  "The user usecase create user operation has been failed",
		Code:     "USECASE_CREATE_USER",
		HTTPCode: http.StatusBadRequest,
	}

	UsecaseUpdateUserError = AppError{
		Message:  "The user usecase update user operation has been failed",
		Code:     "USECASE_UPDATE_USER",
		HTTPCode: http.StatusBadRequest,
	}

	UsecaseCheckUserByNicknameError = AppError{
		Message:  "The user usecase check user by nickname operation has been failed",
		Code:     "USECASE_CHECK_USER_BY_NICKNAME",
		HTTPCode: http.StatusBadRequest,
	}

	UsecaseDeleteUserError = AppError{
		Message:  "The user usecase delete user operation has been failed",
		Code:     "USECASE_DELETE_USER",
		HTTPCode: http.StatusBadRequest,
	}

	UsecaseGetUsersError = AppError{
		Message:  "The user usecase get users operation has been failed",
		Code:     "USECASE_GET_USERS",
		HTTPCode: http.StatusBadRequest,
	}

	AuthUsecaseSignupGetUserIDBySessionID = AppError{
		Message:  "The auth usecase signup get user id by session id operation has been failed",
		Code:     "AUTH_USECASE_SIGNUP_GET_USER_ID_BY_SESSION_ID",
		HTTPCode: http.StatusBadRequest,
	}

	AuthUsecaseSignupGetUser = AppError{
		Message:  "The auth usecase signup get user operation has been failed",
		Code:     "AUTH_USECASE_SIGNUP_GET_USER",
		HTTPCode: http.StatusBadRequest,
	}

	AuthUsecaseSignupSessionRepoCacheGetUserSessionBySessionID = AppError{
		Message:  "The auth usecase signup session repo cache get user session by session id operation has been failed",
		Code:     "AUTH_USECASE_SIGNUP_SESSION_REPO_CACHE_GET_USER_SESSION_BY_SESSION_ID",
		HTTPCode: http.StatusBadRequest,
	}

	AuthUsecaseSignupSessionRepoGetUserSessionBySessionID = AppError{
		Message:  "The auth usecase signup session repo get user session by session id operation has been failed",
		Code:     "AUTH_USECASE_SIGNUP_SESSION_REPO_GET_USER_SESSION_BY_SESSION_ID",
		HTTPCode: http.StatusBadRequest,
	}

	AuthUsecaseSessionRepoCacheGetUserSessionBySessionID = AppError{
		Message:  "The auth usecase session repo cache get user session by session id operation has been failed",
		Code:     "AUTH_USECASE_SESSION_REPO_CACHE_GET_USER_SESSION_BY_SESSION_ID",
		HTTPCode: http.StatusBadRequest,
	}

	AuthUsecaseSessionRepoGetUserSessionBySessionID = AppError{
		Message:  "The auth usecase session repo get user session by session id operation has been failed",
		Code:     "AUTH_USECASE_SESSION_REPO_GET_USER_SESSION_BY_SESSION_ID",
		HTTPCode: http.StatusBadRequest,
	}

	AuthUsecaseLoginGeUserSessionBySessionId = AppError{
		Message:  "The auth usecase login ge user session by session id operation has been failed",
		Code:     "AUTH_USECASE_LOGIN_GE_USER_SESSION_BY_SESSION_ID",
		HTTPCode: http.StatusBadRequest,
	}

	AuthUsecaseLoginUserSessionNotFound = AppError{
		Message:  "The auth usecase login user session not found operation has been failed",
		Code:     "AUTH_USECASE_LOGIN_USER_SESSION_NOT_FOUND",
		HTTPCode: http.StatusBadRequest,
	}

	AuthUsecaseSignupGeUserSessionBySessionId = AppError{
		Message:  "The auth usecase signup ge user session by session id operation has been failed",
		Code:     "AUTH_USECASE_SIGNUP_GE_USER_SESSION_BY_SESSION_ID",
		HTTPCode: http.StatusBadRequest,
	}

	AuthUsecaseGeUserSessionBySessionId = AppError{
		Message:  "The auth usecase ge user session by session id operation has been failed",
		Code:     "AUTH_USECASE_GE_USER_SESSION_BY_SESSION_ID",
		HTTPCode: http.StatusBadRequest,
	}

	AuthUsecaseSignupSetUserSession = AppError{
		Message:  "The auth usecase signup set user session operation has been failed",
		Code:     "AUTH_USECASE_SIGNUP_SET_USER_SESSION",
		HTTPCode: http.StatusBadRequest,
	}

	AuthUsecaseSignupCreateUser = AppError{
		Message:  "The auth usecase signup create user operation has been failed",
		Code:     "AUTH_USECASE_SIGNUP_CREATE_USER",
		HTTPCode: http.StatusBadRequest,
	}

	AuthUsecaseSignupGetUserByID = AppError{
		Message:  "The auth usecase signup get user by id operation has been failed",
		Code:     "AUTH_USECASE_SIGNUP_GET_USER_BY_ID",
		HTTPCode: http.StatusBadRequest,
	}

	AuthUsecaseSessionRepoSetUserSession = AppError{
		Message:  "The auth usecase session repo set user session operation has been failed",
		Code:     "AUTH_USECASE_SESSION_REPO_SET_USER_SESSION",
		HTTPCode: http.StatusBadRequest,
	}

	AuthUsecaseSessionRepoCacheSetUserSession = AppError{
		Message:  "The auth usecase session repo cache set user session operation has been failed",
		Code:     "AUTH_USECASE_SESSION_REPO_CACHE_SET_USER_SESSION",
		HTTPCode: http.StatusBadRequest,
	}

	AuthUsecaseGetUserBySessionID = AppError{
		Message:  "The auth usecase get user by session id operation has been failed",
		Code:     "AUTH_USECASE_GET_USER_BY_SESSION_ID",
		HTTPCode: http.StatusBadRequest,
	}

	AuthUsecaseBindUserIDToSessionID = AppError{
		Message:  "The auth usecase bind user id to session id operation has been failed",
		Code:     "AUTH_USECASE_BIND_USER_ID_TO_SESSION_ID",
		HTTPCode: http.StatusBadRequest,
	}

	AuthUsecaseSignupBindUserIDToSessionID = AppError{
		Message:  "The auth usecase signup bind user id to session id operation has been failed",
		Code:     "AUTH_USECASE_SIGNUP_BIND_USER_ID_TO_SESSION_ID",
		HTTPCode: http.StatusBadRequest,
	}
)
