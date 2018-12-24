package example1

type User struct {
	Province string
	City     string
}

type Address struct {
	Username string
	Sex      int
	*User
}

type User01 struct {
	City     string
	Username string
	Sex      int
	*Address
}

//
//func main() {
//	var user User01
//	user.City = "张三"
//	user.Address = new(Address)
//	user.Address.User = new(User)
//	user.Address.User.City = "王五"
//	user.Username = "赵四"
//	fmt.Printf("%#v\n", user.Address.User)
//
//}
