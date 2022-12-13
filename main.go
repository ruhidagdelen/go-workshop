package main

import "fmt"

func main() {

	//i := 1
	//for i <= 3 {
	//	fmt.Println(i)
	//	i = i + 1
	//}
	//
	//for j := 7; j <= 9; j++ {
	//	fmt.Println(j)
	//}
	//
	//for {
	//	fmt.Println("loop")
	//	break
	//}
	//
	//for n := 0; n <= 5; n++ {
	//	if n%2 == 0 {
	//		continue
	//	}
	//	fmt.Println(n)
	//}

	nums := []int{2, 3, 4}
	sum := 0
	for _, num := range nums {
		sum += num
	}
	fmt.Println("sum:", sum)

	for i, num := range nums {
		if num == 3 {
			fmt.Println("index:", i)
		}
	}

	kvs := map[string]string{"a": "apple", "b": "banana"}
	for k, v := range kvs {
		fmt.Printf("%s -> %s\n", k, v)
	}

	for k := range kvs {
		fmt.Println("key:", k)
	}

	for i, c := range "go" {
		fmt.Println(i, c)
	}

}
