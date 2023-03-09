package main

import (
	"fmt"
	"math/rand"
	"time"
)

func remove(slice []int, index int) []int {
	// fmt.Println("yang dihapus index ke -", index, " dari ", slice)
	copy(slice[index:], slice[index+1:])
	return slice[:len(slice)-1]
}

func main() {
	// Input jumlah pemain dan jumlah dadu
	var n, m int
	fmt.Print("Masukkan jumlah pemain: ")
	fmt.Scanln(&n)
	fmt.Print("Masukkan jumlah dadu: ")
	fmt.Scanln(&m)
	fmt.Println("Pemain=", n, ", Dadu=", m)
	fmt.Println("====================")
	// Membuat slice untuk menampung dadu pemain dan poin masing-masing pemain
	dice := make([][]int, n)
	points := make([]int, n)
	tempPoint := make(map[int]int, n)
	player := []int{}
	// Memasukkan dadu awal untuk setiap pemain
	rand.Seed(time.Now().UnixNano())
	for i := 0; i < n; i++ {
		dice[i] = make([]int, m)
		for j := 0; j < m; j++ {
			dice[i][j] = rand.Intn(6) + 1
		}
		tempPoint[i+1] = 0
		player = append(player, i+1)
	}
	// fmt.Println(tempPoint)
	// Memulai permainan
	round := 1
	for n > 1 {
		fmt.Printf("Giliran %d lempar dadu:\n", round)

		// Setiap pemain melemparkan dadu
		for i := 0; i < n; i++ {
			for j := 0; j < m; j++ {
				dice[i][j] = rand.Intn(6) + 1
			}
			fmt.Printf("\tPemain #%d melempar dadu: %v\n", player[i], dice[i])
		}

		// Evaluasi hasil lemparan dadu
		for i := 0; i < n; i++ {
			for j := 0; j < m; j++ {
				if dice[i][j] == 6 {
					points[i]++
					dice[i][j] = 0
				} else if dice[i][j] == 1 {
					// Memberikan dadu 1 ke pemain disampingnya
					if i == n-1 {
						points[0]++
					} else {
						points[i+1]++
					}
					dice[i][j] = 0
				}
			}
			// Menghitung jumlah dadu yang masih dimiliki oleh pemain
			remainingDice := 0
			for j := 0; j < m; j++ {
				if dice[i][j] != 0 {
					remainingDice++
				}
			}
			if remainingDice == 0 {
				fmt.Printf("\t(Pemain %d berhenti karena tidak memiliki dadu lagi)\n", player[i])
				player = remove(player, i)
			}
		}

		// Menghapus pemain yang sudah tidak memiliki dadu
		newDice := make([][]int, 0)
		newPoints := make([]int, 0)
		for i := 0; i < n; i++ {
			remainingDice := 0
			for j := 0; j < m; j++ {
				if dice[i][j] != 0 {
					remainingDice++
				}
			}
			if remainingDice > 0 {
				newDice = append(newDice, dice[i])
				newPoints = append(newPoints, points[i])
			}
		}
		dice = newDice
		points = newPoints
		n = len(dice)

		// Menampilkan poin setiap pemain
		fmt.Println("Setelah evaluasi")
		for i := 0; i < n; i++ {
			fmt.Printf("\tPemain %d: %d poin\n", player[i], points[i])
			tempPoint[player[i]] = points[i]
		}
		round++
		// fmt.Println(points)
		// fmt.Println(player)
	}
	// fmt.Println(tempPoint)
	// Hasil Akhir
	winner := []int{0, 0}
	for i, p := range tempPoint {
		if p > winner[1] {
			winner[1] = p
			winner[0] = i
		}
	}
	fmt.Println("RESULT :")
	for i, hasil := range tempPoint {
		fmt.Println("Pemain #", i, " = ", hasil, " point")
	}
	fmt.Println("")
	if winner[0] == 0 && winner[1] == 0 {
		fmt.Println("Semua pemain tidak mendapatkan point sehingga DRAW")
	} else {
		fmt.Printf("Pemain %d memenangkan permainan dengan %d poin!\n", winner[0], winner[1])
	}
}
