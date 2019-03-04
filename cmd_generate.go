package main

import (
	"fmt"
)

func cmdGenerate(keychain Keychain, site string) {
	cmdGenerateWithUID(keychain, site, "__default")
}

func cmdGenerateWithUID(keychain Keychain, site string, uid string) {
	var secretKey string
	if uid == "__default" {
		fmt.Printf("%s : ", site)
	} else {
		fmt.Printf("%s @ %s : ", uid, site)
	}
	secretKey = getSecret(keychain, site, uid)
	fmt.Printf("%s\n", generateOTP(secretKey))
}