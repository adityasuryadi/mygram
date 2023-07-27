package tests

import (
	"fmt"
	"log"
	"math"
	dbConfig "mygram/infrastructures/database"
	repository "mygram/infrastructures/repository/postgres"
	"sort"
	"testing"
)

func TestFindPhotoById(t *testing.T) {
	db:=dbConfig.NewTestPostgresDB()
	id := "aditsss@mail.com"
	photoRepository := repository.NewPhotoRepository(db)
	result,err := photoRepository.FindById(id)
	if err != nil {
		log.Print(result)
	}
	log.Print(result)
}


func TestBubleSort(t *testing.T){
	arr := []int{1,4,7,2,5,9}
	var tmp int
	pjg := len(arr)
	for i := 0; i < pjg; i++ {
		for j := pjg-1; j > i; j-- {
			fmt.Print(j)
			if arr[j] < arr[j-1] {
				tmp = arr[j] //1
				// fmt.Print("a",tmp)
				arr[j] = arr[j-1] //3
				// fmt.Print("b",arr[i])
				arr[j-1] = tmp
				// fmt.Print("c",arr[i-1])	
			}
		}
	}
	fmt.Print(arr)
}

func TestSort(t *testing.T){
	arr := []int{1,4,7,2,5,9}
	sort.Ints(arr)
	fmt.Print(arr)
}

func RoundUp(v int){

}

func TestGrade(t *testing.T){
	arr := []int{73,67,38,33}
	for _, v := range arr {
		if v >= 38 {
			x := math.Ceil((float64(v)/5)) * 5
			v = v
			if math.Abs(float64(int(x) - v)) < 3 {
				v = int(x)
			}
		}
		fmt.Println(v)
	}
	
	

}