package main

import (
	"flag"
	"fmt"

	app "github.com/strongpasswordchecker/app"
)

func main() {
	var pwdinput string
	flag.StringVar(&pwdinput, "input", "", "-pwd xxx")
	flag.Parse()
	if inBlackList(pwdinput) {
		fmt.Printf("pwd %s is invalid \n", pwdinput)
		return
	}
	if !app.CheckPassword(pwdinput) {
		fmt.Printf("pwd %s is invalid \n", pwdinput)
		return
	}
	fmt.Printf("pwd %s is check ok \n", pwdinput)
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
