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
	V                  = verbose.V
	appVersion         = fmt.Sprintf("%d.%d.%d", AppVersionMajor, AppVersionMinor, AppVersionPatch)
	appVersionWithDate = fmt.Sprintf("%s %d", appVersion, AppVersionDate)
	app                = kingpin.
				New("gotp", "Generate OATH-TOTP one-time passwords from the command line.")
	mVerbose = app.
			Flag("verbose", "Show more detail.").
			Short('v').
			Bool()
	storeCommand = app.
			Command("store", "Short form: 's'. Store a new account.").
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
			Flag("uid", "An optional account identifier. Useful if you have multiple accounts on a site.").
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
			Command("generate", "Short form: 'g'. Generate OTP(s) for a site.").
			Alias("g")
	generateSite = generateCommand.
			Arg("site", "The site to generate OTP(s) for.").
			Required().
			String()
	generateUid = generateCommand.
			Flag("uid", "An optional account identifier. Useful if you have multiple accounts on a site.").
			Short('u').
			String()
	deleteCommand = app.
			Command("delete", "Short form: 'd'. Delete a site or account.").
			Alias("d")
	deleteSite = deleteCommand.
			Arg("site", "The site to delete.").
			Required().
			String()
	deleteUid = deleteCommand.
			Flag("uid", "An optional account identifier. Useful if you have multiple accounts on a site.").
			Short('u').
			String()
	listSitesCommand = app.
			Command("list-sites", "Short form: 'ls'. List the sites you have added keys for.").
			Alias("ls")
	listUidsCommand = app.
			Command("list-uids", "Short form: 'lu'. List the accounts you have added for a site.").
			Alias("lu")
	listUidsSite = listUidsCommand.
			Arg("site", "The site to list accounts for.").
			Required().
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
	keychain, errL := storage.LoadConfig()
	if errL != nil {
		panic(errL)
	}
	if len(keychain.Sites) < 1 {
		keychain.Sites = make(map[string]storage.Site)
	}
	switch kingpin.MustParse(appCmd, appErr) {
	case storeCommand.FullCommand():
		V("storing a new item")
		if *storeUid != "" {
			errSU := commands.Store(keychain, *storeSite, *storeKey, *storeUid)
			if errSU != nil {
				panic(errSU)
			}
		} else {
			errS := commands.Store(keychain, *storeSite, *storeKey, "__default")
			if errS != nil {
				panic(errS)
			}
		}
	case generateCommand.FullCommand():
		V("generating an OTP")
		if *generateUid != "" {
			errGU := commands.Generate(keychain, *generateSite, *generateUid)
			if errGU != nil {
				panic(errGU)
			}
		} else {
			errG := commands.Generate(keychain, *generateSite, "__default")
			if errG != nil {
				panic(errG)
			}
		}
	case deleteCommand.FullCommand():
		V("deleting something")
		if *deleteUid != "" {
			V("deleting a uid")
			errDU := commands.DeleteUID(keychain, *deleteSite, *deleteUid)
			if errDU != nil {
				panic(errDU)
			}
		} else {
			V("deleting a site")
			errD := commands.DeleteSite(keychain, *deleteSite)
			if errD != nil {
				panic(errD)
			}
		}
	case listSitesCommand.FullCommand():
		V("listing sites")
		errLS := commands.ListSites(keychain)
		if errLS != nil {
			panic(errLS)
		}
	case listUidsCommand.FullCommand():
		V(fmt.Sprintf("listing uids for site %s", *listUidsSite))
		errLU := commands.ListUids(keychain, *listUidsSite)
		if errLU != nil {
			panic(errLU)
		}
	// TODO URI parsing glue
	default:
		V("BUG: invalid command uncaught by CLI parser")
	}
	errSa := storage.SaveConfig(&keychain)
	if errSa != nil {
		panic(errSa)
	}
}
