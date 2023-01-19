package main

import "fmt"

func main() {
	var (
		day1, day2, month1, month2           string
		date1, date2, year1, year2, mo1, mo2 int
	)

	fmt.Scanln(&day1, &date1, &month1, &year1)
	mo1 = angkaBulan(month1)
	retrieval(date1, mo1, year1, day1, &date2, &mo2, &year2, &day2)
	month2 = bulanAngka(mo2)
	fmt.Println(day2, date2, month2, year2)

}

func kabisat(tahun int) bool {
	if tahun%400 == 0 || (tahun%4 == 0 && tahun%100 != 0) {
		return true
	} else {
		return false
	}
}

func angkaBulan(bulan string) int {
	var result int
	if bulan == "januari" {
		result = 1
	} else if bulan == "februari" {
		result = 2
	} else if bulan == "maret" {
		result = 3
	} else if bulan == "april" {
		result = 4
	} else if bulan == "mei" {
		result = 5
	} else if bulan == "juni" {
		result = 6
	} else if bulan == "juli" {
		result = 7
	} else if bulan == "agustus" {
		result = 8
	} else if bulan == "september" {
		result = 9
	} else if bulan == "oktober" {
		result = 10
	} else if bulan == "november" {
		result = 11
	} else if bulan == "desember" {
		result = 12
	}
	return result
}

func bulanAngka(angka int) string {
	var result string
	if angka == 1 {
		result = "januari"
	} else if angka == 2 {
		result = "februari"
	} else if angka == 3 {
		result = "maret"
	} else if angka == 4 {
		result = "april"
	} else if angka == 5 {
		result = "mei"
	} else if angka == 6 {
		result = "juni"
	} else if angka == 7 {
		result = "juli"
	} else if angka == 8 {
		result = "agustus"
	} else if angka == 9 {
		result = "september"
	} else if angka == 10 {
		result = "oktober"
	} else if angka == 11 {
		result = "november"
	} else if angka == 12 {
		result = "deseember"
	}
	return result
}

func jumlahHari(bulan, tahun int) int {
	var result int
	if bulan == 1 {
		result = 31
	} else if bulan == 2 {
		if kabisat(tahun) {
			result = 29
		} else {
			result = 28
		}
	} else if bulan == 3 {
		result = 31
	} else if bulan == 4 {
		result = 30
	} else if bulan == 5 {
		result = 31
	} else if bulan == 6 {
		result = 30
	} else if bulan == 7 {
		result = 31
	} else if bulan == 8 {
		result = 31
	} else if bulan == 9 {
		result = 30
	} else if bulan == 10 {
		result = 31
	} else if bulan == 11 {
		result = 30
	} else if bulan == 12 {
		result = 31
	}
	return result
}

func mencariDurasi(hari1 string, hari2 *string, durasi *int) {
	if hari1 == "senin" {
		*hari2 = "rabu"
		*durasi = 2
	} else if hari1 == "selasa" {
		*hari2 = "kamis"
		*durasi = 2
	} else if hari1 == "rabu" {
		*hari2 = "jumat"
		*durasi = 2
	} else if hari1 == "kamis" {
		*hari2 = "senin"
		*durasi = 4
	} else if hari1 == "jumat" {
		*hari2 = "selasa"
		*durasi = 4
	} else if hari1 == "sabtu" {
		*hari2 = "selasa"
		*durasi = 3
	} else if hari1 == "minggu" {
		*hari2 = "selasa"
		*durasi = 2
	}
}

func retrieval(tanggal1, bulan1, tahun1 int, hari1 string, tanggal2, bulan2, tahun2 *int, hari2 *string) {

	var duration, totalday int

	duration = 0
	mencariDurasi(hari1, &*hari2, &duration)
	totalday = jumlahHari(bulan1, tahun1)
	*tanggal2 = tanggal1 + duration
	*bulan2 = bulan1
	*tahun2 = tahun1

	if *tanggal2 > totalday {
		*tanggal2 = *tanggal2 - totalday
		*bulan2 = *bulan2 + 1
	}

	if *bulan2 > 12 {
		*tahun2 = *tahun2 + 1
		*bulan2 = *bulan2 - 12
	}

}
