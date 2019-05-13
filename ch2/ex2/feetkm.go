package ex2

import "fmt"

type Foot float64

type Meter float64

func FootToMeter(f Foot) Meter{return Meter(f * 0.384)}

func MeterToFoot(m Meter)Foot{return Foot(m/0.384)}

func (f Foot)String()string{
	return fmt.Sprintf("%gfeet",f)
}

func (m Meter)String()string{
	return fmt.Sprintf("%gmeter", m)
}