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

type JoystickInterfaceSuite struct {
	iface interfaces.Interface
	slot  *interfaces.Slot
	plug  *interfaces.Plug
}

var _ = Suite(&JoystickInterfaceSuite{
	iface: builtin.MustInterface("joystick"),
})

func (s *JoystickInterfaceSuite) SetUpTest(c *C) {
	// Mock for OS Snap
	osSnapInfo := snaptest.MockInfo(c, `
name: ubuntu-core
type: os
slots:
  test-joystick:
    interface: joystick
`, nil)
	s.slot = &interfaces.Slot{SlotInfo: osSnapInfo.Slots["test-joystick"]}

	// Snap Consumers
	consumingSnapInfo := snaptest.MockInfo(c, `
name: client-snap
apps:
  app-accessing-joystick:
    command: foo
    plugs: [joystick]
`, nil)
	s.plug = &interfaces.Plug{PlugInfo: consumingSnapInfo.Plugs["joystick"]}
}

func (s *JoystickInterfaceSuite) TestName(c *C) {
	c.Assert(s.iface.Name(), Equals, "joystick")
}

func (s *JoystickInterfaceSuite) TestSanitizeSlot(c *C) {
	c.Assert(s.slot.Sanitize(s.iface), IsNil)
	slot := interfaces.Slot{SlotInfo: &snap.SlotInfo{
		Snap:      &snap.Info{SuggestedName: "some-snap"},
		Name:      "joystick",
		Interface: "joystick",
	}}
	c.Assert(slot.Sanitize(s.iface), ErrorMatches,
		"joystick slots are reserved for the core snap")
}

func (s *JoystickInterfaceSuite) TestSanitizePlug(c *C) {
	c.Assert(s.plug.Sanitize(s.iface), IsNil)
}

func (s *JoystickInterfaceSuite) TestUsedSecuritySystems(c *C) {
	expectedSnippet := `
# Description: Allow reading and writing to joystick devices (/dev/input/js*).

# Per https://github.com/torvalds/linux/blob/master/Documentation/admin-guide/devices.txt
# only js0-js31 is valid so limit the /dev and udev entries to those devices.
/dev/input/js{[0-9],[12][0-9],3[01]} rw,
/run/udev/data/c13:{[0-9],[12][0-9],3[01]} r,
`

	// connected plugs have a non-nil security snippet for apparmor
	apparmorSpec := &apparmor.Specification{}
	err := apparmorSpec.AddConnectedPlug(s.iface, s.plug, nil, s.slot, nil)
	c.Assert(err, IsNil)
	c.Assert(apparmorSpec.SecurityTags(), DeepEquals, []string{"snap.client-snap.app-accessing-joystick"})
	aasnippet := apparmorSpec.SnippetForTag("snap.client-snap.app-accessing-joystick")
	c.Assert(aasnippet, Equals, expectedSnippet)
}

func (s *JoystickInterfaceSuite) TestInterfaces(c *C) {
	c.Check(builtin.Interfaces(), testutil.DeepContains, s.iface)
}
