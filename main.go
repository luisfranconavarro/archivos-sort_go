package main

import(
	"fmt"
  	"os"
	"strings"
	"sort"
)

type asc []string

func (a asc) Len() int           { return len(a) }
func (a asc) Swap(i, j int)      { a[i], a[j] = a[j], a[i] }
func (a asc) Less(i, j int) bool { return a[i] < a[j] }

type des []string

func (a des) Len() int           { return len(a) }
func (a des) Swap(i, j int)      { a[i], a[j] = a[j], a[i] }
func (a des) Less(i, j int) bool { return a[i] > a[j] }

func main(){

	file, err := os.Create("asacendente.txt")
	if err != nil{
		fmt.Println(err)
		return
	}
	defer file.Close()

	file2, err := os.Create("descendente.txt")
	if err != nil{
		fmt.Println(err)
		return
	}
	defer file2.Close()

	slicePal := []string{}
	var palabra, separador string

	for i := 0; i < 5; i++ {
		fmt.Print("ingresa una palabra: ")
		fmt.Scan(&palabra)
		slicePal = append(slicePal, palabra)
	}

	sort.Sort(asc(slicePal))
	separador = strings.Join(slicePal,"\n")
	file.WriteString(separador)

	sort.Sort(des(slicePal))
	separador = strings.Join(slicePal,"\n")
	file2.WriteString(separador)

}