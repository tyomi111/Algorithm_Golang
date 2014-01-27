package main

import(
	"fmt"
)

func main(){
	s := []int{2,3,5,6,4,1}
	for i := 0;i < len(s)-1;i++{
		for j := len(s)-1;j > i;j--{
			if s[j] < s[j-1] {
				desc := s[j]
				s[j] = s[j-1]
				s[j-1] = desc
			}
		}
	}
	//ソート後の出力
	for i := 0;i < len(s);i++{
		fmt.Print(s[i]," ")
	}
	for k := 0;k < len(s)-1;k++{
		for l := 0;l < len(s)-k-1;l++{
			if s[l] < s[l+1]{
				asc := s[l]
				s[l] = s[l+1]
				s[l+1] = asc
			}
		}
	}
	fmt.Println("")
	//ソート後の処理
	for j := 0;j < len(s);j++{
		fmt.Print(s[j]," ")
	}
}