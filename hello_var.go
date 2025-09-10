package main
import "fmt"
 const (
    Red = iota   // 0
    Green        // 1
    Blue         // 2
)
func main() {
		fmt.Println(Red, Green, Blue)
    var hello string
    hello = "Hello world"
    fmt.Println(hello)
	
    fmt.Println(hello)  // Hello world
     
    hello = "Hello Go"
    fmt.Println(hello)  // Hello Go
    name := "Tom"
    fmt.Println(name)
    hello = "Go Go Go Ole Ole Ole"
    fmt.Println(hello)  // Go Go Go Ole Ole Ole
		var a int8 = -1
    var b uint8 = 2
    var c byte = 3 // byte - синоним типа uint8
    var d int16 = -4
    var f uint16 = 5
    var g int32 = -6
    var h rune = -7 // rune - синоним типа int32
    var j uint32 = 8
    var k int64 = -9
    var l uint64 = 10
    var m int = 102
    var n uint = 105
 
    fmt.Println("a: ", a)
    fmt.Println("b: ", b)
    fmt.Println("c: ", c)
    fmt.Println("d: ", d)
    fmt.Println("f: ", f)
    fmt.Println("g: ", g)
    fmt.Println("h: ", h)
    fmt.Println("j: ", j)
    fmt.Println("k: ", k)
    fmt.Println("l: ", l)
    fmt.Println("m: ", m)
    fmt.Println("n: ", n)
		var isAlive bool = true
    var isEnabled bool = false
 
    fmt.Println("isAlive: ", isAlive)
    fmt.Println("isEnabled: ", isEnabled)
}
