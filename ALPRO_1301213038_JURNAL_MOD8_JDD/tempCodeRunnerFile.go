package main

import "fmt"

func main() {
	var winner, player rune
	var ronde, nilai, answer int

	winner = 'A'
	player = 'B'
	ronde = 1

	mulaiRonde(ronde, winner, &nilai)
	for nilai != -101 {

		answer = tebakan(player, nilai)

		tukerPemenang(answer == nilai, &winner, &player)

		fmt.Printf("%c adalah pemenangnya\n\n", winner)
		ronde++
		mulaiRonde(ronde, winner, &nilai)
	}
	fmt.Println("Permainan selesai")
}

func tebakan(player rune, nilai int) int {
	var jawaban int
	i := 1
	fmt.Printf("%c - masukkan angka tebakan ke-%d :", player, i)
	fmt.Scanln(&jawaban)
	for i < 3 && jawaban != nilai {
		i++
		fmt.Printf("%c - masukkan angka tebakan ke-%d :", player, i)
		fmt.Scanln(&jawaban)
	}
	return jawaban
}

func tukerPemenang(benar bool, winner, player *rune) {
	if benar {
		var temp rune = *winner
		*winner = *player
		*player = temp
	}
}

func mulaiRonde(ronde int, winner rune, nilai *int) {
	fmt.Println("Ronde", ronde, ":")
	fmt.Printf("%c - masukkan angka rahasia: ", winner)
	fmt.Scanln(nilai)
}
