package main

import (
       "fmt"
       "os"

       )


func FrequencyNames(strs []string) map[string]int{

freq:=make(map[string]int)
var i=0
  for i<len(strs){
   freq[strs[i]]++

i++
}

  return freq
}


func ReadInput() []string{
   var nums int
   var str string
     
   fmt.Scanf("%d",&nums)

  strs:= make([]string,nums)

  for j:=0;j<nums;j++{
  fmt.Scanf("%s",&str)
   for ch:=range str{
     if str[ch]<57 && str[ch]>46 {
        fmt.Println("Numerals are not allowed !! please try again")
     os.Exit(1)
    }
      
  }
  strs[j]=str
 } 
  return strs
}

func Print(frequen map[string]int){
  for i,k:=range frequen{

	fmt.Println(i,k)
	}
}

func main() {
     
   strs:= ReadInput()
    frequen:= FrequencyNames(strs)

    Print(frequen)
   
}
