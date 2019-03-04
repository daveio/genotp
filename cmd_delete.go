package main

func cmdDelete(keychain Keychain, site string) {
	cmdDeleteWithUID(keychain, site, "__default")
}

func cmdDeleteWithUID(keychain Keychain, site string, uid string) {
	_ = keychain
}