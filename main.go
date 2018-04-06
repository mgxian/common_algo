package main

import (
	"fmt"
	"math/rand"
	"time"

	"github.com/mgxian/common_algo/sort"
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
	randomIntSlice := getRandomIntSlice(1, 100, 10)
	fmt.Println(randomIntSlice)
	sort.BubbleSort(randomIntSlice)
	fmt.Println(randomIntSlice)
}
