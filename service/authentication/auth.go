package authentication

import "fmt"

type Auth struct {
	Username string
	Password string
}

func (a *Auth) Check() (bool, error) {
	userName := a.Username
	passWord := a.Password
	ret := false
	if userName == "cwj" && passWord == "123" {
		ret = true
		fmt.Println("login success")
	} else {
		fmt.Printf("login fail")
	}
	return ret, nil
}
