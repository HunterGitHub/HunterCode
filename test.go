package main

import (
	"fmt"
	"reflect"
	"regexp"
	"os"
	"bytes"
)

type A struct {
	b int
}

func main() {
	matched, err := regexp.MatchString("foo.*", "seafood")
	fmt.Println(matched, err)
	matched, err = regexp.MatchString("bar.*", "seafood")
	fmt.Println(matched, err)
	matched, err = regexp.MatchString(`a(b`, "seafood")
	fmt.Println(matched, err)
	matched, err = regexp.MatchString(`/(css|js)/[a-z]+\.css`, "/css/fdfd.css")
	matched, err = regexp.MatchString(`/(css|js)/[a-z]+\.(css|js)`, "/js/fdfd.js")
	fmt.Println(matched, err)
	fmt.Println(reflect.TypeOf(matched))
	var a []A = []A{{1},{2},{3}} 
	for _, aa := range a {
		print(aa.b)
	}


	w := &A{}
	fmt.Println(w)

	file := "./db"
	f, err := os.OpenFile(file, os.O_RDWR|os.O_CREATE, 0); if err != nil { fmt.Println(err) }
	defer f.Close()
	b := make([]byte, 2, 4)
	f.Read(b)
	fmt.Println(string(b))

	sss := S{1}
	fmt.Println("sss", sss.hh())

	b2 := make([]byte, 30)
	b2[4] = 3
	b2[5] = 4
	fmt.Println(b2[4:6])
	if bytes.Equal(b2[4:6], []byte{3,4}) {
		fmt.Println("ok")
	}

	if b2[4] == 3 {
		fmt.Println("okkk")
	}
}

type S struct {
	k int
}

func (s *S) hh() int {
	return s.k
}
