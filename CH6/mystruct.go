package CH6

import (
	"image/color"
	"log"
)

type ColoredPoint struct {
	color.Color
	x,y int
}

func mystruct() {
	var c ColoredPoint
	c.x=3
	c.y=4
	log.Println(c.Color)
}