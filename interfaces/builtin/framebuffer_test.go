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

type FramebufferInterfaceSuite struct {
	iface interfaces.Interface
	slot  *interfaces.Slot
	plug  *interfaces.Plug
}

var _ = Suite(&FramebufferInterfaceSuite{
	iface: builtin.MustInterface("framebuffer"),
})

func (s *FramebufferInterfaceSuite) SetUpTest(c *C) {
	// Mock for OS Snap
	osSnapInfo := snaptest.MockInfo(c, `
name: ubuntu-core
type: os
slots:
  test-framebuffer:
    interface: framebuffer
`, nil)
	s.slot = &interfaces.Slot{SlotInfo: osSnapInfo.Slots["test-framebuffer"]}

	// Snap Consumers
	consumingSnapInfo := snaptest.MockInfo(c, `
name: client-snap
apps:
  app-accessing-framebuffer:
    command: foo
    plugs: [framebuffer]
`, nil)
	s.plug = &interfaces.Plug{PlugInfo: consumingSnapInfo.Plugs["framebuffer"]}
}

func (s *FramebufferInterfaceSuite) TestName(c *C) {
	c.Assert(s.iface.Name(), Equals, "framebuffer")
}

func (s *FramebufferInterfaceSuite) TestSanitizeSlot(c *C) {
	c.Assert(s.slot.Sanitize(s.iface), IsNil)
	slot := &interfaces.Slot{SlotInfo: &snap.SlotInfo{
		Snap:      &snap.Info{SuggestedName: "some-snap"},
		Name:      "framebuffer",
		Interface: "framebuffer",
	}}
	c.Assert(slot.Sanitize(s.iface), ErrorMatches,
		"framebuffer slots are reserved for the core snap")
}

func (s *FramebufferInterfaceSuite) TestSanitizePlug(c *C) {
	c.Assert(s.plug.Sanitize(s.iface), IsNil)
}

func (s *FramebufferInterfaceSuite) TestUsedSecuritySystems(c *C) {
	expectedSnippet1 := `
# Description: Allow reading and writing to the universal framebuffer (/dev/fb*) which
# gives privileged access to the console framebuffer.

/dev/fb[0-9]* rw,
/run/udev/data/c29:[0-9]* r,
`
	// connected plugs have a non-nil security snippet for apparmor
	apparmorSpec := &apparmor.Specification{}
	err := apparmorSpec.AddConnectedPlug(s.iface, s.plug, nil, s.slot, nil)
	c.Assert(err, IsNil)
	c.Assert(apparmorSpec.SecurityTags(), DeepEquals, []string{"snap.client-snap.app-accessing-framebuffer"})
	aasnippet := apparmorSpec.SnippetForTag("snap.client-snap.app-accessing-framebuffer")
	c.Assert(aasnippet, Equals, expectedSnippet1, Commentf("\nexpected:\n%s\nfound:\n%s", expectedSnippet1, aasnippet))
}

func (s *FramebufferInterfaceSuite) TestInterfaces(c *C) {
	c.Check(builtin.Interfaces(), testutil.DeepContains, s.iface)
}
