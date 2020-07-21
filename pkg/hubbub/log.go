// Copyright 2020 Google LLC
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//      http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package hubbub

import (
	"fmt"
	"github.com/google/triage-party/pkg/models"

	"k8s.io/klog/v2"
)

func (h *Engine) logRate(r models.Rate) {
	msg := fmt.Sprintf("GitHub API hourly quota remaining: %d of %d, resets at %s", r.Remaining, r.Limit, r.Reset)

	if r.Remaining < 25 {
		klog.Error(msg)
		return
	}

	if r.Remaining < 250 {
		klog.Warning(msg)
		return
	}

	if r.Remaining%100 == 1 {
		klog.Info(msg)
	}
}
