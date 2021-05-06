package main

import (
	"fmt"
	"log"
	"strconv"
	"strings"
	"syscall/js"
)

func registerCallbacks() {
	js.Global().Set("result_arr",js.FuncOf(result_arr))
}

func result_arr(this js.Value,v []js.Value) interface{}{
	log.Println("v: ",v[0])
	var str1= v[0].String()
	log.Print("str1: ",str1)
	var str_operation []string
	for _,iteam:=range str1{
		if iteam =='+'{
			str_operation =append(str_operation,"+")
			str1 = strings.Replace(str1, "+", ",", 1)
		}
		if iteam =='-'{
			str_operation =append(str_operation,"-")
			str1 = strings.Replace(str1, "-", ",", 1)
		}
		if iteam =='*'{
			str_operation =append(str_operation,"*")
			str1 = strings.Replace(str1, "*", ",", 1)
		}
		if iteam =='/'{
			str_operation =append(str_operation,"/")
			str1 = strings.Replace(str1, "/", ",", 1)
		}
	}

	str_number:= strings.Split(str1,",")
	var str_number1 = []int{}

	for _, i := range str_number {
		j, err := strconv.Atoi(i)
		if err != nil {
			panic(err)
		}
		str_number1 = append(str_number1, j)
	}
	var sum=0
	var i=0
	// đổi -A thành +(-A)
	for i=0;i<len(str_operation);i++{
		if str_operation[i]=="-"{
			str_operation[i]="+"
			str_number1[i+1]=str_number1[i+1]*(-1)
		}

	}
	for index, item := range(str_operation){
		if item == "/"{
			str_number1[index+1] = str_number1[index] / str_number1[index+1]
			str_number1[index] = 0
		}

		if item == "*"{
			str_number1[index+1] = str_number1[index] * str_number1[index+1]
			str_number1[index] = 0
		}

	}
	for _, item := range(str_number1){
		sum = sum + item
	}
	fmt.Println(str1)
	fmt.Println("str_operation",str_operation)
	fmt.Println("str_number1",str_number1)
	return sum
}

func main() {
	c := make(chan struct{}, 0)

	println("WASM Go Initialized")
	// register functions
	registerCallbacks()
	<-c
}