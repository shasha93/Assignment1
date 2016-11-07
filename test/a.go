package Assignment1

import(
   
   "fmt"
    
    )
func init(){
 fmt.Println("init called first")
}

func NumberFrequency(numbers []int) []int{
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
   //fmt.Printf("%d ",freq[i])
 }
 }
 return result
}



func FrequencyNames(strs []string) map[string]int{

freq:=make(map[string]int)
var i=0
  for i<len(strs){
   freq[strs[i]]++

i++
}

  return freq
}


