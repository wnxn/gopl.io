package myreceiver_test

import (
	"fmt"
	"github.com/adonovan/gopl.io/ch6/myreceiver"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("Point Pointer Receiver", func() {
	Describe("Pass Pointer", func(){
		It("should change input value", func(){
			ptr:= myreceiver.NewPointPtr(4,6)
			ptr.ScaleBy(2)
			Expect(myreceiver.PointEqual(ptr,myreceiver.NewPointPtr(8,12))).To(Equal(true))
		})
	})

	Describe("Pass Value", func(){
		It("should change input value", func(){
			val := myreceiver.NewPointValue(4,6)
			val.ScaleBy(2)
			Expect(myreceiver.PointEqual(&val,myreceiver.NewPointPtr(8,12))).To(Equal(true))
		})
	})

	Describe("Pass nil", func() {
		It("should return zeor value", func(){
			var ptr *myreceiver.Point = nil
			ptr.ScaleBy(2)
			Expect(ptr).To(BeNil())
		})
	})
})

var _ = Describe("Point Value Receiver", func(){
	Describe("Pass Pointer", func(){
		It("should not change input value", func(){
			ptr := myreceiver.NewPointPtr(4,6)
			res := ptr.Add(2)
			Expect(myreceiver.PointEqual(ptr, myreceiver.NewPointPtr(4,6))).To(Equal(true))
			Expect(myreceiver.PointEqual(&res,myreceiver.NewPointPtr(6,8))).To(Equal(true))
		})
	})

	Describe("Pass Value", func(){
		It("should not change input value", func(){
			val := myreceiver.NewPointValue(4,6)
			res := val.Add(2)
			Expect(myreceiver.PointEqual(&val, myreceiver.NewPointPtr(4,6))).To(Equal(true))
			Expect(myreceiver.PointEqual(&res, myreceiver.NewPointPtr(6,8))).To(Equal(true))
		})
	})

	Describe("Pass temp value", func(){
		It("should not change input value", func(){
			res:= myreceiver.NewPointValue(4,6).Add(2)
			Expect(res).To(Equal(myreceiver.NewPointValue(6,8)))
			fmt.Fprintf(GinkgoWriter, "res = %v", res)
		})
	})
})


var _ = Describe("Chain Pointer Receiver", func() {
	Describe("Pass Pointer", func(){
		It("should change input value", func(){
			ptr:= &myreceiver.Chain{4,6}
			ptr.ScaleBy(2)
			Expect(ptr).To(Equal(&myreceiver.Chain{8,12}))
		})
	})

	Describe("Pass Value", func(){
		It("should change input value", func(){
			val := myreceiver.Chain{4,6}
			val.ScaleBy(2)
			Expect(val).To(Equal(myreceiver.Chain{8,12}))
		})
	})
})

var _ = Describe("Chain Value Receiver", func(){
	Describe("Pass Pointer", func(){
		It("should change input value", func(){
			ptr := &myreceiver.Chain{4,6}
			res := ptr.Add(2)
			Expect(ptr).To(Equal(&myreceiver.Chain{6,8}))
			Expect(res).To(Equal(myreceiver.Chain{6,8}))
		})
	})

	Describe("Pass Value", func(){
		It("should change input value", func(){
			val := myreceiver.Chain{4,6}
			res := val.Add(2)
			Expect(val).To(Equal(myreceiver.Chain{6,8}))
			Expect(res).To(Equal(myreceiver.Chain{6,8}))
		})
	})

	Describe("Pass temp value", func(){
		It("should change input value", func(){
			res:= myreceiver.Chain{4,6}.Add(2)
			Expect(res).To(Equal(myreceiver.Chain{6,8}))
		})
	})
})