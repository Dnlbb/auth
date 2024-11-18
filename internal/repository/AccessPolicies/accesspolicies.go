package AccessPolicies

import (
	"github.com/Dnlbb/auth/internal/repository/repointerface"
)

// AccessPolicyRepository хранилище с политиками доступа для юзеров.
type AccessPolicyRepository struct {
	policies map[string][]string
}

// NewAccessPolicyRepository конструктор для хранилища с политиками доступа.
func NewAccessPolicyRepository() repointerface.AccessPolicies {
	return AccessPolicyRepository{policies: map[string][]string{
		"USER":             {"api.chat/SendMessage"},
		"ADMIN":            {"api.chat/SendMessage", "api.chat/Delete", "api.chat/Create"},
		"ROLE_UNSPECIFIED": {"api.chat/SendMessage"}}}
}
