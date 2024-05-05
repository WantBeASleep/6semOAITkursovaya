package lib

import (
	"fmt"
	"math/rand"
	"kra/constants"
)

func GetRandString(n int) string {
	str := make([]byte, 0, n)
	for range n {
		str = append(str, byte(97 + rand.Intn(122 - 97)))
	}
	return string(str)
}

var domens = []string{"mail.ru", "yandex.ru", "gmail.com", "hotmail.com"}
func GenUser() User {
	return User{
		Login: "user@" + GetRandString(constants.LoginMaxGenLen),
		Email: "mail" + GetRandString(constants.EmailMaxGenLen) + "@" + domens[rand.Intn(len(domens))],
		Password: GetRandString(constants.PasswordMaxGenLen),
		Role: 0,
	}
}

var network = []string{"twitter", "vk", "insta", "telegram", "facebook", "linkedin"}
func GenResource() ExternalResource {
	id := rand.Intn(len(network))
	return ExternalResource{
		Link: "https://" + network[id] + ".com/" + fmt.Sprint(rand.Intn(constants.LinkMaxId)),
		RType: network[id],
	}
}