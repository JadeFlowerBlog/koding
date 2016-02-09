package softlayer

import (
	"koding/db/mongodb/modelhelper"
	"koding/kites/kloud/machinestate"

	"golang.org/x/net/context"
)

// Stop stops the given machine
func (m *Machine) Stop(ctx context.Context) error {
	if err := modelhelper.ChangeMachineState(m.ObjectId, "Machine is stopping", machinestate.Stopping); err != nil {
		return err
	}

	//Get the SoftLayer virtual guest service
	svc, err := m.Session.SLClient.GetSoftLayer_Virtual_Guest_Service()
	if err != nil {
		return err
	}

	meta, err := m.GetMeta()
	if err != nil {
		return err
	}

	_, err = svc.PowerOff(meta.Id)
	if err != nil {
		return err
	}

	if err := waitState(svc, meta.Id, "HALTED", m.StateTimeout); err != nil {
		return err
	}

	if err := m.deleteDomains(); err != nil {
		m.Log.Warning("couldn't delete domains while stopping machine: %s", err)
	}

	return m.MarkAsStoppedWithReason("Machine is stopped")
}
