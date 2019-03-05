package storage

import (
	"github.com/daveio/gotp/errors"
	"github.com/daveio/gotp/verbose"
	"github.com/tucnak/store"
	"sort"
)

var (
	V = verbose.V
)

func InitStorage(appName string) {
	store.Init(appName)
}

func LoadConfig() (keychain Keychain, error error) {
	var k Keychain
	err := store.Load("keychain.json", &k)
	if err == nil {
		return k, nil
	} else {
		panic(err)
	}
}

func SaveConfig(keychain *Keychain) (error error) {
	err := store.Save("keychain.json", &keychain)
	if err == nil {
		return nil
	} else {
		return err
	}
}

func StoreSecret(keychain Keychain, site string, key string, uid string) (error error) {
	if len(keychain.Sites[site].UIDs) < 1 {
		var newSite Site
		newSite.UIDs = make(map[string]string)
		newSite.UIDs[uid] = key
		keychain.Sites[site] = newSite
	} else {
		keychain.Sites[site].UIDs[uid] = key
	}
	return nil
}

func GetSecret(keychain Keychain, site string, uid string) (string, error) {
	kSite, present := keychain.Sites[site]
	if present {
		kSecret, present := kSite.UIDs[uid]
		if present {
			return kSecret, nil
		} else {
			return "", &errors.UidError{Site: site, Uid: uid}
		}
	} else {
		return "", &errors.SiteError{Site: site}
	}
}

func GetSiteNames(keychain Keychain, sortNames bool) (siteNames []string) {
	sites := keychain.Sites
	var allNames []string
    for thisName := range sites {
        allNames = append(allNames, thisName)
    }
    if sortNames {
    	sort.Strings(allNames)
	}
    return allNames
}

func GetUidsForSite(keychain Keychain, siteName string, sortNames bool) (uids []string) {
	site, pres := keychain.Sites[siteName]
	if pres {
		uids := site.UIDs
		var allUids []string
		for thisUid := range uids {
			allUids = append(allUids, thisUid)
		}
		if sortNames {
			sort.Strings(allUids)
		}
		return allUids
	} else {
		return []string{}
	}
}

func DeleteUID(keychain Keychain, siteName string, uid string) (error error) {
	site, sPres := keychain.Sites[siteName]
	if sPres {
		_, uPres := site.UIDs[uid]
		if uPres {
			delete(site.UIDs, uid)
		}
	}
	return nil
}

func DeleteSite(keychain Keychain, siteName string) (error error) {
	_, sPres := keychain.Sites[siteName]
	if sPres {
		delete(keychain.Sites, siteName)
	}
	return nil
}