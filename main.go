package main

import (
	"fmt"
	"github.com/tucnak/store"
	"gopkg.in/alecthomas/kingpin.v2"
	"os"
	"time"
)

var (
	app = kingpin.
		New("genotp", "Generate OATH-TOTP one-time passwords from the command line.")
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
	Site string `yaml:"site"`
	UID  string `yaml:"uid"`
	Key  string `yaml:"key"`
}

type Keychain struct {
	Accounts []Account `yaml:"accounts"`
}

func init() {
	store.Init("genotp")
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
	app.Version("0.0.1")
	switch kingpin.MustParse(app.Parse(os.Args[1:])) {
	case storeCommand.FullCommand():
		fmt.Printf("Storing --> ")
		fmt.Printf("Site: %s ", *storeSite)
		if *storeUid != "" {
			fmt.Printf("UID: %s ", *storeUid)
		}
		fmt.Printf("Key: %s", *storeKey)
		fmt.Println()
	case generateCommand.FullCommand():
		fmt.Printf("Generating --> ")
		fmt.Printf("Site: %s ", *generateSite)
		if *generateUid != "" {
			fmt.Printf("UID: %s", *generateUid)
		}
		fmt.Println()
	case deleteCommand.FullCommand():
		fmt.Printf("Deleting --> ")
		fmt.Printf("Site: %s ", *deleteSite)
		if *deleteUid != "" {
			fmt.Printf("UID: %s", *deleteUid)
		}
		fmt.Println()
	}
	time.Sleep(10 * time.Second)
	saveConfig(&keychain)
}
