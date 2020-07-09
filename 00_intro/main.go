package main

import "fmt"

func main() {
	var publisher, writer, artist, title string
	var year, pageNumber int
	var grade float32

	title = "Mr. GoToSleep"
	writer = "Tracey Hatchet"
	artist = "Jewel Tampson"
	publisher = "DizzyBooks Publishing Inc."
	year = 1997
	pageNumber = 14
	grade = 6.5

	fmt.Println(title + " written by : " + writer + ", published by : " + publisher + ", drawn by " + artist)
	fmt.Println("in the year ", year, ". Pages : ", pageNumber, " for grade : ", grade)

	title = "Epic Vol. 1"
	writer = "Ryan N. Shawn"
	artist = "Pheobe Paperclips"
	publisher = "DizzyBooks Publishing Inc."
	year = 2013
	pageNumber = 160
	grade = 9.5

	fmt.Println(title + " written by : " + writer + ", published by : " + publisher + ", drawn by " + artist)
	fmt.Println("in the year ", year, ". Pages : ", pageNumber, " for grade : ", grade)

}
