package main

import (
	"fmt"
	"io"
	"log"
	"math"
	"os"
	"runtime"
	"runtime/pprof"
	"strings"
)


var pow = []int{1, 2, 4, 8, 16, 32, 64, 128}

//func main() {
//	var arr [8]int
//	for j := 0;j < 8;j++ {
//		a := pow[j]
//		p := &a
//		arr[j] = pow[j]
//		fmt.Println(p)
//	}
//	for i, v := range pow {
//		var b = &v
//		//arr[i] = v
//		fmt.Printf("2**%d = %d = %p\n", i, v,b)
//	}
//
//	for j := 0;j < 8;j++ {
//		a := arr[j]
//		p := &a
//		fmt.Println(p)
//	}
//
//	var x = "aaa"
//	var y = x
//	var xp = &x
//	var yp = &y
//	var r = (x == y)
//	fmt.Println(xp)
//	fmt.Println(yp)
//	fmt.Println(r)
//}

type rot13Reader struct {
	r io.Reader
}
func (rot rot13Reader) Read(p []byte) (n int, err error){
	n, err = rot.r.Read(p)
	for i := 0; i < len(p); i++ {
		if p[i] >= 'A' && p[i] < 'Z' {
			p[i] = 65 + (((p[i] - 65) + 13) % 26)
		} else if p[i] >= 'a' && p[i] <= 'z' {
			p[i] = 97 + (((p[i] - 97) + 13) % 26)
		}
	}
	return
}
func main() {
	runtime.SetCPUProfileRate(500)
	cpuf, err := os.Create("cpu_profile")
	if err != nil {
		log.Fatal(err)
	}
	pprof.StartCPUProfile(cpuf)
	defer pprof.StopCPUProfile()
	s := strings.NewReader("Lbh penpxrq gur pbqr!")
	r := rot13Reader{s}
	io.Copy(os.Stdout, &r)

}

func printSlice(s []int) {
	fmt.Printf("len=%d cap=%d %v\n", len(s), cap(s), s)
}

func fibonacci() func() int {
	a := 0
	b := 1
	return func() int{
		temp := b
		b = a + b
		a = temp
		return b
	}

}

type Abser interface {
	Abs() float64
}
//
//func main() {
//	r := strings.NewReader("Hello, Reader!")
//
//
//	b := make([]byte, 8)
//	for {
//		n, err := r.Read(b)
//		fmt.Printf("n = %v err = %v b = %v\n", n, err, b)
//		fmt.Printf("b[:n] = %q\n", b[:n])
//		if err == io.EOF {
//			break
//		}
//	}
//}

//func main() {
//	var a Abser
//	f := MyFloat(-math.Sqrt2)
//	v := Vertex{3, 4}
//
//	a = f  // a MyFloat implements Abser
//
//	a = &v // a *Vertex implements Abser
//
//	// In the following line, v is a Vertex (not *Vertex)
//	// and does NOT implement Abser.
//	a = v
//	fmt.Println(a.Abs())
//
//
//}

type MyFloat float64

func (f MyFloat) Abs() float64 {
	if f < 0 {
		return float64(-f)
	}
	return float64(f)
}

type Vertex struct {
	X, Y float64
}

func (v Vertex) Abs() float64 {
	return math.Sqrt(v.X*v.X + v.Y*v.Y)
}


