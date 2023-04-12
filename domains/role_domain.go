package domains

import entities "mygram/domains/entity"

// repository contract
type RoleRepository interface {
	Insert(role *entities.Role) error
}