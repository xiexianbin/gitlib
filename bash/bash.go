// Copyright 2020 xiexianbin<me@xiexianbin.cn>
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

// https://github.com/xiexianbin/goto/blob/master/utils/bash.go

package bash

import (
	"os/exec"
)

import (
	"fmt"
	"runtime"
	"strings"
)

// Run is a shared function used by check and reload
// to run the given command.
// It returns nil if the given cmd returns 0.
// The command can be run on unix and windows.
func Run(cmd string, maskStrs ...string) (string, error) {
	maskStr := ""
	if len(maskStrs) >= 1 {
		maskStr = maskStrs[0]
	}
	if maskStr != "" {
		fmt.Println("Running:", strings.Replace(cmd, maskStr, "******", -1))
	}
	var c *exec.Cmd
	if runtime.GOOS == "windows" {
		c = exec.Command("cmd", "/C", cmd)
	} else {
		c = exec.Command("/bin/sh", "-c", cmd)
	}

	output, err := c.CombinedOutput()
	outStr := fmt.Sprintf("%s", string(output))
	return outStr, err
}
