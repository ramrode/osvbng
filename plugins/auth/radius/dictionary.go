// Copyright 2026 The osvbng Authors
// Licensed under the GNU General Public License v3.0 or later.
// SPDX-License-Identifier: GPL-3.0-or-later

package radius

import "github.com/veesix-networks/osvbng/pkg/aaa"

// DefaultVendorID is the enterprise number under which osvbng's own
// vendor-specific attributes are encoded. 32473 is the IANA-reserved
// documentation enterprise number (RFC 5612), used as an interim value
// until the project's own PEN assignment lands; deployments can pin a
// different number with the plugin's vendor_id setting, and the
// matching FreeRADIUS definitions live in contrib/freeradius/.
const DefaultVendorID = 32473

const (
	vsaL2GWHandoffGroup = 1
	vsaL2GWSVLAN        = 2
	vsaL2GWCVLAN        = 3
)

// osvbngVendorMappings are the built-in tier-2 response mappings under
// the osvbng vendor id: Access-Accept VSAs decoded into internal
// attributes.
func osvbngVendorMappings(vendorID uint32) []vendorMapping {
	return []vendorMapping{
		{vendorID: vendorID, vendorType: vsaL2GWHandoffGroup, internal: aaa.AttrL2GWHandoffGroup, decode: decodeVSAString},
		{vendorID: vendorID, vendorType: vsaL2GWSVLAN, internal: aaa.AttrL2GWSVLAN, decode: decodeVSAString},
		{vendorID: vendorID, vendorType: vsaL2GWCVLAN, internal: aaa.AttrL2GWCVLAN, decode: decodeVSAString},
	}
}

// osvbngAcctMappings are the built-in accounting mappings: resolved
// internal attributes reported back to the OSS as VSAs on
// Start/Interim/Stop.
func osvbngAcctMappings(vendorID uint32) []compiledRequestMapping {
	return []compiledRequestMapping{
		{internal: aaa.AttrL2GWHandoffGroup, vendorID: vendorID, vendorType: vsaL2GWHandoffGroup},
		{internal: aaa.AttrL2GWSVLAN, vendorID: vendorID, vendorType: vsaL2GWSVLAN},
		{internal: aaa.AttrL2GWCVLAN, vendorID: vendorID, vendorType: vsaL2GWCVLAN},
	}
}
