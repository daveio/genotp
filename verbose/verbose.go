package verbose

import (
	"fmt"
	"log"
)

var (
	verbose bool = true
	appVersion string = "VER?"
)

func InitV(setVerbose bool, setAppVersion string) {
	verbose = setVerbose
	appVersion = setAppVersion
}

func V(logLine string) {
	if verbose {
		prefix := fmt.Sprintf("--- gotp-%s ", appVersion)
		log.SetPrefix(prefix)
		log.Println(logLine)
	}
}