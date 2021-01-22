package main

import "github.com/mallvielfrass/coloredPrint/fmc"

//	yst    FgYellow
//	ybt    FgYellow, Bold

//	rst    FgRed
//	rbt    FgRed, Bold

//	gst    FgGreen
//	gbt    FgGreen, Bold

//	bst    FgBlue
//	bbt    FgBlue, Bold

//	wst    FgWhite
//	wbt    FgWhite, Bold
func main() {
	colors := [10]string{"yst", "ybt", "rst", "rbt", "gst", "gbt", "bst", "bbt", "wst", "wbt"}
	for _, value := range colors {
		fmc.Printf("color: #"+value+value+" %s\n", "test")
	}

	fmc.Printf("#ybt I#rbt  love#bbt  Go#gbt!\n")
	fmc.Printf("#ybt %s#rbt  %s#bbt  %s#gbt%s\n", "I", "love", "Go", "!")
	fmc.Print("#ybt lol")
	fmc.Printfln("#ybt  %s", "I")
	fmc.Println("#bbt Println")

}
