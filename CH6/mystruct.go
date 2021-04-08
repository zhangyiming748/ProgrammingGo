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

type Optioner interface {
	Name() string
	IsValid() string
}
type OptionCommon struct {
	ShortName string "short option name"
	LongName string "kong option name"
}
type IntOption struct {
	OptionCommon
	Value,min,max int
}
func(option IntOption)Name()string{
	return name(option.ShortName,option.LongName)
}
func (option IntOption)IsValid()bool  {
	return option.min<=option.Value&&option.Value<=option.max
}
func name(shortName,longName string) string{
	if longName==""{
		return shortName
	}
	return longName
}