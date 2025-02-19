package model

import (
	"errors"

	"supportchat/internal/apperrors"
)

const (
	RoleClient                        = "client"
	RoleUser                          = "user"
	RoleModerator                     = "moderator"
	RoleAdmin                         = "admin"
	PermissionUpdate                  = "update"
	PermissionDelete                  = "delete"
	hasNoPermissionsToUpdateUserError = "auth user can't change this user"
	hasNoPermissionsToDeleteUserError = "auth user can't delete this user"
	hasNoPermissionsError             = "auth user can't delete this user"
)

func (u *User) GetDefaultRole() string {
	return RoleUser
}

func (u *User) GetRoles() []string {
	return []string{RoleUser, RoleModerator, RoleAdmin}
}

func (u *User) IsAdmin() bool {
	return (u.Role == RoleAdmin)
}

func (u *User) Can(permission string) error {
	switch permission {
	case PermissionUpdate:
		return u.HasPermissionsToUpdateUser()
	case PermissionDelete:
		return u.HasPermissionsToDeleteUser()
	default:
		err := errors.New(hasNoPermissionsError)
		return apperrors.RoleCanNoPermission.AppendMessage(err)
	}
}

func (u *User) HasPermissionsToUpdateUser() error {
	if u.IsAdmin() {
		return nil
	}
	err := errors.New(hasNoPermissionsToUpdateUserError)
	return apperrors.HasPermissionsToUpdateUser.AppendMessage(err)
}

func (u *User) HasPermissionsToDeleteUser() error {
	if u.IsAdmin() {
		return nil
	}
	err := errors.New(hasNoPermissionsToDeleteUserError)
	return apperrors.HasPermissionsDeleteUser.AppendMessage(err)
}
