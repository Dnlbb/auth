package AccessPolicies

import (
	"errors"
)

// Check проверка доступа к эндпоинту.
func (a AccessPolicyRepository) Check(path string, role string) error {
	paths, ok := a.policies[role]
	if !ok {
		return errors.New("access role not found")
	}
	for _, p := range paths {
		if p == path {
			return nil
		}
	}

	return errors.New("access policy not found")
}
