package problem599

import (
	"testing"
	"fmt"
)

func Test_findRestaurant(t *testing.T) {
	list1 := []string{"Shogun", "Tapioca Express", "Burger King", "KFC"}
	list2 := []string{"Piatti", "The Grill at Torrey Pines", "Hungry Hunter Steakhouse", "Shogun"}

	fmt.Println(findRestaurant(list1, list2))
}
