package pojo

type User struct {
	name string
	age  int
	sex  string
}

type FuncAttr func(*User)
type FuncAttrs []FuncAttr

func NewUser(fs ...FuncAttr) *User {
	user := &User{}
	FuncAttrs(fs).Apply(user)
	return user
}

func UserWithName(name string) FuncAttr {
	return func(user *User) {
		user.name = name
	}
}

func UserWithAge(age int) FuncAttr {
	return func(user *User) {
		user.age = age
	}
}

func UserWithSex(sex string) FuncAttr {
	return func(user *User) {
		user.sex = sex
	}
}

func (fs FuncAttrs) Apply(user *User) {
	for _, f := range fs {
		f(user)
	}
}
