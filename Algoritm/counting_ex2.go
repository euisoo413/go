package main

import "fmt"

func main(){

	// arr := string[]{a,b,c,g,e,f,a,c,e,b,e,a,z,g,e,f,q,e,w,z,x,v,b,n,n,d,f,d,s,e,w,f,r,h,u,l,i,u,j,m,h,o,p,p,p}
	str:="azzzzzzzzzzzzzzzzzzsdfabdaqwerqwenrkjnjqwuanxszxcvnkqwernkjasdfyh"
	var count [26]int

	for i:=0;i<len(str);i++{
		count[str[i]-'a']++
	}

	maxCount := 0
	var maxCh byte
	for i:=0;i<len(count);i++{
		if count[i]>maxCount{
			maxCount = count[i]
			maxCh = byte(i+'a')
		}
	}

	fmt.Printf("%c, %d", maxCh, maxCount)


}