package problem605

import (
	"testing"
	"fmt"
)

func Test_canPlaceFlowers(t *testing.T) {
	flowerbed := []int{1, 0, 0, 0, 1}

	fmt.Println(canPlaceFlowers(flowerbed, 1))
	fmt.Println(canPlaceFlowers(flowerbed, 2))
}
