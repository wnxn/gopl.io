package myreceiver

import "fmt"

type Point struct{
	x,y int
}

func NewPointPtr(x,y int)*Point{
	return &Point{x,y}
}

func NewPointValue(x,y int)Point{
	return Point{x,y}
}

func (p Point)ToString()string{
	return fmt.Sprintf("x=%d,y=%d", p.x, p.y)
}

func (p *Point)ScaleBy(factor int){
	if p == nil{
		return
	}
	p.x*=factor
	p.y*=factor
}


func PointEqual(p1,p2 *Point)bool{
	if p1 == nil && p2 == nil{
		return true
	}
	if p1 != nil && p2 != nil && p1.y == p2.y && p1.x == p2.x{
		return true
	}else{
		return false
	}
}

func (p Point)Add(factor int)Point{
	p.x += factor
	p.y += factor
	return p
}

type Chain []int

func (c Chain)Add(factor int)Chain{
	for i :=range c{
		c[i]+=factor
	}
	return c
}

func(c *Chain)ScaleBy(factor int){
	for i:=range *c{
		(*c)[i]*=factor
	}
}