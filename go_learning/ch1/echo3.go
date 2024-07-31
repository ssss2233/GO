package ch1
import(
	"fmt"
	"os"
)

func Echo3(){
	for i,arg := range os.Args[0:]{
		fmt.Println(i,arg)
	}
}