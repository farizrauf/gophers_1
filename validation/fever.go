package validation

import (
	"fmt"

	"github.com/fatih/color"
)

func CheckFever() {
	var suhu int
	fmt.Println("Masukkan suhu badan Anda:")
	fmt.Scan(&suhu)

	if suhu > 37 {
		color.Red("Anda memiliki fever.")
	} else {
		color.Green("Anda tidak memiliki fever.")
	}
}