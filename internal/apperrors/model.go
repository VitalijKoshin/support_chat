package apperrors

import "net/http"

var (
	UserHashPasswordGenerateFromPassword = AppError{
		Message:  "The user hash password generate operation has been failed",
		Code:     "USER_REPO_FIND_USER_BY_UUID_GET_CONTEXT",
		HTTPCode: http.StatusBadRequest,
	}

	UserComparePasswordsCompareHashAndPassword = AppError{
		Message:  "The user compare hash and password operation has been failed",
		Code:     "USER_COMPARE_PASSWORDS_COMPARE_HASH_AND_PASSWORD",
		HTTPCode: http.StatusBadRequest,
	}

	RoleCanNoPermission = AppError{
		Message:  "Auth user doesn't have permissions",
		Code:     "ROLE_CAN_NO_PERMISSION",
		HTTPCode: http.StatusBadRequest,
	}

	HasPermissionsToUpdateUser = AppError{
		Message:  "Auth user doesn't have permission to update user",
		Code:     "HAS_PERMISSIONS_TO_UPDATE_USER",
		HTTPCode: http.StatusBadRequest,
	}

	HasPermissionsDeleteUser = AppError{
		Message:  "Auth user doesn't have permission to delete user",
		Code:     "HAS_PERMISSIONS_DELETE_USER",
		HTTPCode: http.StatusBadRequest,
	}
)
