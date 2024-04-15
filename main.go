package main

import (
	"fmt"
	"strconv"
)
var deckSize int = 20

func swap (a,b *int){
   *a,*b = *b,*a
}

func sum(list ...int)(sum string){
   for _,val := range list {
      sum += strconv.Itoa(val)
  }
  return
}
func swapByValue(a,b int) {
   a,b = b,a
}
func main() {
   card := newCard()
   fmt.Println(card)
   var x,y,z int;
   fmt.Printf("x %d y %d z %d \n",x,y,z)
   var ptr *int
   ptr = &deckSize
   *ptr = 12

   anInt,err := strconv.Atoi("123*")
   if err != nil {
      fmt.Println(anInt);
      fmt.Println(err)
   }
   for i:=0;i<10;i++ {
      fmt.Println(i)
   }
   str := "go routines"
   for ind,val := range str {
      fmt.Printf("value at %d is %c \n",ind,val)
   }
   a:=10
   b:=20 
   swap(&a,&b)
   swap(&a,&b)

   fmt.Printf("a %d b %d \n",a,b)
   fmt.Println(ptr)
   fmt.Printf("value at %p is %d \n",ptr,*ptr)
   fmt.Printf("IntSize %d\n",strconv.IntSize)

   fmt.Printf("%s\n",sum(1,2,2,5))
}

func newCard () string{
	return "Five of Diamonds"
}

// v declared and not used.
// no new variables on the left side of :=".
// myPhone := "Iqoo" should used inside funtion only not in package level