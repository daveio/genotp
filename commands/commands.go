package commands

import (
	"fmt"
	"github.com/daveio/gotp/otputils"
	"github.com/daveio/gotp/storage"
	"github.com/daveio/gotp/verbose"
)

var (
	V = verbose.V
)

func Generate(keychain storage.Keychain, site string, uid string) (error error) {
	secretKey, errGet := storage.GetSecret(keychain, site, uid)
	if errGet == nil {
		otp, errGen := otputils.GenerateOTP(secretKey)
		if errGen == nil {
			if uid == "__default" {
				fmt.Printf("%s : ", site)
			} else {
				fmt.Printf("%s @ %s : ", uid, site)
			}
			fmt.Printf("%s\n", otp)
			return nil
		} else {
			return errGen
		}
	} else {
		return errGet
	}
}

func Store(keychain storage.Keychain, site string, key string, uid string) (error error) {
	err := storage.StoreSecret(keychain, site, key, uid)
	if err == nil {
		return nil
	} else {
		return err
	}
}

func ListSites(keychain storage.Keychain) (error error) {
	siteNames := storage.GetSiteNames(keychain, true)
	for siteIdx := range siteNames {
		fmt.Println(siteNames[siteIdx])
	}
	return nil
}

func ListUids(keychain storage.Keychain, siteName string) (error error) {
	siteUids := storage.GetUidsForSite(keychain, siteName, true)
	for uidIdx := range siteUids {
		fmt.Println(siteUids[uidIdx])
	}
	return nil
}

func DeleteUID(keychain storage.Keychain, siteName string, uid string) (error error) {
	err := storage.DeleteUID(keychain, siteName, uid)
	return err
}

func DeleteSite(keychain storage.Keychain, siteName string) (error error) {
	err := storage.DeleteSite(keychain, siteName)
	return err
}

func Hello() {
	fmt.Println("https://github.com/daveio/gotp")
}

// TODO URI parsing command