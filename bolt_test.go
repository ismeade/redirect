//
// bolt_test.go
// Copyright (C) 2020 liyang <yang.li@51vcheck.cn>
//
// Distributed under terms of the MIT license.
//

package main

import (
	"fmt"
	"testing"
)

func Test(t *testing.T) {

	defer db.Close()

	insert("hello", "Hello World,This is Bolt Database11..")
	insert("hello2", "Hello World,This is Bolt Database22..")
	fmt.Println("===test===")
	fmt.Println(read("hello"))
	fmt.Println(read("hello2"))
	//rm("hello")
	//insert("news1", "this is a title.")

	//fetchAll(bucket)
}



