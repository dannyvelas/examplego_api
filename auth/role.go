package auth

import (
	"errors"
	"strings"
)

type Role uint8

const (
	UndefinedRole Role = iota
	AdminRole
	UserRole
)

func NewRole(value string) (Role, error) {
	switch strings.ToLower(value) {
	case "admin":
		return AdminRole, nil
	case "user":
		return UserRole, nil
	default:
		return UndefinedRole, errors.New("Invalid auth role of " + value)
	}
}

func String(role Role) string {
	switch role {
	case AdminRole:
		return "Admin"
	case UserRole:
		return "User"
	default:
		return ""
	}
}
