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

	err := put("", "Hello World,This is Bolt Database11..")
	if err != nil {
		fmt.Println(err.Error())
	}
	put("hello", "Hello World,This is Bolt Database11..")
	put("hello2", "Hello World,This is Bolt Database22..")
	fmt.Println("===test===")
	fmt.Println(get("hello"))
	fmt.Println(get("hello2"))
	//rm("hello")
	//insert("news1", "this is a title.")

	//fetchAll(bucket)
}
