// Copyright 2022 Jonas Kwiedor. All rights reserved.
// Use of this source code is governed by the MIT
// license that can be found in the LICENSE file.

package main

import (
	"fmt"
	"os"
)

func main() {

	fmt.Println(os.Getenv("GITHUB_OUTPUT"))

}
