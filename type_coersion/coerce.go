package main

import (
	"encoding/json"
	"fmt"
)

type S struct {
	sixth string
}

type Q struct {
	sixth string
}

type X struct {
	Sixth string
}

type Y struct {
	Sixth string
}

type A struct {
	first  int
	second string
	third  float64
	fourth S
}

type B struct {
	first  int
	second string
	third  float64
	fourth Q
}

type C struct {
	fifth  int
	second string
	third  float64
	fourth S
}

type D struct {
	First  int `json:"First"`
	Second string
	Third  float64
	Fourth X
}

type E struct {
	First  int
	Second string
	Third  float64
	Fourth Y
}

func main() {
	a := A{first: 5, second: "hi", third: 5.4, fourth: S{sixth: "hey"}}

	fmt.Println(a)
	b := B{}
	/*
		b = B(a)
		fmt.Println(b)
	*/
	export, err := json.Marshal(&a)
	fmt.Println(string(export))
	if err != nil {
		fmt.Println(err)
		return
	}
	//	c := C{}
	if err = json.Unmarshal(export, &b); err != nil {
		fmt.Println(err)
	}
	fmt.Println(b)

	d := D{First: 5, Second: "hi", Third: 5.4, Fourth: X{Sixth: "hey"}}
	fmt.Println(d)
	e := E{}
	export, err = json.Marshal(&d)
	fmt.Println(string(export))
	if err != nil {
		fmt.Println(err)
		return
	}
	if err = json.Unmarshal(export, &e); err != nil {
		fmt.Println(err)
	}
	fmt.Println(e)

	fmt.Println(d)
	export, err = json.Marshal(&d)
	fmt.Println(string(export))
	if err != nil {
		fmt.Println(err)
		return
	}
	if err = json.Unmarshal(export, &b); err != nil {
		fmt.Println(err)
	}
	fmt.Println(b)
}

