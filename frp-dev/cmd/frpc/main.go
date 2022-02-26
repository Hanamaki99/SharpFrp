// Copyright 2016 fatedier, fatedier@gmail.com
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package main

import "C"
import (
	"fmt"
	_ "github.com/fatedier/frp/assets/frpc"
	"github.com/fatedier/frp/cmd/frpc/sub"
	"github.com/fatedier/golib/crypto"
	"math/rand"
	"os"
	"time"
)

func main() {
	//mainDelegate()
	//sub.Execute()
}

//start_init :
//export start_init
func start_init(charargs *C.char) {
	var stringargs string
	stringargs = C.GoString(charargs)
	crypto.DefaultSalt = "frp"
	rand.Seed(time.Now().UnixNano())
	err := sub.RunClient(stringargs)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
