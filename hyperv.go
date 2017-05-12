package virt

import (
	"os/exec"
)

// HyperVVM record
type HyperVVM struct {
	VMName string
}

// NewHyperVVM from machine name
func NewHyperVVM(n string) (v *HyperVVM) {
	return &HyperVVM{
		VMName: n,
	}
}

// Start VM
func (v *HyperVVM) Start() (err error) {

	return exec.Command(
		"powershell",
		"-exec",
		"bypass",
		"-Command",
		"Start-VM",
		"-Name",
		v.VMName,
	).Run()
}

// Stop VM
func (v *HyperVVM) Stop() (err error) {

	return exec.Command(
		"powershell",
		"-exec",
		"bypass",
		"-Command",
		"Stop-VM",
		"-Name",
		v.VMName,
		"-Force",
	).Run()

}

// Restore VM to Snapshot
func (v *HyperVVM) Restore(Snapshot string) (err error) {

	return exec.Command(
		"powershell",
		"-exec",
		"bypass",
		"-Command",
		"Restore-VMSnapshot",
		"-VMName",
		v.VMName,
		"-Name",
		Snapshot,
		"-Confirm:$false",
	).Run()

}
