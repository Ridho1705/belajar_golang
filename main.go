package main

import "fmt"

func main() {
	fmt.Println("Hello, 世界")
	nama :="ridho"
	if nama=="ridho"{
		fmt.Println("nama saya",nama)

	}else {
		fmt.Println("bukan",nama)
		}
	list:=[]int64{1,2,3,4}
	for i :=0;i<=len(list)-1;{
		fmt.Println(list[i])
		i=i+1
	}
	for _,v:=range(list){
		fmt.Println(v)
	}
}
