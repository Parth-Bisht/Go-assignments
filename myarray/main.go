package main

import (
	"fmt"
	// "sort"
)

func main() {
	// fmt.Println("Arrays")

	// var fruitList [4] string
	// fruitList[0] = "Apple"
	// fruitList[1] = "Tomato"
	// fruitList[2] = "Peach"
	// fruitList[3] = "Banana"

	// fmt.Println("Fruit list is: ",fruitList)
	// fmt.Println("Fruit list is: ",len(fruitList))

	// var vegList = [3] string{"potatoe","beans","mushroom"}
	// fmt.Println("Vegy list is: ",len(vegList))

	// fmt.Println("Slices")

	// var fruitList = []string{"Apple", "Tomato", "Peach"}
	// fmt.Printf("Type of Fruitlsit is %T\n", fruitList)

	// fruitList = append(fruitList, "Mango", "Banana")

	// fmt.Println(fruitList)

	// fruitList = append(fruitList[:3])
	// fmt.Println(fruitList)

	// highScores := make([]int, 4)

	// highScores[0] = 234
	// highScores[1] = 345
	// highScores[2] = 987
	// highScores[3] = 467

	// highScores = append(highScores, 555, 666, 321)

	// fmt.Println(highScores)

	// fmt.Println("Is it sorted?",sort.IntsAreSorted(highScores))
	// sort.Ints(highScores)
	// fmt.Println(highScores)

	// //how to remove a value from slices based on index
	// var courses = [] string{"reactjs","js","swift","python","ruby"}
	// fmt.Println(courses)
	// var index int = 2
	// courses = append(courses[:index],courses[index+1:]...)
	// fmt.Println(courses)

	languages := make(map[string]string)

	languages["JS"] = "Javascript"
	languages["RB"] = "Ruby"
	languages["Py"] = "Python"

	fmt.Println("List of all languages: ",languages)
	fmt.Println("JS shorts for: ",languages["JS"])

	delete(languages,"RB")
	fmt.Println("List of all languages: ",languages)

	for key,value := range languages{
		fmt.Printf("For key %v, vlaue is %v\n",key,value)
	}
}
