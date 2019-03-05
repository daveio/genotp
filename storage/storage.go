package storage

import (
	"github.com/daveio/gotp/errors"
	"github.com/tucnak/store"
)

func InitStorage(appName string) {
	store.Init(appName)
}

func LoadConfig() Keychain {
	var keychain Keychain
	err := store.Load("keychain.json", &keychain)
	if err != nil {
		panic(err)
	}
	return keychain
}

func SaveConfig(keychain *Keychain) {
	err := store.Save("keychain.json", &keychain)
	if err != nil {
		panic(err)
	}
}

func StoreSecret(keychain Keychain, site string, key string, uid string) {
	if len(keychain.Sites[site].UIDs) < 1 {
		var newSite Site
		newSite.UIDs = make(map[string]string)
		newSite.UIDs[uid] = key
		keychain.Sites[site] = newSite
	} else {
		keychain.Sites[site].UIDs[uid] = key
	}
}

func GetSecret(keychain Keychain, site string, uid string) (string, error) {
	kSite, present := keychain.Sites[site]
	if present {
		kSecret, present := kSite.UIDs[uid]
		if present {
			return kSecret, nil
		} else {
			return "", &errors.UidError{site, uid}
		}
	} else {
		return "", &errors.SiteError{site}
	}
}