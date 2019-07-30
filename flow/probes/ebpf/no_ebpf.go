// +build !ebpf

/*
 * Copyright (C) 2017 Red Hat, Inc.
 *
 * Licensed under the Apache License, Version 2.0 (the "License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy ofthe License at
 *
 *     http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specificlanguage governing permissions and
 * limitations under the License.
 *
 */

package ebpf

import (
	"github.com/skydive-project/skydive/flow/probes"
	"github.com/skydive-project/skydive/probe"
)

// NewProbe returns a new eBPF probe
func NewProbe(ctx probes.Context, bundle *probe.Bundle) (probes.FlowProbeHandler, error) {
	return nil, probe.ErrNotCompiled
}
