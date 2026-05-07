package condicionales

import (
	"fmt"
	"runtime"
)

func CondicionIF() {
	if os := runtime.GOOS; os == "windows" || os == "linux" {
		fmt.Println("Estas usando Windows.")
	} else {
		fmt.Println("Esta usando OS X")
	}

	switch os := runtime.GOOS; os {
	case "windows":
		fmt.Println("Estas usando Windows.")
	case "linux":
		fmt.Println("Estas usando Linux.")
	default:
		fmt.Println("Esta usando OS X \n os")
	}

	fmt.Println(runtime.GOOS)
}
