package main

import "fmt"

func main() {
	fmt.Println("here i'll try to implement smth with arrays...")
	var myArr [5]int
	myArr[3] = 33
	fmt.Println(myArr)

	var geru [35]string
	geru[1] = "geru geru"
	fmt.Println(geru)

	var blu [2][3]int
	blu[1][2] = 9
	blu[0][1] = 7
	fmt.Println(blu)

	var clu [4][2]string
	clu[1][1] = "fu fi"
	fmt.Print(clu)
	fmt.Println("end of line")
}
