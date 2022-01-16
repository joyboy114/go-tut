package main

import (
	"fmt"
	"reflect"
)

var b = 3
//a := 9

//constants name in capital
const PI = 3.14


//struct can be passed as variable similar to other types
type Person struct {
	name string
	age int
	job string
	salary int
}

func main() {
	var a string
	var b int 
	var c bool 
	fmt.Println(a) //default value is set for a = ""
	fmt.Println(b) //default value is set for b = 0
	fmt.Println(c) //default value is set for c = false
	var student1 string
	student1 = "Jonh"
	fmt.Println(student1)	
	//difference between var and := 
	//var can be inside and outside of function
	// := can be used only inside of function 
	//comment out := line outside main and get error
	
	//value unpacking
	var a1,b1,c1 int = 4,5,6
	var d1,e1 = 5, "Hello world"
	fmt.Println(a1,b1,c1);
	fmt.Println(d1,e1);
	fmt.Println(reflect.TypeOf(PI));
	
	//type view
	var i = 15.5
	var txt = "Hello World!"

	fmt.Printf("%v\n", i)
	fmt.Printf("%#v\n", i)
	fmt.Printf("%v%%\n", i)
	fmt.Printf("%T\n", i)

	fmt.Printf("%v\n", txt)
	fmt.Printf("%#v\n", txt)
	fmt.Printf("%T\n", txt)


	 var arr1 = [3]int{1,2,3}
	 arr2 := [5]int{4,5,6,7,8}
	 fmt.Println(arr1)
	 fmt.Println(arr2)
	 //var arr3 = [2]int{1,2,3} //this gives error : out of bound
	 
	 //var arr3 = [4]int{1,2,3} //this is allowed
	 arr5 := [5]int{1:10,2:40}
	 fmt.Println(arr5)
	 fmt.Println(len(arr5))
	 
	 fmt.Println(arr5[2:4])
	 fmt.Println(len(arr5[2:4]))
	 
	 var person1 Person
	 //var person2 Person
	 person1.name = "Alice"
	 person1.age = 32
	 person1.job = "Teacher"
	 person1.salary = 4000
	 
	 fmt.Println(person1)
	 
	 var aMap = map[string]string{"brand": "Ford", "model": "Mustang", "year": "1964"}
  	 bMap := map[string]int{"Oslo": 1, "Bergen": 2, "Trondheim": 3, "Stavanger": 4}
  	 
	 fmt.Printf("a\t%v\n", aMap)
 	 fmt.Printf("b\t%v\n", bMap)
 	 fmt.Println(aMap["brand"]);
 	 fmt.Println(aMap["brand1"]==""); //gives empty value
	 
	 
	 
	 
	 
	 
	 
	 
	 
	 
	 
	 
	 
	 
	 
	 
	 
	 
} 
