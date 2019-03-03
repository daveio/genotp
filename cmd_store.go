package main

import (
	"fmt"
)

func cmdStore(site string, key string) {
	fmt.Printf("-- in cmdStore(%s, %s)\n", site, key)
	cmdStoreWithUID(site, key, "__default")
}

func cmdStoreWithUID(site string, key string, uid string) {
	fmt.Printf("-- in cmdStoreWithUID(%s, %s, %s)\n", site, key, uid)
}