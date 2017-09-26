package main

import (

"image"
"image/color"
"image/png"
"os"
)

var img = image.NewRGBA(image.Rect(0, 0, 2000, 1000))
var col color.Color

func main() {
col = color.Black
Rect(300, 300,500 , 500,)

f, err := os.Create("draw.png")
if err != nil {
panic(err)
}
defer f.Close()
png.Encode(f, img)
}

// HLine draws a horizontal line
func HLine(x1, y, x2 int) {
for ; x1 <= x2; x1++ {
img.Set
}
}

// VLine draws a veritcal line
func VLine(x, y1, y2 int) {
for ; y1 <= y2; y1++ {
img.Set(x, y1, col)
}
}

func Rect(x1, y1, x2, y2,x3,y3,x4,y4 int) {
HLine(x1, y1, x2,y2)
HLine(x2,y2,x3,y3)
VLine(x3,y3,x4,y4)
VLine(x4,y4,x1,y1)
	grap
}