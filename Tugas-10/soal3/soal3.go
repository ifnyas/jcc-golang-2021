package soal3

import "strconv"

func LuasPersegi(sisi uint, tampilkan bool) interface{} {
	switch {
	case sisi == 0 && !tampilkan:
		return nil
	case sisi == 0 && tampilkan:
		return "Maaf anda belum menginput sisi dari persegi"
	case !tampilkan:
		return sisi * sisi
	default:
		return "luas persegi dengan sisi " + strconv.Itoa(int(sisi)) + "cm adalah " + strconv.Itoa(int(sisi*sisi)) + " cm"
	}
}
