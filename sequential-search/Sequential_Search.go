package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"time"
)

func sequential_search(array []int, size int, nilai int) bool {
	ketemu := false

	for i := 0; i < size && !ketemu; {
		if array[i] == nilai {
			ketemu = true
		} else {
			i++
		}
	}
	return ketemu

}

func print_array(array []int, size int) {
	fmt.Printf("[")

	for i := 0; i < size; i++ {
		fmt.Printf("%d", array[i])

		if i != size-1 {
			fmt.Printf(", ")
		}
	}
	fmt.Printf("]\n")

}

func baca_file() []int {
	file, err := os.Open("data.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	var s []int
	// printSlice(s)
	scanner := bufio.NewScanner(file) // inisialisasi membaca file
	for scanner.Scan() {              // loop reading datanya
		lineStr := scanner.Text()       // ambil 1 data
		num, _ := strconv.Atoi(lineStr) // ubah ke int
		s = append(s, num)
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	return s
}

func main() {
	now := time.Now()
	var data = baca_file()
	defer func() {
		fmt.Println("\nEksekusi waktu ", time.Now().Sub(now))
		fmt.Printf("Dari jumlah %d data\n", len(data))
	}()
	fmt.Println("================= ")
	fmt.Println("Sequential Search")
	fmt.Println("================= ")
	fmt.Println("Masukkan angka yang dicari: ")
	var cari int
	fmt.Scanln(&cari)
	// fmt.Printf("Isi array: \n")
	// print_array(data, len(data)) // tampilkan data jika mau
	status_200 := sequential_search(data, len(data), cari)
	fmt.Printf("\nNilai %d %t\n", cari, status_200)
}
