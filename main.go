package testSever

import "fmt"

func main()  {
	//$ GOSSAFUNC=main go build main.go 生成ssa
	defer func(){
		if err:=recover();err!=nil{
			fmt.Println("recover failed")
		}
	}()
	panic("panic recover test")
}
