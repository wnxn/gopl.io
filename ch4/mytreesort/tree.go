package main

import (
	"fmt"
	"math/rand"
)

func main(){
	s := genSlice()
	fmt.Println(s)
	var t *treeNode
	for _,v:=range s{
		t = add(t, v)
	}
	midOrder(t)
}

func genSlice()[]int{
	res := make([]int,20)
	for i:=0;i<20;i++{
		res[i]=rand.Int()%20
	}
	return res
}

type treeNode struct{
	value int
	left, right *treeNode
}

func add(t *treeNode, i int)*treeNode{
	if t == nil{
		t = &treeNode{value:i}
		return t
	}
	if t.value > i{
		t.left = add(t.left,i)
	}else{
		t.right = add(t.right,i)
	}
	return t
}

func midOrder(t *treeNode){
	if t == nil{
		return
	}
	midOrder(t.left)
	fmt.Printf("%d,", t.value)
	midOrder(t.right)
}