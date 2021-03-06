// -*- Mode: Go; indent-tabs-mode: t -*-

/*
 * Copyright (C) 2017 Canonical Ltd
 *
 * This program is free software: you can redistribute it and/or modify
 * it under the terms of the GNU General Public License version 3 as
 * published by the Free Software Foundation.
 *
 * This program is distributed in the hope that it will be useful,
 * but WITHOUT ANY WARRANTY; without even the implied warranty of
 * MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
 * GNU General Public License for more details.
 *
 * You should have received a copy of the GNU General Public License
 * along with this program.  If not, see <http://www.gnu.org/licenses/>.
 *
 */

package builtin

import (
	"strings"

	"github.com/snapcore/snapd/interfaces"
	"github.com/snapcore/snapd/interfaces/apparmor"
	"github.com/snapcore/snapd/interfaces/kmod"
	"github.com/snapcore/snapd/interfaces/udev"
)

const broadcomAsicControlSummary = `allows using the broadcom-asic kernel module`

const broadcomAsicControlBaseDeclarationSlots = `
  broadcom-asic-control:
    allow-installation:
      slot-snap-type:
        - core
    deny-auto-connection: true
`
const broadcomAsicControlConnectedPlugAppArmor = `
# Description: Allow access to broadcom asic kernel module.

/sys/module/linux_bcm_knet/{,**} r,
/sys/module/linux_kernel_bde/{,**} r,
/sys/module/linux_user_bde/{,**} r,
/dev/linux-user-bde rw,
/dev/linux-kernel-bde rw,
/dev/linux-bcm-knet rw,
`

const broadcomAsicControlConnectedPlugUDev = `
KERNEL=="linux-user-bde", TAG+="###SLOT_SECURITY_TAGS###"
KERNEL=="linux-kernel-bde", TAG+="###SLOT_SECURITY_TAGS###"
KERNEL=="linux-bcm-knet", TAG+="###SLOT_SECURITY_TAGS###"
`

// The upstream linux kernel doesn't come with support for the
// necessary kernel modules we need to drive a Broadcom ASIC.
// All necessary modules need to be loaded on demand if the
// kernel the device runs with provides them.
var broadcomAsicControlConnectedPlugKMod = []string{
	"linux-user-bde",
	"linux-kernel-bde",
	"linux-bcm-knet",
}

// The type for broadcom-asic-control interface
type broadcomAsicControlInterface struct{}

// Getter for the name of the broadcom-asic-control interface
func (iface *broadcomAsicControlInterface) Name() string {
	return "broadcom-asic-control"
}

func (iface *broadcomAsicControlInterface) StaticInfo() interfaces.StaticInfo {
	return interfaces.StaticInfo{
		Summary:              broadcomAsicControlSummary,
		ImplicitOnCore:       true,
		ImplicitOnClassic:    true,
		BaseDeclarationSlots: broadcomAsicControlBaseDeclarationSlots,
	}
}

// Check validity of the defined slot
func (iface *broadcomAsicControlInterface) SanitizeSlot(slot *interfaces.Slot) error {
	return sanitizeSlotReservedForOS(iface, slot)
}

func (iface *broadcomAsicControlInterface) AppArmorConnectedPlug(spec *apparmor.Specification, plug *interfaces.Plug, plugAttrs map[string]interface{}, slot *interfaces.Slot, slotAttrs map[string]interface{}) error {
	spec.AddSnippet(broadcomAsicControlConnectedPlugAppArmor)
	return nil
}

func (iface *broadcomAsicControlInterface) UDevConnectedPlug(spec *udev.Specification, plug *interfaces.Plug, plugAttrs map[string]interface{}, slot *interfaces.Slot, slotAttrs map[string]interface{}) error {
	old := "###SLOT_SECURITY_TAGS###"
	for appName := range plug.Apps {
		tag := udevSnapSecurityName(plug.Snap.Name(), appName)
		snippet := strings.Replace(broadcomAsicControlConnectedPlugUDev, old, tag, -1)
		spec.AddSnippet(snippet)
	}
	return nil
}

func (iface *broadcomAsicControlInterface) KModConnectedPlug(spec *kmod.Specification, plug *interfaces.Plug, plugAttrs map[string]interface{}, slot *interfaces.Slot, slotAttrs map[string]interface{}) error {
	for _, kmod := range broadcomAsicControlConnectedPlugKMod {
		if err := spec.AddModule(kmod); err != nil {
			return err
		}
	}
	return nil
}

func (iface *broadcomAsicControlInterface) AutoConnect(*interfaces.Plug, *interfaces.Slot) bool {
	// Allow what is allowed in the declarations; see https://github.com/snapcore/snapd/blob/master/interfaces/policy/basedeclaration.go
	// for more details.
	return true
}

func init() {
	registerIface(&broadcomAsicControlInterface{})
}
