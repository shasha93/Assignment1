package main

import "fmt"

func NumberFrequency(numbers []int){
freq:=make([]int ,10)
result:=make([]int ,0,10)
 i:=0
  for i<len(numbers){
   freq[numbers[i]]++

i++
}

for i=0;i<len(freq);i++{
  if freq[i]>0{
   result=append(result,freq[i])
   fmt.Printf("%d ",freq[i])
 }
 }
 return result
}

func main() {
var num,nums int
  
  fmt.Scanf("%d",&nums)
numbers:= make([]int,0,nums)
  for j:=0;j<nums;j++{
  fmt.Scanf("%d",&num)
numbers=append(numbers,num)
} 

NumberFrequency(numbers)

	
}
