package main

import "fmt"

func storeSecret(keychain Keychain, site string, key string, uid string) {
	if len(keychain.Sites[site].UIDs) < 1 {
		fmt.Println("UIDs empty, creating")
		var newSite Site
		newSite.UIDs = make(map[string]string)
		newSite.UIDs[uid] = key
		keychain.Sites[site] = newSite
	} else {
		fmt.Println("UIDs non empty, setting")
		keychain.Sites[site].UIDs[uid] = key
	}
}

func getSecret(keychain Keychain, site string, uid string) string {
	_, _, _ = keychain, site, uid
	return "secret key"
}