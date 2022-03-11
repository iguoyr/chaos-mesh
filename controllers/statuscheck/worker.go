// Copyright Chaos Mesh Authors.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
// http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.
//

package statuscheck

type worker struct {
	stopCh chan struct{}

	spec struct{} // TODO StatusCheck spec
}

func newWorker() *worker {
	return &worker{
		stopCh: make(chan struct{}),
	}
}

// run periodically execute the status check.
func (w *worker) run() {

}

func (w *worker) stop() {
	close(w.stopCh)
}

// execute the status check once and records the result.
// Returns whether the worker should continue.
func (w *worker) execute() bool {
	return true
}