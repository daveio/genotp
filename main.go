package main

import (
	"fmt"
	"github.com/tucnak/store"
	"gopkg.in/alecthomas/kingpin.v2"
	"os"
)

var (
	app = kingpin.
		New("gotp", "Generate OATH-TOTP one-time passwords from the command line.")
	debug = app.
		Flag("debug", "Enable debug mode.").
		Short('d').
		Bool()
	storeCommand = app.
			Command("store", "Store a new account.").
			Alias("s")
	storeSite = storeCommand.
			Arg("site", "An identifier for the site which you will use to generate OTPs.").
			Required().
			String()
	storeKey = storeCommand.
			Arg("key", "Your secret key for the site.").
			Required().
			String()
	storeUid = storeCommand.
			Flag("uid", "An optional username or other identifier for the account this key is for. Useful if you have multiple accounts on a single site.").
			Short('u').
			String()
	generateCommand = app.
			Command("generate", "Generate OTP(s) for a site.").
			Alias("g")
	generateSite = generateCommand.
			Arg("site", "The site to generate OTP(s) for.").
			Required().
			String()
	generateUid = generateCommand.
			Flag("uid", "An optional identifier for the key to use. Only necessary if you have multiple keys for a site.").
			Short('u').
			String()
	deleteCommand = app.
			Command("delete", "Generate OTP(s) for a site.").
			Alias("d")
	deleteSite = deleteCommand.
			Arg("site", "The site to delete.").
			Required().
			String()
	deleteUid = deleteCommand.
			Flag("uid", "An optional identifier for the key to delete. Only necessary if you have multiple keys for a site.").
			Short('u').
			String()
)

type Site struct {
	UIDs map[string]string `json:"uids"`
}

type Keychain struct {
	Sites map[string]Site `json:"sites"`
}

func init() {
	store.Init("gotp")
}

func loadConfig() Keychain {
	var keychain Keychain
	err := store.Load("keychain.json", &keychain)
	if err != nil {
		panic(err)
	}
	return keychain
}

func saveConfig(keychain *Keychain) {
	err := store.Save("keychain.json", &keychain)
	if err != nil {
		panic(err)
	}
}

func main() {
	keychain := loadConfig()
	if len(keychain.Sites) < 1 {
		keychain.Sites = make(map[string]Site)
	}
	if *debug {
		fmt.Println("Debug mode enabled.")
	}
	appVersion := fmt.Sprintf("%d.%d.%d (%d)", AppVersionMajor, AppVersionMinor,
		AppVersionPatch, AppVersionDate)
	app.Version(appVersion)
	switch kingpin.MustParse(app.Parse(os.Args[1:])) {
	case storeCommand.FullCommand():
		if *storeUid != "" {
			cmdStoreWithUID(keychain, *storeSite, *storeKey, *storeUid)
		} else {
			cmdStore(keychain, *storeSite, *storeKey)
		}
	case generateCommand.FullCommand():
		if *generateUid != "" {
			cmdGenerateWithUID(keychain, *generateSite, *generateUid)
		} else {
			cmdGenerate(keychain, *generateSite)
		}
	case deleteCommand.FullCommand():
		if *deleteUid != "" {
			cmdDeleteWithUID(keychain, *deleteSite, *deleteUid)
		} else {
			cmdDelete(keychain, *deleteSite)
		}
	}
	saveConfig(&keychain)
}
