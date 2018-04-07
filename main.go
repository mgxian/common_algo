package main

import (
	"fmt"
	"math/rand"
	"time"

	"github.com/mgxian/common_algo/search"
)

func getRandomIntSlice(min, max, size int) (randomIntSlice []int) {
	if size <= 0 {
		return nil
	}

	s := rand.NewSource(time.Now().UnixNano())
	r := rand.New(s)
	for i := 0; i < size; i++ {
		randomInt := min + r.Intn(max-min)
		randomIntSlice = append(randomIntSlice, randomInt)
	}

	return randomIntSlice
}

func main() {
	data := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	fmt.Println(search.BinarySearchV1(data, 5, 0, 9))
	fmt.Println(search.BinarySearchV2(data, 5))
}
