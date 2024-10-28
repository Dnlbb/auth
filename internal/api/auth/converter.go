package auth

import (
	"time"

	"github.com/Dnlbb/auth/internal/models"
	authv1 "github.com/Dnlbb/auth/pkg/auth_v1"
	"google.golang.org/protobuf/types/known/timestamppb"
)

const (
	StrRoleUnspecified = "ROLE_UNSPECIFIED"
	StrUser            = "USER"
	StrAdmin           = "ADMIN"
)

func mappingUserParams(userParams *authv1.GetRequest) (*models.GetUserParams, error) {
	var params models.GetUserParams
	switch nameOrID := userParams.NameOrId.(type) {
	case *authv1.GetRequest_Id:
		params.ID = &nameOrID.Id
	case *authv1.GetRequest_Username:
		params.Username = &nameOrID.Username
	}

	return &params, nil
}

func toModelRole(role authv1.Role) string {
	switch role {
	case authv1.Role_ROLE_UNSPECIFIED:
		return StrRoleUnspecified
	case authv1.Role_ADMIN:
		return StrAdmin
	case authv1.Role_USER:
		return StrUser
	}

	return "ROLE_UNSPECIFIED"
}

func toModelUser(user *authv1.CreateRequest) models.User {
	role := toModelRole(user.GetUser().GetRole())

	return models.User{
		Name:     user.GetUser().GetName(),
		Email:    user.GetUser().GetEmail(),
		Role:     role,
		Password: user.Password}
}

func toProtoUserProfile(user models.User) *authv1.GetResponse {
	userRole := role2String(user.Role)

	return &authv1.GetResponse{
		Id: user.ID,
		User: &authv1.User{
			Name:  user.Name,
			Email: user.Email,
			Role:  userRole,
		},
		CreatedAt: toTimestampProto(user.CreatedAt),
		UpdatedAt: toTimestampProto(user.UpdatedAt),
	}
}

func toUpdateUser(update *authv1.UpdateRequest) *models.User {
	return &models.User{ID: update.GetId(),
		Name:  update.Name.Value,
		Email: update.Email.Value,
		Role:  toModelRole(update.GetRole())}
}

// Role2String Определяем функцию конвертации из строки в Role
func role2String(roleStr string) authv1.Role {
	switch roleStr {
	case StrRoleUnspecified:
		return authv1.Role_ROLE_UNSPECIFIED
	case StrUser:
		return authv1.Role_USER
	case StrAdmin:
		return authv1.Role_ADMIN
	}
	return authv1.Role_ROLE_UNSPECIFIED
}

// ToTimestampProto конвертация времени
func toTimestampProto(time time.Time) *timestamppb.Timestamp {
	return timestamppb.New(time)
}
