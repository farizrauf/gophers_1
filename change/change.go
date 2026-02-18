package change

import (
	"fmt"

	"github.com/fatih/color"
)

func CheckChange() {
	var uang, harga int
	fmt.Print("Harga Barang: ")
	fmt.Scan(&harga)
	fmt.Print("Uang Pembeli: ")
	fmt.Scan(&uang)

	if uang < harga {
		color.Red("[SISTEM]: Transaksi Ditolak!. Uang Anda kurang %d", harga-uang)
	} else {
		color.Green("[SISTEM]: Transaksi Berhasil. Kembalian Anda adalah %d", uang-harga)
	}
}
