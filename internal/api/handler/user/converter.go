package user

import (
	"time"

	"github.com/Dnlbb/auth/internal/models"
	userv1 "github.com/Dnlbb/auth/pkg/user_v1"
	"google.golang.org/protobuf/types/known/timestamppb"
)

const (
	strRoleUnspecified = "ROLE_UNSPECIFIED"
	strUser            = "USER"
	strAdmin           = "ADMIN"
)

func toModelRole(role userv1.Role) string {
	switch role {
	case userv1.Role_ROLE_UNSPECIFIED:
		return strRoleUnspecified
	case userv1.Role_ADMIN:
		return strAdmin
	case userv1.Role_USER:
		return strUser
	}

	return "ROLE_UNSPECIFIED"
}

func toModelUser(user *userv1.CreateRequest) models.User {
	role := toModelRole(user.GetUser().GetRole())

	return models.User{
		Name:     user.GetUser().GetName(),
		Email:    user.GetUser().GetEmail(),
		Role:     role,
		Password: user.Password}
}

func toProtoUserProfile(user models.User) *userv1.GetByResponse {
	userRole := role2String(user.Role)

	return &userv1.GetByResponse{
		Id: user.ID,
		User: &userv1.User{
			Name:  user.Name,
			Email: user.Email,
			Role:  userRole,
		},
		CreatedAt: toTimestampProto(user.CreatedAt),
		UpdatedAt: toTimestampProto(user.UpdatedAt),
	}
}

func toUpdateUser(update *userv1.UpdateRequest) *models.User {
	return &models.User{ID: update.GetId(),
		Name:  update.Name.Value,
		Email: update.Email.Value,
		Role:  toModelRole(update.GetRole())}
}

// Role2String Определяем функцию конвертации из строки в Role
func role2String(roleStr string) userv1.Role {
	switch roleStr {
	case strRoleUnspecified:
		return userv1.Role_ROLE_UNSPECIFIED
	case strUser:
		return userv1.Role_USER
	case strAdmin:
		return userv1.Role_ADMIN
	}
	return userv1.Role_ROLE_UNSPECIFIED
}

// ToTimestampProto конвертация времени
func toTimestampProto(time time.Time) *timestamppb.Timestamp {
	return timestamppb.New(time)
}
