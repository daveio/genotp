package commands

import (
	"fmt"
	"github.com/daveio/gotp/otputils"
	"github.com/daveio/gotp/storage"
)

func Generate(keychain storage.Keychain, site string, uid string) {
	secretKey, err := storage.GetSecret(keychain, site, uid)
	if err == nil {
		if uid == "__default" {
			fmt.Printf("%s : ", site)
		} else {
			fmt.Printf("%s @ %s : ", uid, site)
		}
		fmt.Printf("%s\n", otputils.GenerateOTP(secretKey))
	} else {
		fmt.Printf("%s", err)
	}
}

func Store(keychain storage.Keychain, site string, key string, uid string) {
	storage.StoreSecret(keychain, site, key, uid)
}

func Delete(keychain storage.Keychain, site string, uid string) {
	_ = keychain
}