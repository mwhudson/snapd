// -*- Mode: Go; indent-tabs-mode: t -*-

/*
 * Copyright (C) 2016 Canonical Ltd
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
	"github.com/snapcore/snapd/interfaces"
	"github.com/snapcore/snapd/interfaces/apparmor"
	"github.com/snapcore/snapd/interfaces/seccomp"
)

const lxdSummary = `allows access to the LXD socket`

const lxdBaseDeclarationSlots = `
  lxd:
    allow-installation: false
    deny-connection: true
    deny-auto-connection: true
`

const lxdConnectedPlugAppArmor = `
# Description: allow access to the LXD daemon socket. This gives privileged
# access to the system via LXD's socket API.

/var/snap/lxd/common/lxd/unix.socket rw,
`

const lxdConnectedPlugSecComp = `
# Description: allow access to the LXD daemon socket. This gives privileged
# access to the system via LXD's socket API.

shutdown
socket AF_NETLINK - NETLINK_GENERIC
`

type lxdInterface struct{}

func (iface *lxdInterface) Name() string {
	return "lxd"
}

func (iface *lxdInterface) StaticInfo() interfaces.StaticInfo {
	return interfaces.StaticInfo{
		Summary:              lxdSummary,
		BaseDeclarationSlots: lxdBaseDeclarationSlots,
	}
}

func (iface *lxdInterface) AppArmorConnectedPlug(spec *apparmor.Specification, plug *interfaces.Plug, plugAttrs map[string]interface{}, slot *interfaces.Slot, slotAttrs map[string]interface{}) error {
	spec.AddSnippet(lxdConnectedPlugAppArmor)
	return nil
}

func (iface *lxdInterface) SecCompConnectedPlug(spec *seccomp.Specification, plug *interfaces.Plug, plugAttrs map[string]interface{}, slot *interfaces.Slot, slotAttrs map[string]interface{}) error {
	spec.AddSnippet(lxdConnectedPlugSecComp)
	return nil
}

func (iface *lxdInterface) AutoConnect(*interfaces.Plug, *interfaces.Slot) bool {
	// allow what declarations allowed
	return true
}

func init() {
	registerIface(&lxdInterface{})
}
