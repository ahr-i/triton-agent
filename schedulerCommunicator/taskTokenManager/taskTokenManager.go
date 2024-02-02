package taskTokenManager

import "log"

var tokenMap map[string]bool

func Init() {
	log.Println("* (System) Set up the tokenMap.")

	tokenMap = make(map[string]bool)
}

func Register(token string) {
	tokenMap[token] = true
}

func IsValidToken(token string) bool {
	_, err_ := tokenMap[token]

	return !err_
}

func UseToken(token string) {
	delete(tokenMap, token)
}
