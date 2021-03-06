// Copyright (c) 2018, Sylabs Inc. All rights reserved.
// This software is licensed under a 3-clause BSD license. Please consult the
// LICENSE file distributed with the sources of this project regarding your
// rights to use or distribute this software.

package singularity

import (
	"github.com/singularityware/singularity/src/runtime/engines/common/config"
)

// EngineOperations describes a runtime engine
type EngineOperations struct {
	CommonConfig *config.Common `json:"-"`
	EngineConfig *EngineConfig  `json:"engineConfig"`
}

// InitConfig stores the pointer to config.Common
func (e *EngineOperations) InitConfig(cfg *config.Common) {
	e.CommonConfig = cfg
}

// Config returns a pointer to a singularity.EngineConfig literal as a
// config.EngineConfig interface. This pointer gets stored in the Engine.Common
// field.
func (e *EngineOperations) Config() config.EngineConfig {
	return e.EngineConfig
}

// IsRunAsInstance returns true if the runtime engine was run as an instance
func (e *EngineOperations) IsRunAsInstance() bool {
	return e.EngineConfig.GetInstance()
}

// IsAllowSUID always returns true to allow SUID workflow
func (e *EngineOperations) IsAllowSUID() bool {
	return true
}
