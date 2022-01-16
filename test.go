package main

//import "fmt"

func canIDrrink(age int) bool {
	
	/*
	if koreanAge := age + 2; koreanAge < 18 {
		return false
	} else {
		return true
	}
	*/
	switch koreanAge := age + 2; koreanAge{
	case 10:
		return false
	case 18:
		return true	
	}
	return false
}

/*
func main() {
	names := []string{"nico","lynn","dal"}
	names = append(names,"flynn")
	//append는 새로운 slice를 준다. slice는 array인데 개수가 정해지지 않은 
	fmt.Println(names)
	/*
	a := 2
	b := &a
	//a = 5
	*b = 20
	fmt.Println(a)
	*/	
//}



