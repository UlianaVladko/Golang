package main

func InsertXInMap(x int) {
	testmap := make(map[int]int, 0)
	for i := 0; i < x; i++ {
		testmap[i] = i
	}
}

func InsertXInMapInterface(x int) {
	testmap := make(map[interface{}]int, 0)
	for i := 0; i < x; i++ {
		testmap[i] = i
	}
}

func InsertXInSlice(x int) {
	testmap := make([]int, 0)
	for i := 0; i < x; i++ {
		testmap = append(testmap, i)
	}
}

func main() {
	//any code
}
