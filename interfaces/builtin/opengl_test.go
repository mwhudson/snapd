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

package builtin_test

import (
	. "gopkg.in/check.v1"

	"github.com/snapcore/snapd/interfaces"
	"github.com/snapcore/snapd/interfaces/apparmor"
	"github.com/snapcore/snapd/interfaces/builtin"
	"github.com/snapcore/snapd/snap"
	"github.com/snapcore/snapd/snap/snaptest"
	"github.com/snapcore/snapd/testutil"
)

type OpenglInterfaceSuite struct {
	iface interfaces.Interface
	slot  *interfaces.Slot
	plug  *interfaces.Plug
}

var _ = Suite(&OpenglInterfaceSuite{
	iface: builtin.MustInterface("opengl"),
})

func (s *OpenglInterfaceSuite) SetUpTest(c *C) {
	// Mock for OS Snap
	osSnapInfo := snaptest.MockInfo(c, `
name: ubuntu-core
type: os
slots:
  test-opengl:
    interface: opengl
`, nil)
	s.slot = &interfaces.Slot{SlotInfo: osSnapInfo.Slots["test-opengl"]}

	// Snap Consumers
	consumingSnapInfo := snaptest.MockInfo(c, `
name: client-snap
apps:
  app-accessing-opengl:
    command: foo
    plugs: [opengl]
`, nil)
	s.plug = &interfaces.Plug{PlugInfo: consumingSnapInfo.Plugs["opengl"]}
}

func (s *OpenglInterfaceSuite) TestName(c *C) {
	c.Assert(s.iface.Name(), Equals, "opengl")
}

func (s *OpenglInterfaceSuite) TestSanitizeSlot(c *C) {
	c.Assert(s.slot.Sanitize(s.iface), IsNil)
	slot := &interfaces.Slot{SlotInfo: &snap.SlotInfo{
		Snap:      &snap.Info{SuggestedName: "some-snap"},
		Name:      "opengl",
		Interface: "opengl",
	}}
	c.Assert(slot.Sanitize(s.iface), ErrorMatches,
		"opengl slots are reserved for the core snap")
}

func (s *OpenglInterfaceSuite) TestSanitizePlug(c *C) {
	c.Assert(s.plug.Sanitize(s.iface), IsNil)
}

func (s *OpenglInterfaceSuite) TestUsedSecuritySystems(c *C) {
	expectedSnippet := `
# Description: Can access opengl.

  # specific gl libs
  /var/lib/snapd/lib/gl/ r,
  /var/lib/snapd/lib/gl/** rm,

  /dev/dri/ r,
  /dev/dri/card0 rw,
  # nvidia
  @{PROC}/driver/nvidia/params r,
  @{PROC}/modules r,
  /dev/nvidiactl rw,
  /dev/nvidia-modeset rw,
  /dev/nvidia* rw,
  unix (send, receive) type=dgram peer=(addr="@nvidia[0-9a-f]*"),

  # eglfs
  /dev/vchiq rw,
  /sys/devices/pci[0-9]*/**/config r,
  /sys/devices/pci[0-9]*/**/{,subsystem_}device r,
  /sys/devices/pci[0-9]*/**/{,subsystem_}vendor r,

  # FIXME: this is an information leak and snapd should instead query udev for
  # the specific accesses associated with the above devices.
  /sys/bus/pci/devices/** r,
  /run/udev/data/+drm:card* r,
  /run/udev/data/+pci:[0-9]* r,

  # FIXME: for each device in /dev that this policy references, lookup the
  # device type, major and minor and create rules of this form:
  # /run/udev/data/<type><major>:<minor> r,
  # For now, allow 'c'haracter devices and 'b'lock devices based on
  # https://www.kernel.org/doc/Documentation/devices.txt
  /run/udev/data/c226:[0-9]* r,  # 226 drm
`

	// connected plugs have a non-nil security snippet for apparmor
	apparmorSpec := &apparmor.Specification{}
	err := apparmorSpec.AddConnectedPlug(s.iface, s.plug, nil, s.slot, nil)
	c.Assert(err, IsNil)
	c.Assert(apparmorSpec.SecurityTags(), DeepEquals, []string{"snap.client-snap.app-accessing-opengl"})
	aasnippet := apparmorSpec.SnippetForTag("snap.client-snap.app-accessing-opengl")
	c.Assert(aasnippet, Equals, expectedSnippet)
}

func (s *OpenglInterfaceSuite) TestInterfaces(c *C) {
	c.Check(builtin.Interfaces(), testutil.DeepContains, s.iface)
}
