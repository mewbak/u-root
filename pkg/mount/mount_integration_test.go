// Copyright 2019 the u-root Authors. All rights reserved
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// +build !race

package mount

import (
	"testing"

	"github.com/u-root/u-root/pkg/qemu"
	"github.com/u-root/u-root/pkg/uroot"
	"github.com/u-root/u-root/pkg/vmtest"
)

func TestIntegration(t *testing.T) {
	o := &vmtest.Options{
		BuildOpts: uroot.Opts{
			Commands: uroot.BusyBoxCmds(
				"github.com/u-root/u-root/cmds/core/ls",
			),
		},
		QEMUOpts: qemu.Options{
			Devices: []qemu.Device{
				// CONFIG_ATA_PIIX is required for this option to work.
				qemu.ArbitraryArgs{"-hda", "testdata/1MB.ext4_vfat"},
				qemu.ArbitraryArgs{"-hdb", "testdata/12Kzeros"},
			},
		},
	}
	vmtest.GolangTest(t, []string{"github.com/u-root/u-root/pkg/mount"}, o)
}
