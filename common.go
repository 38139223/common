package common
import "fmt"
func NewLoginer() loginer{
 return defaultLogin(0)
}
type Loginer interface{
login()
}
type defaultLogin int
func (d defaultLogin) Login(){
fmt.Println("login in ...")
}

package main
import "common"
func main(){
l:=common.NewLoginer()
l.login()
}
