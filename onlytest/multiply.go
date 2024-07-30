package main

import "strconv"

func multiply(num1 string, num2 string) string {
    l1,l2:=len(num1),len(num2)
	es:=[]int{}
	nextadd:=0
	if l1<l2{
		for i := 0; i < l1; i++ {
			for j:=0;j<l2;j++{
				a1:=int(num1[i]-'0')
				a2:=int(num2[j]-'0')
				csum:=a1*a2+nextadd
				down:=csum/10
				nextadd=csum%10
				es=append(es,down)
			}

		}

	}
	if l2<l1
}
