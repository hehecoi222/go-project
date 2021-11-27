package main

func powWithMod(a,b int) int {
	result:=1;
	for i:=0;i<b;i++ {
		result=(result*a)%1337
	}
	return result
}

func superPow(a int, b []int) int {
	if len(b) <2 {
		return a
	}
	a=a%1337
	base :=a
	result := powWithMod(base,b[len(b)-1])
	for i:=len(b)-2;i>=0;i-- {
		base= powWithMod(base,10)
		if base == 0 {
			return 0
		}
		result = (result*powWithMod(base,b[i]))%1337
	}
	return result
}