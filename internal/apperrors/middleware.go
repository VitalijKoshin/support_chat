package apperrors

import "net/http"

var (
	HasPermissionUuidParse = AppError{
		Message:  "Failed to parse user id",
		Code:     "HAS_PERMISSION_UUID_PARSE",
		HTTPCode: http.StatusBadRequest,
	}

	MiddlewareJWTAuthValid = AppError{
		Message:  "JWT token is not valid",
		Code:     "MIDDLEWARE_JWT_AUTH_VALID",
		HTTPCode: http.StatusUnauthorized,
	}

	MiddlewareJWTAuthVerifyJwtUser = AppError{
		Message:  "Failed to verify jwt user",
		Code:     "MIDDLEWARE_JWT_AUTH_VERIFY_JWT_USER",
		HTTPCode: http.StatusUnauthorized,
	}

	MiddlewareVerifyJwtUserGetUserByNickname = AppError{
		Message:  "Failed to get user by nickname",
		Code:     "MIDDLEWARE_VERIFY_JWT_USER_GET_USER_BY_NICKNAME",
		HTTPCode: http.StatusUnauthorized,
	}

	MiddlewareVerifyAuthUserGetUserByNickname = AppError{
		Message:  "Failed to get user by nickname",
		Code:     "MIDDLEWARE_VERIFY_AUTH_USER_GET_USER_BY_NICKNAME",
		HTTPCode: http.StatusUnauthorized,
	}

	MiddlewareVerifyAuthUserComparePasswords = AppError{
		Message:  "Failed to compare passwords",
		Code:     "MIDDLEWARE_VERIFY_AUTH_USER_COMPARE_PASSWORDS",
		HTTPCode: http.StatusUnauthorized,
	}

	HasPermissionFetchJwtUser = AppError{
		Message:  "Failed to fetch jwt user",
		Code:     "HAS_PERMISSION_FETCH_JWT_USER",
		HTTPCode: http.StatusUnauthorized,
	}

	MiddlewareJWTMiddlewareMissingToken = AppError{
		Message:  "Missing token",
		Code:     "MIDDLEWARE_JWT_MIDDLEWARE_MISSING_TOKEN",
		HTTPCode: http.StatusUnauthorized,
	}
)
