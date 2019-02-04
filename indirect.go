// Copyright 2019 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// +build indirect

package mods

import (
	_ "github.com/containerd/containerd/dialer"
	_ "github.com/docker/docker/api/types/image"
	_ "github.com/docker/libnetwork/ipvs"
	_ "github.com/opencontainers/go-digest"
	_ "github.com/opencontainers/runc/libcontainer"
)
