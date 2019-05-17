package main

func bigger(x int, y int) int {
	if x > y {
		return x
	}
	return y

}

func swap(x int, y int) (int, int) {
	return y, x
}

func main() {

	res := bigger(3, 5)

	print(res)

	r2, r3 := swap(200, 100)

	print(r2, r3)

}
