package main

func cmdStore(keychain Keychain, site string, key string) {
	cmdStoreWithUID(keychain, site, key, "__default")
}

func cmdStoreWithUID(keychain Keychain, site string, key string, uid string) {
	storeSecret(keychain, site, key, uid)
}