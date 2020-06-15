package user

type User struct {
	Name string
}

// UserRepository 用户仓库
type UserRepository interface {
	FindOne(id int) (*User, error)
}
