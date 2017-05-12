package virt

// VMHostType identify VMHost
type VMHostType int

const (
	// HostTypeNone (no control available)
	HostTypeNone VMHostType = iota - 1

	// HostTypeHyperV of vm
	HostTypeHyperV

	// HostTypeVirtualBox of vm
	HostTypeVirtualBox
)

// Driver for VM
type Driver interface {
	Stop() error
	Start() error
	Restore(Snapshot string) error
}

// VM controller
type VM struct {
	Controller Driver
}

// Ready VM to Snapshot
func (v *VM) Ready(Snapshot string) (err error) {

	err = v.Controller.Stop()
	if err != nil {
		return
	}

	err = v.Controller.Restore(Snapshot)
	if err != nil {
		return
	}

	err = v.Controller.Start()

	return
}

// NewVM from type and name
func NewVM(t VMHostType, n string) (v *VM) {

	var drv Driver

	switch t {
	case HostTypeVirtualBox:
		drv = NewVirtualBoxVM(n)
	case HostTypeHyperV:
		drv = NewHyperVVM(n)
	default:
		drv = NewNopObject()
	}

	return &VM{
		Controller: drv,
	}
}
