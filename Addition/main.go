package Addition

func Add(a ...int,name string) (int,string){
	sum := 0
	for _,v := range a {
		sum+=v
	}
	return sum,name
}
