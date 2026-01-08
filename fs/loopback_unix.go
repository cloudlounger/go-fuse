//go:build !freebsd

// Copyright 2024 the Go-FUSE Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.
package fs

import (
	"context"
	"fmt"
	"syscall"

	"golang.org/x/sys/unix"
)

var _ = (NodeListxattrer)((*LoopbackNode)(nil))

func (n *LoopbackNode) Listxattr(ctx context.Context, dest []byte) (uint32, syscall.Errno) {
	sz, err := unix.Llistxattr(n.path(), dest)
	if err != nil {
		fmt.Printf("LoopbackNode.Listxattr, path %s err %v\n", n.path(), err)
	}
	return uint32(sz), ToErrno(err)
}
