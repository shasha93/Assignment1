package Assignment1

import (
	"testing"
         "fmt"
)

func TestNumbers(t *testing.T){
  
  values:=[]int{1,2,3,3}
  expectedOutput:=[]int{1,2,1}
 res:=NumberFrequency(values)
	fmt.Println(res)
  for i:=0;i<len(res);i++ {
  if res[i]!=expectedOutput[i]{
  t.Fatal("failed")
  }
 }
}



func TestStrings(t *testing.T){
  
  values:=[]string{"go","we","go","out"}
  commits := map[string]int{
    "go": 2,
    "we": 1,
    "out": 1,
}
 res:=FrequencyNames(values)
	fmt.Println("good job")
  for i,j:=range res {


  if j!=commits[i]{
  

  t.Fatal("failed")
  }
 }
}









