package main

import (
	"fmt"
)

func cmdDelete(site string) {
	fmt.Printf("-- in cmdDelete(%s)\n", site)
	cmdDeleteWithUID(site, "__default")
}

func cmdDeleteWithUID(site string, uid string) {
	fmt.Printf("-- in cmdDeleteWithUID(%s, %s)\n", site, uid)
}