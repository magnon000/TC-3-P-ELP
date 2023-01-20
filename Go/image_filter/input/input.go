package input

import (
	"fmt"
	"image/png"
	"os"
)

var(
	filename="test.png"
)
//structure?
type Image struct {
	name string
	
}
func load(filename string) {
	f, err := os.Open(filename)
    if err != nil {
		panic(err)
	}
	m, err := png.Decode(f)
    if err != nil {
        panic(err)
    }
	fmt.Printf("%v\n", m.Bounds())     
    fmt.Printf("%v\n", m.ColorModel())
    fmt.Printf("%v\n", m.At(100,100))
}
