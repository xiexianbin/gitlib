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

package logger

import "testing"

func TestDebug(t *testing.T) {
	SetLogLevel(DEBUG)
	Debug("Debug")
	Debugf("Debugf %s", "format")
	Info("Info")
	Infof("Infof %s", "format")
	Warn("Warn")
	Warnf("Warnf %s", "format")
	Error("Error")
	Errorf("Errorf %s", "format")
	Critical("Critical")
	Criticalf("Criticalf %s", "format")
	//Fatal("Fatal")
	//Fatalf("Fatalf %s", "format")
	Print("print")
	Printf("print %s", "format")
}

func TestInfo(t *testing.T) {
	SetLogLevel(INFO)
	Debug("Debug")
	Debugf("Debugf %s", "format")
	Info("Info")
	Infof("Infof %s", "format")
	Warn("Warn")
	Warnf("Warnf %s", "format")
	Error("Error")
	Errorf("Errorf %s", "format")
	Critical("Critical")
	Criticalf("Criticalf %s", "format")
	//Fatal("Fatal")
	//Fatalf("Fatalf %s", "format")
	Print("print")
	Printf("print %s", "format")
}
