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

type Account struct {
	Site string `json:"site"`
	UID  string `json:"uid"`
	Key  string `json:"key"`
}

type Keychain struct {
	Accounts []Account `json:"accounts"`
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
	if *debug {
		fmt.Println("Debug mode enabled.")
	}
	// fmt.Printf(keychain.Accounts[0].Site)
	appVersion := fmt.Sprintf("%d.%d.%d (%d)", AppVersionMajor, AppVersionMinor,
		AppVersionPatch, AppVersionDate)
	app.Version(appVersion)
	switch kingpin.MustParse(app.Parse(os.Args[1:])) {
	case storeCommand.FullCommand():
		if *storeUid != "" {
			cmdStoreWithUID(*storeSite, *storeKey, *storeUid)
		} else {
			cmdStore(*storeSite, *storeKey)
		}
	case generateCommand.FullCommand():
		if *generateUid != "" {
			cmdGenerateWithUID(*generateSite, *generateUid)
		} else {
			cmdGenerate(*generateSite)
		}
	case deleteCommand.FullCommand():
		if *deleteUid != "" {
			cmdDeleteWithUID(*deleteSite, *deleteUid)
		} else {
			cmdDelete(*deleteSite)
		}
	}
	saveConfig(&keychain)
}
