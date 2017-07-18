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
