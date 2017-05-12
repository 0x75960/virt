package virt

import (
	"os/exec"
	"time"
)

// VirtualBoxVM record
type VirtualBoxVM struct {
	VMName string
}

// NewVirtualBoxVM from machine name
func NewVirtualBoxVM(n string) (v *VirtualBoxVM) {
	return &VirtualBoxVM{
		VMName: n,
	}
}

// Start VM
func (v *VirtualBoxVM) Start() (err error) {

	defer time.Sleep(3 * time.Second)

	return exec.Command(
		"vboxmanage",
		"startvm",
		v.VMName,
	).Run()
}

// Stop VM
func (v *VirtualBoxVM) Stop() (err error) {

	defer time.Sleep(3 * time.Second)

	return exec.Command(
		"vboxmanage",
		"controlvm",
		v.VMName,
		"poweroff",
	).Run()
}

// Restore VM to Snapshot
func (v *VirtualBoxVM) Restore(Snapshot string) (err error) {

	defer time.Sleep(3 * time.Second)

	return exec.Command(
		"vboxmanage",
		"snapshot",
		v.VMName,
		"restore",
		Snapshot,
	).Run()
}
