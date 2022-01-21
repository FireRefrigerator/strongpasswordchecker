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
	if app.CheckPassword(pwdinput) {
		fmt.Printf("pwd %s is strong \n", pwdinput)
	} else {
		fmt.Printf("pwd %s is week \n", pwdinput)
	}

	step := app.CheckPasswordStep(pwdinput)
	fmt.Printf("pwd %s modify to stronger need %d step \n", pwdinput, step)
}
