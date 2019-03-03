package main

import (
	"fmt"
)

func cmdGenerate(site string) {
	fmt.Printf("-- in cmdGenerate(%s)\n", site)
	cmdGenerateWithUID(site, "__default")
}

func cmdGenerateWithUID(site string, uid string) {
	fmt.Printf("-- in cmdGenerateWithUID(%s, %s)\n", site, uid)
}