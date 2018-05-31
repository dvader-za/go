package main

import "fmt"

func main() {

	/*fmt.Println("Hello, earth!")
	numbers := []int{0, 1, 2, 3, 4, 5, 6, 7, 8}
	num := common.Reverse(numbers)

	for i, val := range num {
		fmt.Println("Slice item", i, "is", val)
	}

	var countryCapitalMap map[string]string*/
	/* create a map*/
	/*countryCapitalMap = make(map[string]string)

	/* insert key-value pairs in the map*/
	/*countryCapitalMap["France"] = "Paris"
	countryCapitalMap["Italy"] = "Rome"
	countryCapitalMap["Japan"] = "Tokyo"
	countryCapitalMap["India"] = "New Delhi"
	for key, value := range countryCapitalMap {
		fmt.Println("country - ", key, " capital ", value)
	}

	if _, ok := countryCapitalMap["Italy"]; ok {
		fmt.Println("Found")
	} else {
		fmt.Println("Not found")
	}*/

	var a string
	a, _ = swap("Stephen", "Booth")
	fmt.Println(a)

}

func swap(x, y string) (string, string) {
	return y, x
}
