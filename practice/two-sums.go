package main

import (
	"fmt"
)

// findPairs mencari semua pasangan unik dari array yang jumlahnya sama dengan target.
//
// Parameter:
//   - arr: slice integer yang akan dicari pasangannya
//   - target: nilai target yang harus dicapai dari penjumlahan dua angka
//
// Return:
//   - slice of slice integer yang berisi pasangan-pasangan unik yang jumlahnya sama dengan target
//
// Cara kerja:
//  1. Membuat map 'seen' untuk menyimpan angka yang sudah dilihat
//  2. Membuat map 'pairs' untuk menghindari duplikasi pasangan
//  3. Membuat slice 'result' untuk menyimpan hasil akhir
//  4. Iterasi setiap angka dalam array:
//     - Hitung nilai yang dibutuhkan (need = target - num)
//     - Jika nilai yang dibutuhkan sudah ada di 'seen':
//     * Urutkan pasangan (a, b) agar yang lebih kecil di depan
//     * Buat key unik dari pasangan tersebut (format: "a,b")
//     * Jika pasangan belum ada di 'pairs', tambahkan ke result
//     - Tandai angka saat ini sebagai sudah dilihat
//  5. Kembalikan semua pasangan yang ditemukan
//
// Contoh:
//
//	findPairs([]int{1, 2, 3, 4, 5}, 5) // returns [][]int{{1, 4}, {2, 3}}

func findPairs(arr []int, target int) [][]int {
	seen := make(map[int]bool)
	pairs := make(map[string]bool)
	result := [][]int{}

	for _, num := range arr {

		// Hitung angka yang dibutuhkan untuk mencapai target
		need := target - num

		println("Current number:", num, "Needs:", need)

		// Cek apakah angka yang dibutuhkan sudah pernah dilihat
		if seen[need] {
			a, b := num, need
			println("Found pair:", a, b)

			// Urutkan pasangan agar yang lebih kecil di depan
			if a > b {
				a, b = b, a
			}

			// Buat key unik untuk pasangan
			key := fmt.Sprintf("%d,%d", a, b)

			println("Pair key:", key)

			// Cek apakah pasangan sudah ada untuk menghindari duplikasi
			if !pairs[key] {
				pairs[key] = true
				result = append(result, []int{a, b})
			}
		}
		seen[num] = true

	}
	return result

}

func findPairsP(arr []int, target int) [][]int {
	//sort.Ints(arr) // O(n log n)

	// Gunakan dua pointer untuk menemukan pasangan
	result := [][]int{}
	l, r := 0, len(arr)-1

	// O(n)
	for l < r {
		// Hitung jumlah dari dua pointer
		sum := arr[l] + arr[r]

		// Jika jumlah sesuai dengan target, simpan pasangan dan pindahkan kedua pointer
		if sum == target {
			result = append(result, []int{arr[l], arr[r]})

			// skip duplicate kiri
			leftVal := arr[l]
			for l < r && arr[l] == leftVal {
				l++
			}

			// skip duplicate kanan
			rightVal := arr[r]
			for l < r && arr[r] == rightVal {
				r--
			}

			// Jika jumlah kurang dari target, geser pointer kiri ke kanan
		} else if sum < target {
			l++
		} else {
			r--
		}
	}
	return result
}

func main() {
	arr := []int{1, 2, 3, 4, 5, 6}
	target := 7

	pairs := findPairs(arr, target)
	fmt.Println(pairs)

	pairsP := findPairsP(arr, target)
	fmt.Println(pairsP)
}
