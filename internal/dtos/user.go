package dtos

type User struct {
	Id      uint64
	Name    string
	Phone   string
	isAdmin bool
}

func (user *User) IsAdmin() (bool, error) {
	return user.isAdmin && len(user.Phone) > 0, nil
}
