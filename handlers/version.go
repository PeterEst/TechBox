package handlers

import (
	"fmt"
	"techbox/config"
)

func HandleVersion() {
	fmt.Printf("Techbox v%s\n", config.GetVersion())
}
