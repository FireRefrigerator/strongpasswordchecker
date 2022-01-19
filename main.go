package main

import (
	"flag"
	"fmt"

	app "github.com/strongpasswordchecker/app"
)

func main() {
	var pwdinput string
	flag.StringVar(&pwdinput, "pwd", "pwd", "-pwd xxx")
	if inBlackList(pwdinput) {
		fmt.Printf("pwd %s is invalid \n", pwdinput)
	}
	if app.CheckPassword(pwdinput) {
		fmt.Printf("pwd %s is invalid \n", pwdinput)
	} else {
		fmt.Printf("pwd %s is check ok \n", pwdinput)
	}
}

func inBlackList(pwd string) bool {
	blackList := pwdBlackList()
	if _, ok := blackList[pwd]; ok {
		return true
	}
	return false
}

func pwdBlackList() map[string]string {
	return map[string]string{
		"123456": "",
		"abcdef": "",
	}
}
