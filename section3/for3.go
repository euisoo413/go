package main

import "fmt"

func main(){	
	
Loop1:
	for i:=0;i<5;i++{
		for j:=0;j<5;j++{
			if i==2 && j==4{
				break Loop1 //break를 하면서 Loop1자리로 돌아간다.
			}
			fmt.Println("i:",i," j:",j)
		}
	}

//루프를 없앨 경우
//Loop1:
	for i:=0;i<5;i++{
		for j:=0;j<5;j++{
			if i==2 && j==4{
				break //Loop1을 없앨 경우 해당 자리의 것만 제외 시키고 *즉 가장 가까운 반목문으로만 나간다.
			}
			fmt.Println("i:",i," j:",j)
		}
	}
//Loop1:
for i:=0;i<5;i++{
	for j:=0;j<5;j++{
		if i==2 && j==4{
			break //Loop1을 없앨 경우 해당 자리의 것만 제외 시키고 *즉 가장 가까운 반목문으로만 나간다.
		}
		fmt.Println("i:",i," j:",j)
	}
}

for i:=0;i<10;i++{
	if i%2==0{
		continue //여기서의 컨틴뉴는 for문으로 가는걸 말하므로 아래의 출력은 제외된다
	}
	fmt.Println(i)
}

Loop2:
	for i:=0;i<5;i++{
		for j:=0;j<5;j++{
			if i==2 && j==4{
				continue Loop2 //break를 하면서 Loop2자리로 돌아간다. // 출력에서 제외된다.
			}
			fmt.Println("i:",i," j:",j)
		}
	}

	

}