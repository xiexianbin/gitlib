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

// https://github.com/xiexianbin/goto/blob/master/utils/sshpass.go

package bash

//import (
//	"fmt"
//	"time"
//
//	expect "github.com/google/goexpect"
//)
//
//func SshPassword(ip, username, password string, port, timeout time.Duration) error {
//	//sshClt, err := ssh.Dial("tcp", *addr, &ssh.ClientConfig{
//	//	User:            *user,
//	//	Auth:            []ssh.AuthMethod{ssh.Password(*pass1)},
//	//	HostKeyCallback: ssh.InsecureIgnoreHostKey(),
//	//})
//	//if err != nil {
//	//	log.Fatalf("ssh.Dial(%q) failed: %v", *addr, err)
//	//}
//	//defer sshClt.Close()
//
//	//e, _, err := expect.SpawnSSH(sshClt, timeout)
//	//if err != nil {
//	//	log.Fatal(err)
//	//}
//	//defer e.Close()
//	//
//	//e.ExpectBatch([]expect.Batcher{
//	//	&expect.BCas{[]expect.Caser{
//	//		&expect.Case{R: regexp.MustCompile(`router#`), T: expect.OK()},
//	//		&expect.Case{R: regexp.MustCompile(`Login: `), S: *user,
//	//			T: expect.Continue(expect.NewStatus(codes.PermissionDenied, "wrong username")), Rt: 3},
//	//		&expect.Case{R: regexp.MustCompile(`Password: `), S: *pass1, T: expect.Next(), Rt: 1},
//	//		&expect.Case{R: regexp.MustCompile(`Password: `), S: *pass2,
//	//			T: expect.Continue(expect.NewStatus(codes.PermissionDenied, "wrong password")), Rt: 1},
//	//	}},
//	//}, timeout)
//	e, _, err := expect.Spawn(fmt.Sprintf("ssh %s", ip), -1)
//	if err != nil {
//		//log.Fatal(err)
//		fmt.Println("expect.Spawn:", err)
//		return err
//	}
//	defer e.Close()
//
//	e.ExpectBatch([]expect.Batcher{
//		&expect.BExp{R: "username:"},
//		&expect.BSnd{S: username + "\n"},
//		&expect.BExp{R: "password:"},
//		&expect.BSnd{S: password + "\n"},
//	}, timeout)
//
//	return nil
//}
