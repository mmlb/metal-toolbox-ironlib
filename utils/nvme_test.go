package utils

import (
	"context"
	"fmt"
	"testing"

	"github.com/bmc-toolbox/common"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func Test_NvmeComponents(t *testing.T) {
	expected := []*common.Drive{
		{Common: common.Common{
			Serial: "Z9DF70I8FY3L", Vendor: "TOSHIBA", Model: "KXG60ZNV256G TOSHIBA", Description: "KXG60ZNV256G TOSHIBA", Firmware: &common.Firmware{Installed: "AGGA4104"}, ProductName: "NULL",
			Metadata: map[string]string{
				"Block Erase Sanitize Operation Supported":                          "false",
				"Crypto Erase Applies to All/Single Namespace(s) (t:All, f:Single)": "false",
				"Crypto Erase Sanitize Operation Supported":                         "false",
				"Crypto Erase Supported as part of Secure Erase":                    "true",
				"Format Applies to All/Single Namespace(s) (t:All, f:Single)":       "false",
				"No-Deallocate After Sanitize bit in Sanitize command Supported":    "false",
				"Overwrite Sanitize Operation Supported":                            "false",
			},
		}},
		{Common: common.Common{
			Serial: "Z9DF70I9FY3L", Vendor: "TOSHIBA", Model: "KXG60ZNV256G TOSHIBA", Description: "KXG60ZNV256G TOSHIBA", Firmware: &common.Firmware{Installed: "AGGA4104"}, ProductName: "NULL",
			Metadata: map[string]string{
				"Block Erase Sanitize Operation Supported":                          "false",
				"Crypto Erase Applies to All/Single Namespace(s) (t:All, f:Single)": "false",
				"Crypto Erase Sanitize Operation Supported":                         "false",
				"Crypto Erase Supported as part of Secure Erase":                    "true",
				"Format Applies to All/Single Namespace(s) (t:All, f:Single)":       "false",
				"No-Deallocate After Sanitize bit in Sanitize command Supported":    "false",
				"Overwrite Sanitize Operation Supported":                            "false",
			},
		}},
	}

	n := NewFakeNvme()

	drives, err := n.Drives(context.TODO())
	if err != nil {
		t.Error(err)
	}

	assert.Equal(t, expected, drives)
}

func Test_NvmeDriveCapabilities(t *testing.T) {
	n := NewFakeNvme()

	d := &nvmeDeviceAttributes{DevicePath: "/dev/nvme0"}

	capabilities, err := n.DriveCapabilities(context.TODO(), d.DevicePath)
	if err != nil {
		t.Fatal(err)
	}

	assert.Equal(t, fixtureNvmeDeviceCapabilities, capabilities)
}

var fixtureNvmeDeviceCapabilities = []*common.Capability{
	{Name: "fmns", Description: "Format Applies to All/Single Namespace(s) (t:All, f:Single)", Enabled: false},
	{Name: "cens", Description: "Crypto Erase Applies to All/Single Namespace(s) (t:All, f:Single)", Enabled: false},
	{Name: "cese", Description: "Crypto Erase Supported as part of Secure Erase", Enabled: true},
	{Name: "cer", Description: "Crypto Erase Sanitize Operation Supported", Enabled: false},
	{Name: "ber", Description: "Block Erase Sanitize Operation Supported", Enabled: false},
	{Name: "owr", Description: "Overwrite Sanitize Operation Supported", Enabled: false},
	{Name: "ndi", Description: "No-Deallocate After Sanitize bit in Sanitize command Supported", Enabled: false},
}

func Test_NvmeParseFna(t *testing.T) {
	// These are laid out so if you squint and pretend false/true are 0/1 they match the bit pattern of the int
	// Its a map so order doesn't matter but I think it makes it easier to match a broken test to the code
	wants := []map[string]bool{
		{"cese": false, "cens": false, "fmns": false},
		{"cese": false, "cens": false, "fmns": true},
		{"cese": false, "cens": true, "fmns": false},
		{"cese": false, "cens": true, "fmns": true},
		{"cese": true, "cens": false, "fmns": false},
		{"cese": true, "cens": false, "fmns": true},
		{"cese": true, "cens": true, "fmns": false},
		{"cese": true, "cens": true, "fmns": true},
	}
	for i, want := range wants {
		t.Run(fmt.Sprintf("0b0...%03b", i), func(t *testing.T) {
			caps := parseFna(uint(i))
			require.Len(t, caps, len(want))
			for _, cap := range caps {
				require.Equal(t, want[cap.Name], cap.Enabled)
			}
		})
	}
}

// This test tests parsing sanicap values that we're interested in
// The test is broken up into 3 sections
// First it iterates through all the bit patterns for the low bits that we're interested in, namely 0-7 aka 0b000-b111
// Then it tests the high bits we're interested in: bits 29-31.
// Finally it tests that if bit 31 (the 32nd bit) is set the function panics because we don't know what to do with that bit's data.
func Test_NvmeParseSanicap(t *testing.T) {
	// These are laid out so if you squint and pretend false/true are 0/1 they match the bit pattern of the int
	// Its a map so order doesn't matter but I think it makes it easier to match a broken test to the code

	// lower bits only
	wants := []map[string]bool{
		{"owr": false, "ber": false, "cer": false},
		{"owr": false, "ber": false, "cer": true},
		{"owr": false, "ber": true, "cer": false},
		{"owr": false, "ber": true, "cer": true},
		{"owr": true, "ber": false, "cer": false},
		{"owr": true, "ber": false, "cer": true},
		{"owr": true, "ber": true, "cer": false},
		{"owr": true, "ber": true, "cer": true},
	}
	for i, want := range wants {
		// not testing ndi yet but its being returned
		// don't want to add it above to avoid noise
		want["ndi"] = false
		t.Run(fmt.Sprintf("0b000...%03b", i), func(t *testing.T) {
			caps, err := parseSanicap(uint(i))
			require.NoError(t, err)
			require.Len(t, caps, len(want))
			for _, cap := range caps {
				require.Equal(t, want[cap.Name], cap.Enabled)
			}
		})
	}

	// higher bits only
	wants = []map[string]bool{
		{}, // already tested above and want to avoid t.Run duplicate name handling
		{"ndi": true},
		{"nodmmas": false, "ndi": false},
		{"nodmmas": false, "ndi": true},
		{"nodmmas": true, "ndi": false},
		{"nodmmas": true, "ndi": true},
	}
	for i, want := range wants {
		if i == 0 {
			// already tested above
			continue
		}
		// not testing these now but they are being returned
		// don't want to add them above to avoid noise
		want["owr"] = false
		want["ber"] = false
		want["cer"] = false
		i = (i << 29)
		name := fmt.Sprintf("%032b", i)
		name = "0b" + name[:3] + "...000"
		t.Run(name, func(t *testing.T) {
			caps, err := parseSanicap(uint(i))
			require.NoError(t, err)
			require.Len(t, caps, len(want))
			for _, cap := range caps {
				require.Equal(t, want[cap.Name], cap.Enabled)
			}
		})
	}

	i := 0b11 << 30
	name := fmt.Sprintf("%032b", i)
	name = "0b" + name[:3] + "...000"
	t.Run(name, func(t *testing.T) {
		caps, err := parseSanicap(uint(i))
		require.Error(t, err)
		require.Nil(t, caps)
	})
}
