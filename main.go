// Copyright 2022 Jonas Kwiedor. All rights reserved.
// Use of this source code is governed by the MIT
// license that can be found in the LICENSE file.

package main

import (
	"fmt"
	"os"
)

func main() {

	path := os.Getenv("GITHUB_OUTPUT")
	file, err := os.OpenFile(path, os.O_APPEND|os.O_WRONLY, os.ModeAppend)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	_, err = file.WriteString("\"core_result=2\"")
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	err = file.Close()
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

}
