package virt

// NopObject record
type NopObject struct{}

// NewNopObject from machine name
func NewNopObject() (v *NopObject) {
	return &NopObject{}
}

// Start VM
func (v *NopObject) Start() (err error) {
	return
}

// Stop VM
func (v *NopObject) Stop() (err error) {
	return
}

// Restore VM to Snapshot
func (v *NopObject) Restore(Snapshot string) (err error) {
	return
}
