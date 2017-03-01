package main

import (
	"log"

	"github.com/gizak/termui"
)

func main() {
	err := termui.Init()
	if err != nil {
		log.Fatal(err)
	}
	defer termui.Close()

	//termui.UseTheme("helloworld")

	bc := termui.NewMBarChart()
	math := []int{90, 85, 90, 80}
	english := []int{70, 85, 75, 60}
	science := []int{75, 60, 80, 85}
	compsci := []int{100, 100, 100, 100}
	bc.Data[0] = math
	bc.Data[1] = english
	bc.Data[2] = science
	bc.Data[3] = compsci
	studentsName := []string{"12", "13", "14", "15"}
	bc.ShowScale = true
	bc.BorderLabel = "Student's Marks X-Axis=Name Y-Axis=Marks[Math,English,Science,ComputerScience] in %"
	bc.BarWidth = 20
	bc.Width = 50
	bc.Height = 30
	bc.Y = 3
	bc.BarWidth = 10
	bc.DataLabels = studentsName
	bc.ShowScale = true //Show y_axis scale value (min and max)
	bc.TextColor = termui.ColorGreen
	bc.SetMax(200)
	bc.Align()

	bc.TextColor = termui.ColorGreen    //this is color for label (x-axis)
	bc.BarColor[3] = termui.ColorGreen  //BarColor for computerscience
	bc.BarColor[1] = termui.ColorYellow //Bar Color for english
	bc.NumColor[3] = termui.ColorRed    // Num color for computerscience
	bc.NumColor[1] = termui.ColorRed    // num color for english

	//Other colors are automatically populated, btw All the students seems do well in computerscience. :p

	g0 := termui.NewGauge()
	g0.Percent = 40
	g0.Width = 50
	g0.Height = 3
	g0.BorderLabel = "Slim Gauge"
	g0.BarColor = termui.ColorRed
	g0.BorderFg = termui.ColorWhite
	g0.BorderLabelFg = termui.ColorCyan

	bc1 := termui.NewBarChart()
	data := []int{3, 2, 5, 3, 9, 5, 3, 2, 5, 8, 3, 2, 4, 5, 3, 2, 5, 7, 5, 3, 2, 6, 7, 4, 6, 3, 6, 7, 8, 3, 6, 4, 5, 3, 2, 4, 6, 4, 8, 5, 9, 4, 3, 6, 5, 3, 6}
	bclabels := []string{"S0", "S1", "S2", "S3", "S4", "S5"}
	bc1.BorderLabel = "Bar Chart"
	bc1.Data = data
	bc1.Y = 33
	bc1.Width = 26
	bc1.Height = 10
	bc1.DataLabels = bclabels
	bc1.TextColor = termui.ColorGreen
	bc1.BarColor = termui.ColorRed
	bc1.NumColor = termui.ColorYellow

	termui.Render(bc, bc1, g0)

	termui.Handle("/sys/kbd/q", func(termui.Event) {
		termui.StopLoop()
	})
	termui.Loop()
}
