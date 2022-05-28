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

// https://github.com/xiexianbin/gsync/blob/7431fdacf3/utils/concurrentmap_test.go

package concurrentmap

import (
	"math/rand"
	"strconv"
	"sync"
	"testing"
)

// 100000 set values, and 100000 read test
var times int = 100000

// BenchmarkTestConcurrentMap
func BenchmarkTestConcurrentMap(b *testing.B) {
	for k := 0; k < b.N; k++ {
		b.StopTimer()
		// generate 100000 keys map
		testKV := map[string]int{}
		for i := 0; i < 10000; i++ {
			testKV[strconv.Itoa(i)] = i
		}

		// new ConcurrentMap
		cMap := New()

		// set key and value to cMap
		for k, v := range testKV {
			cMap.Set(k, v)
		}

		// start timer
		b.StartTimer()

		wg := sync.WaitGroup{}
		wg.Add(2)

		// set values
		go func() {
			// rand key set value 100000 counts
			for i := 0; i < times; i++ {
				index := rand.Intn(times)
				cMap.Set(strconv.Itoa(index), index+1)
			}
			wg.Done()
		}()

		// read values
		go func() {
			// rand key, read 100000 counts
			for i := 0; i < times; i++ {
				index := rand.Intn(times)
				cMap.Get(strconv.Itoa(index))
			}
			wg.Done()
		}()

		// wait goroutine done
		wg.Wait()
	}
}
