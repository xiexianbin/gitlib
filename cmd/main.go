// Copyright 2022 xiexianbin<me@xiexianbin.cn>
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

import "github.com/xiexianbin/golib/logger"

func main() {
	logger.Print("print")
	logger.Printf("print %s", "format")
	logger.SetLogLevel(logger.DEBUG)
	logger.Debug("Debug")
	logger.Debugf("Debugf %s", "format")
	logger.Info("Info")
	logger.Infof("Infof %s", "format")
	logger.Warn("Warn")
	logger.Warnf("Warnf %s", "format")
	logger.Error("Error")
	logger.Errorf("Errorf %s", "format")
	logger.Critical("Critical")
	logger.Criticalf("Criticalf %s", "format")
	//logger.Fatal("Fatal")
	logger.Fatalf("Fatalf %s", "format")
}
