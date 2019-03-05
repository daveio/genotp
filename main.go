package main

import (
	"fmt"
	"github.com/daveio/gotp/commands"
	"github.com/daveio/gotp/storage"
	"github.com/daveio/gotp/verbose"
	"gopkg.in/alecthomas/kingpin.v2"
	"os"
)

var (
	V = verbose.V
	appVersion = fmt.Sprintf("%d.%d.%d", AppVersionMajor, AppVersionMinor, AppVersionPatch)
	appVersionWithDate = fmt.Sprintf("%s %d", appVersion, AppVersionDate)
	app = kingpin.
		New("gotp", "Generate OATH-TOTP one-time passwords from the command line.")
	mVerbose = app.
		Flag("verbose", "Show more detail.").
		Short('v').
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
/*
	storeURICommand = app.
		Command("store-uri", "Store a new account using a totp:// URI.").
		Alias("su")
	storeURIURI = storeURICommand.
		Arg("uri", "A valid totp:// URI").
		Required().
		String()
*/
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

func init() {
	storage.InitStorage("gotp")
}

func main() {
	app.Version(appVersionWithDate)
	appCmd, appErr := app.Parse(os.Args[1:])
	verbose.InitV(*mVerbose, appVersion)
	V("Verbose mode enabled")
	keychain := storage.LoadConfig()
	if len(keychain.Sites) < 1 {
		keychain.Sites = make(map[string]storage.Site)
	}
	switch kingpin.MustParse(appCmd, appErr) {
	case storeCommand.FullCommand():
		V("storing a new item")
		if *storeUid != "" {
			commands.Store(keychain, *storeSite, *storeKey, *storeUid)
		} else {
			commands.Store(keychain, *storeSite, *storeKey, "__default")
		}
	case generateCommand.FullCommand():
		V("generating an OTP")
		if *generateUid != "" {
			commands.Generate(keychain, *generateSite, *generateUid)
		} else {
			commands.Generate(keychain, *generateSite, "__default")
		}
	case deleteCommand.FullCommand():
		V("deleting an item")
		if *deleteUid != "" {
			commands.Delete(keychain, *deleteSite, *deleteUid)
		} else {
			commands.Delete(keychain, *deleteSite, "__default")
		}
/*	case storeURICommand.FullCommand():
		site, uid, key, err := parseURI(*storeURIURI)
		if err == nil {
			V(fmt.Sprintf("site: %s, uid: %s, key: %s", site, uid, key))
		} else {
			fmt.Println("Malformed URI")
		} */
	default:
		V("BUG: invalid command uncaught by CLI parser")
	}
	storage.SaveConfig(&keychain)
}
