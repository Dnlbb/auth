package converter

import (
	"errors"
	"fmt"
	"log"
	"time"

	"github.com/Dnlbb/auth/internal/models"
	authv1 "github.com/Dnlbb/auth/pkg/auth_v1"
	"google.golang.org/protobuf/types/known/timestamppb"
)

// Role роль юзера
type Role int32

// RoleUnspecified роль юзера
const (
	RoleUnspecified    int32 = 0
	USER               int32 = 1
	ADMIN              int32 = 2
	StrRoleUnspecified       = "ROLE_UNSPECIFIED"
	StrUser                  = "USER"
	StrAdmin                 = "ADMIN"
)

// Role2String Определяем функцию конвертации из строки в Role
func Role2String(roleStr string) (authv1.Role, error) {
	switch roleStr {
	case StrRoleUnspecified:
		return authv1.Role(RoleUnspecified), nil
	case StrUser:
		return authv1.Role(USER), nil
	case StrAdmin:
		return authv1.Role(ADMIN), nil
	default:
		return authv1.Role(RoleUnspecified), fmt.Errorf("некорректное значение роли: %s", roleStr)
	}
}

// UserModel2ProtoUserProfile конвертер сервис модель -> grpc ответ
func UserModel2ProtoUserProfile(user models.User) *authv1.GetResponse {
	userRole, err := Role2String(user.Role)
	if err != nil {
		log.Println(err)
		log.Fatal("error from converter")
	}
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

// toTimestampProto конвертация времени
func toTimestampProto(time time.Time) *timestamppb.Timestamp {
	return timestamppb.New(time)
}

// ProtoRole2ModelsRole конвертация роли юзера
func ProtoRole2ModelsRole(role int32) (string, error) {
	switch role {
	case 0:
		return "ROLE_UNSPECIFIED", nil
	case 1:
		return "USER", nil
	case 2:
		return "ADMIN", nil
	}
	return "", errors.New("role unspecified")
}

// ProtoAddUser2AddUser конвертер grpc -> сервисная модель добавления пользователя
func ProtoAddUser2AddUser(addUser *authv1.CreateRequest) models.UserAdd {
	role, err := ProtoRole2ModelsRole(int32(addUser.GetUser().GetRole()))
	if err != nil {
		log.Fatal("error from converter")
	}
	return models.UserAdd{
		Name:     addUser.GetUser().GetName(),
		Email:    addUser.GetUser().GetEmail(),
		Role:     role,
		Password: addUser.Password}
}

// GetUserParamsReq2Params конвертер запроса grpc -> в сервисную модель GetUserParams
func GetUserParamsReq2Params(userParams *authv1.GetRequest) (*models.GetUserParams, error) {
	var params models.GetUserParams
	switch nameOrID := userParams.NameOrId.(type) {
	case *authv1.GetRequest_Id:
		params.ID = &nameOrID.Id
	case *authv1.GetRequest_Username:
		params.Username = &nameOrID.Username
	default:
		return nil, fmt.Errorf("необходимо указать либо ID, либо Username")
	}
	return &params, nil
}

// ProtoUpdateUser2UpdateUser конвертер запроса grpc -> в сервисную модель UpdateUser
func ProtoUpdateUser2UpdateUser(updateProto *authv1.UpdateRequest) *models.UpdateUser {
	return &models.UpdateUser{ID: updateProto.GetId(),
		Name:  updateProto.Name.Value,
		Email: updateProto.Email.Value,
		Role:  string(updateProto.GetRole())}
}
