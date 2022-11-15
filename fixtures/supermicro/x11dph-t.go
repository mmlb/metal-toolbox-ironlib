package supermicro

import (
	"github.com/bmc-toolbox/common"
	"github.com/metal-toolbox/ironlib/model"
)

// nolint // test data
var (
	Testdata_X11DPH_T_Components = []*model.Component{
		{
			Serial:            "",
			Vendor:            "Supermicro",
			Type:              "",
			Model:             "Supermicro",
			Name:              "CPLD",
			Slug:              "CPLD",
			FirmwareInstalled: "02.b6.04",
			FirmwareAvailable: "",
		},
		{
			Serial:            "",
			Vendor:            "Supermicro",
			Type:              "",
			Model:             "Supermicro",
			Name:              "BIOS",
			Slug:              "BIOS",
			FirmwareInstalled: "3.3",
			FirmwareAvailable: "",
		},
		{
			Serial:            "",
			Vendor:            "Supermicro",
			Type:              "",
			Model:             "Supermicro",
			Name:              "BMC",
			Slug:              "BMC",
			FirmwareInstalled: "1.71.11",
			FirmwareAvailable: "",
		},
		{
			Serial:            "20032613EE70",
			Vendor:            "Micron",
			Type:              "Sata SSD",
			Model:             "Micron_5200_MTFDDAK960TDN",
			Name:              "Micron_5200_MTFDDAK960TDN",
			Slug:              "Drive",
			FirmwareInstalled: "D1MU020",
			FirmwareAvailable: "",
		},
		{
			Serial:            "20032613D813",
			Vendor:            "Micron",
			Type:              "Sata SSD",
			Model:             "Micron_5200_MTFDDAK960TDN",
			Name:              "Micron_5200_MTFDDAK960TDN",
			Slug:              "Drive",
			FirmwareInstalled: "D1MU020",
			FirmwareAvailable: "",
		},
		{
			Serial:            "VDJKG9LK",
			Vendor:            "hgst",
			Type:              "Unknown",
			Model:             "HGST HUS728T8TALE6L4",
			Name:              "HGST HUS728T8TALE6L4",
			Slug:              "Drive",
			FirmwareInstalled: "V8GNW460",
			FirmwareAvailable: "",
		},
		{
			Serial:            "VDJBVPXD",
			Vendor:            "hgst",
			Type:              "Unknown",
			Model:             "HGST HUS728T8TALE6L4",
			Name:              "HGST HUS728T8TALE6L4",
			Slug:              "Drive",
			FirmwareInstalled: "V8GNW460",
			FirmwareAvailable: "",
		},
		{
			Serial:            "VDJJ7T3K",
			Vendor:            "hgst",
			Type:              "Unknown",
			Model:             "HGST HUS728T8TALE6L4",
			Name:              "HGST HUS728T8TALE6L4",
			Slug:              "Drive",
			FirmwareInstalled: "V8GNW460",
			FirmwareAvailable: "",
		},
		{
			Serial:            "VDJL8U5K",
			Vendor:            "hgst",
			Type:              "Unknown",
			Model:             "HGST HUS728T8TALE6L4",
			Name:              "HGST HUS728T8TALE6L4",
			Slug:              "Drive",
			FirmwareInstalled: "V8GNW460",
			FirmwareAvailable: "",
		},
		{
			Serial:            "VDJBP7SD",
			Vendor:            "hgst",
			Type:              "Unknown",
			Model:             "HGST HUS728T8TALE6L4",
			Name:              "HGST HUS728T8TALE6L4",
			Slug:              "Drive",
			FirmwareInstalled: "V8GNW460",
			FirmwareAvailable: "",
		},
		{
			Serial:            "VDJKVBKK",
			Vendor:            "hgst",
			Type:              "Unknown",
			Model:             "HGST HUS728T8TALE6L4",
			Name:              "HGST HUS728T8TALE6L4",
			Slug:              "Drive",
			FirmwareInstalled: "V8GNW460",
			FirmwareAvailable: "",
		},
		{
			Serial:            "VDJK7E2K",
			Vendor:            "hgst",
			Type:              "Unknown",
			Model:             "HGST HUS728T8TALE6L4",
			Name:              "HGST HUS728T8TALE6L4",
			Slug:              "Drive",
			FirmwareInstalled: "V8GNW460",
			FirmwareAvailable: "",
		},
		{
			Serial:            "VDJL2NUK",
			Vendor:            "hgst",
			Type:              "Unknown",
			Model:             "HGST HUS728T8TALE6L4",
			Name:              "HGST HUS728T8TALE6L4",
			Slug:              "Drive",
			FirmwareInstalled: "V8GNW460",
			FirmwareAvailable: "",
		},
		{
			Serial:            "VDJL8MVK",
			Vendor:            "hgst",
			Type:              "Unknown",
			Model:             "HGST HUS728T8TALE6L4",
			Name:              "HGST HUS728T8TALE6L4",
			Slug:              "Drive",
			FirmwareInstalled: "V8GNW460",
			FirmwareAvailable: "",
		},
		{
			Serial:            "VDJHR67K",
			Vendor:            "hgst",
			Type:              "Unknown",
			Model:             "HGST HUS728T8TALE6L4",
			Name:              "HGST HUS728T8TALE6L4",
			Slug:              "Drive",
			FirmwareInstalled: "V8GNW460",
			FirmwareAvailable: "",
		},
		{
			Serial:            "VDJJ51YK",
			Vendor:            "hgst",
			Type:              "Unknown",
			Model:             "HGST HUS728T8TALE6L4",
			Name:              "HGST HUS728T8TALE6L4",
			Slug:              "Drive",
			FirmwareInstalled: "V8GNW460",
			FirmwareAvailable: "",
		},
		{
			Serial:            "VDJLNERD",
			Vendor:            "hgst",
			Type:              "Unknown",
			Model:             "HGST HUS728T8TALE6L4",
			Name:              "HGST HUS728T8TALE6L4",
			Slug:              "Drive",
			FirmwareInstalled: "V8GNW460",
			FirmwareAvailable: "",
		},
		{
			Serial:            "Z9DF71TJFY3L",
			Vendor:            "Toshiba",
			Type:              "NVMe PCIe SSD",
			Model:             "KXG60ZNV256G TOSHIBA",
			Name:              "KXG60ZNV256G TOSHIBA",
			Slug:              "Drive",
			FirmwareInstalled: "AGGA4104",
			FirmwareAvailable: "",
		},
		{
			Serial:            "Z9DF71YGFY3L",
			Vendor:            "Toshiba",
			Type:              "NVMe PCIe SSD",
			Model:             "KXG60ZNV256G TOSHIBA",
			Name:              "KXG60ZNV256G TOSHIBA",
			Slug:              "Drive",
			FirmwareInstalled: "AGGA4104",
			FirmwareAvailable: "",
		},
		{
			Serial:            "50030480256efa01",
			Vendor:            "LSI",
			Type:              "",
			Model:             "LSI3008-IT",
			Name:              "Serial Attached SCSI controller",
			Slug:              "StorageController",
			FirmwareInstalled: "16.00.01.00",
			FirmwareAvailable: "",
		},
		{
			Serial:            "0c:42:a1:3d:bd:e0",
			Vendor:            "Mellanox",
			Type:              "",
			Model:             "ConnectX4LX",
			Name:              "ConnectX-4 Lx EN network interface card; 25GbE dual-port SFP28; PCIe3.0 x8; ROHS R6",
			Slug:              "NIC",
			FirmwareInstalled: "14.27.1016",
			FirmwareAvailable: "14.28.2006",
			Metadata: map[string]string{
				"firmware_pxe_installed":  "3.5.0901",
				"firmware_pxe_available":  "3.6.0102",
				"firmware_uefi_installed": "14.20.0019",
				"firmware_uefi_available": "14.21.0017",
			},
			Oem:             false,
			FirmwareManaged: true,
		},
	}

	// *model.Device without firmware information
	Testdata_X11DPH_T_Inventory = &common.Device{
		Common: common.Common{
			Oem:         false,
			Description: "",
			Vendor:      "supermicro",
			Model:       "x11dph-t",
			Serial:      "A401031X0502795",
			ProductName: "",
			Firmware:    nil,
			Status:      nil,
		},
		HardwareType: "",
		Chassis:      "",
		BIOS: &common.BIOS{
			Common: common.Common{
				Oem:         false,
				Description: "BIOS",
				Vendor:      "Supermicro",
				Model:       "x11dph-t",
				Serial:      "",
				ProductName: "",
				Firmware: &common.Firmware{
					Installed:  "3.4",
					Available:  "",
					SoftwareID: "",
					Previous:   nil,
					Metadata: map[string]string{
						"build_date": "11/03/2020",
					},
				},
				Status: nil,
			},
			SizeBytes:     65536,
			CapacityBytes: 67108864,
		},
		BMC: &common.BMC{
			Common: common.Common{
				Oem:         false,
				Description: "BMC",
				Vendor:      "Supermicro",
				Model:       "x11dph-t",
				Serial:      "",
				ProductName: "",
				Firmware: &common.Firmware{
					Installed:  "1.71.11",
					Available:  "",
					SoftwareID: "",
					Previous:   nil,
					Metadata: map[string]string{
						"build_date": "10/25/2019",
					},
				},
				Status: nil,
			},
			ID:  "",
			NIC: &common.NIC{},
		},
		Mainboard: &common.Mainboard{
			Common: common.Common{
				Oem:         false,
				Description: "Motherboard",
				Vendor:      "Supermicro",
				Model:       "X11DPH-T",
				Serial:      "ZM19AS026837",
				ProductName: "X11DPH-T",
				Firmware:    nil,
				Status:      nil,
			},
			PhysicalID: "0",
		},
		CPLDs: []*common.CPLD{
			{
				Common: common.Common{
					Oem:         false,
					Description: "CPLD",
					Vendor:      "Supermicro",
					Model:       "x11dph-t",
					Serial:      "",
					ProductName: "",
					Firmware: &common.Firmware{
						Installed:  "02.b6.04",
						Available:  "",
						SoftwareID: "",
						Previous:   nil,
					},
					Status: nil,
				},
			},
		},
		TPMs: []*common.TPM{
			{
				Common: common.Common{
					Oem:         false,
					Description: "",
					Vendor:      "",
					Model:       "",
					Serial:      "",
					ProductName: "",
					Firmware: &common.Firmware{
						Installed:  "",
						Available:  "",
						SoftwareID: "",
						Previous:   nil,
					},
					Status: nil,
					Metadata: map[string]string{
						"Specification Version": "",
					},
				},
				InterfaceType: "",
			},
		},
		GPUs: []*common.GPU{},
		CPUs: []*common.CPU{
			{
				Common: common.Common{
					Oem:         false,
					Description: "CPU",
					Vendor:      "Intel Corp.",
					Model:       "Intel(R) Xeon(R) Silver 4214 CPU @ 2.20GHz",
					Serial:      "",
					ProductName: "Intel(R) Xeon(R) Silver 4214 CPU @ 2.20GHz",
					Firmware:    &common.Firmware{Installed: "83898113", Metadata: map[string]string{}},
					Status:      nil,
				},
				ID:           "",
				Slot:         "CPU1",
				Architecture: "",
				ClockSpeedHz: 100000000,
				Cores:        12,
				Threads:      24,
			},
			{
				Common: common.Common{
					Oem:         false,
					Description: "CPU",
					Vendor:      "Intel Corp.",
					Model:       "Intel(R) Xeon(R) Silver 4214 CPU @ 2.20GHz",
					Serial:      "",
					ProductName: "Intel(R) Xeon(R) Silver 4214 CPU @ 2.20GHz",
					Firmware:    &common.Firmware{Installed: "83898113", Metadata: map[string]string{}},
					Status:      nil,
				},
				ID:           "",
				Slot:         "CPU2",
				Architecture: "",
				ClockSpeedHz: 100000000,
				Cores:        12,
				Threads:      24,
			},
		},
		Memory: []*common.Memory{
			{
				Common: common.Common{
					Oem:         false,
					Description: "DIMM DDR4 Synchronous Registered (Buffered) 2933 MHz (0.3 ns) [empty]",
					Vendor:      "SK Hynix",
					Model:       "HMA82GR7CJR8N-WM",
					Serial:      "53963974",
					ProductName: "HMA82GR7CJR8N-WM",
					Firmware:    nil,
					Status:      nil,
				},
				ID:           "",
				Slot:         "P1-DIMMA1",
				Type:         "",
				SizeBytes:    0,
				FormFactor:   "",
				PartNumber:   "",
				ClockSpeedHz: 2933000000,
			},
			{
				Common: common.Common{
					Oem:         false,
					Description: "DIMM DDR4 Synchronous Registered (Buffered) 2933 MHz (0.3 ns) [empty]",
					Vendor:      "SK Hynix",
					Model:       "HMA82GR7CJR8N-WM",
					Serial:      "53963967",
					ProductName: "HMA82GR7CJR8N-WM",
					Firmware:    nil,
					Status:      nil,
				},
				ID:           "",
				Slot:         "P1-DIMMB1",
				Type:         "",
				SizeBytes:    0,
				FormFactor:   "",
				PartNumber:   "",
				ClockSpeedHz: 2933000000,
			},
			{
				Common: common.Common{
					Oem:         false,
					Description: "DIMM DDR4 Synchronous Registered (Buffered) 2933 MHz (0.3 ns) [empty]",
					Vendor:      "SK Hynix",
					Model:       "HMA82GR7CJR8N-WM",
					Serial:      "53963971",
					ProductName: "HMA82GR7CJR8N-WM",
					Firmware:    nil,
					Status:      nil,
				},
				ID:           "",
				Slot:         "P1-DIMMC1",
				Type:         "",
				SizeBytes:    0,
				FormFactor:   "",
				PartNumber:   "",
				ClockSpeedHz: 2933000000,
			},
			{
				Common: common.Common{
					Oem:         false,
					Description: "DIMM DDR4 Synchronous Registered (Buffered) 2933 MHz (0.3 ns) [empty]",
					Vendor:      "SK Hynix",
					Model:       "HMA82GR7CJR8N-WM",
					Serial:      "53963941",
					ProductName: "HMA82GR7CJR8N-WM",
					Firmware:    nil,
					Status:      nil,
				},
				ID:           "",
				Slot:         "P1-DIMMD1",
				Type:         "",
				SizeBytes:    0,
				FormFactor:   "",
				PartNumber:   "",
				ClockSpeedHz: 2933000000,
			},
			{
				Common: common.Common{
					Oem:         false,
					Description: "DIMM DDR4 Synchronous Registered (Buffered) 2933 MHz (0.3 ns) [empty]",
					Vendor:      "SK Hynix",
					Model:       "HMA82GR7CJR8N-WM",
					Serial:      "53963943",
					ProductName: "HMA82GR7CJR8N-WM",
					Firmware:    nil,
					Status:      nil,
				},
				ID:           "",
				Slot:         "P1-DIMME1",
				Type:         "",
				SizeBytes:    0,
				FormFactor:   "",
				PartNumber:   "",
				ClockSpeedHz: 2933000000,
			},
			{
				Common: common.Common{
					Oem:         false,
					Description: "DIMM DDR4 Synchronous Registered (Buffered) 2933 MHz (0.3 ns) [empty]",
					Vendor:      "SK Hynix",
					Model:       "HMA82GR7CJR8N-WM",
					Serial:      "53963950",
					ProductName: "HMA82GR7CJR8N-WM",
					Firmware:    nil,
					Status:      nil,
				},
				ID:           "",
				Slot:         "P1-DIMMF1",
				Type:         "",
				SizeBytes:    0,
				FormFactor:   "",
				PartNumber:   "",
				ClockSpeedHz: 2933000000,
			},
			{
				Common: common.Common{
					Oem:         false,
					Description: "DIMM DDR4 Synchronous Registered (Buffered) 2933 MHz (0.3 ns) [empty]",
					Vendor:      "SK Hynix",
					Model:       "HMA82GR7CJR8N-WM",
					Serial:      "5396396B",
					ProductName: "HMA82GR7CJR8N-WM",
					Firmware:    nil,
					Status:      nil,
				},
				ID:           "",
				Slot:         "P2-DIMMA1",
				Type:         "",
				SizeBytes:    0,
				FormFactor:   "",
				PartNumber:   "",
				ClockSpeedHz: 2933000000,
			},
			{
				Common: common.Common{
					Oem:         false,
					Description: "DIMM DDR4 Synchronous Registered (Buffered) 2933 MHz (0.3 ns) [empty]",
					Vendor:      "SK Hynix",
					Model:       "HMA82GR7CJR8N-WM",
					Serial:      "53963981",
					ProductName: "HMA82GR7CJR8N-WM",
					Firmware:    nil,
					Status:      nil,
				},
				ID:           "",
				Slot:         "P2-DIMMB1",
				Type:         "",
				SizeBytes:    0,
				FormFactor:   "",
				PartNumber:   "",
				ClockSpeedHz: 2933000000,
			},
			{
				Common: common.Common{
					Oem:         false,
					Description: "DIMM DDR4 Synchronous Registered (Buffered) 2933 MHz (0.3 ns) [empty]",
					Vendor:      "SK Hynix",
					Model:       "HMA82GR7CJR8N-WM",
					Serial:      "53963958",
					ProductName: "HMA82GR7CJR8N-WM",
					Firmware:    nil,
					Status:      nil,
				},
				ID:           "",
				Slot:         "P2-DIMMC1",
				Type:         "",
				SizeBytes:    0,
				FormFactor:   "",
				PartNumber:   "",
				ClockSpeedHz: 2933000000,
			},
			{
				Common: common.Common{
					Oem:         false,
					Description: "DIMM DDR4 Synchronous Registered (Buffered) 2933 MHz (0.3 ns) [empty]",
					Vendor:      "SK Hynix",
					Model:       "HMA82GR7CJR8N-WM",
					Serial:      "53963964",
					ProductName: "HMA82GR7CJR8N-WM",
					Firmware:    nil,
					Status:      nil,
				},
				ID:           "",
				Slot:         "P2-DIMMD1",
				Type:         "",
				SizeBytes:    0,
				FormFactor:   "",
				PartNumber:   "",
				ClockSpeedHz: 2933000000,
			},
			{
				Common: common.Common{
					Oem:         false,
					Description: "DIMM DDR4 Synchronous Registered (Buffered) 2933 MHz (0.3 ns) [empty]",
					Vendor:      "SK Hynix",
					Model:       "HMA82GR7CJR8N-WM",
					Serial:      "53963959",
					ProductName: "HMA82GR7CJR8N-WM",
					Firmware:    nil,
					Status:      nil,
				},
				ID:           "",
				Slot:         "P2-DIMME1",
				Type:         "",
				SizeBytes:    0,
				FormFactor:   "",
				PartNumber:   "",
				ClockSpeedHz: 2933000000,
			},
			{
				Common: common.Common{
					Oem:         false,
					Description: "DIMM DDR4 Synchronous Registered (Buffered) 2933 MHz (0.3 ns) [empty]",
					Vendor:      "SK Hynix",
					Model:       "HMA82GR7CJR8N-WM",
					Serial:      "53963965",
					ProductName: "HMA82GR7CJR8N-WM",
					Firmware:    nil,
					Status:      nil,
				},
				ID:           "",
				Slot:         "P2-DIMMF1",
				Type:         "",
				SizeBytes:    0,
				FormFactor:   "",
				PartNumber:   "",
				ClockSpeedHz: 2933000000,
			},
		},
		NICs: []*common.NIC{
			{
				Common: common.Common{
					Oem:         false,
					Description: "Ethernet interface",
					Vendor:      "Intel Corporation",
					Model:       "Ethernet Connection X722 for 10GBASE-T",
					Serial:      "ac:1f:6b:dc:19:c2",
					ProductName: "Ethernet Connection X722 for 10GBASE-T",
					Firmware: &common.Firmware{
						Installed:  "1.1747.0",
						Available:  "",
						SoftwareID: "",
						Previous:   nil,
					},
					Status: nil,
					Metadata: map[string]string{
						"driver":   "i40e",
						"firmware": "3.31 0x80000c92 1.1747.0",
						"link":     "no",
					},
				},
				ID:          "",
				Description: "Ethernet interface",
				SpeedBits:   10000000000,
				PhysicalID:  "0",
				BusInfo:     "pci@0000:1a:00.0",
				MacAddress:  "",
			},
			{
				Common: common.Common{
					Oem:         false,
					Description: "Ethernet interface",
					Vendor:      "Intel Corporation",
					Model:       "Ethernet Connection X722 for 10GBASE-T",
					Serial:      "ac:1f:6b:dc:19:c3",
					ProductName: "Ethernet Connection X722 for 10GBASE-T",
					Firmware: &common.Firmware{
						Installed:  "1.1747.0",
						Available:  "",
						SoftwareID: "",
						Previous:   nil,
					},
					Status: nil,
					Metadata: map[string]string{
						"driver":   "i40e",
						"firmware": "3.31 0x80000c92 1.1747.0",
						"link":     "no",
					},
				},
				ID:          "",
				Description: "Ethernet interface",
				SpeedBits:   10000000000,
				PhysicalID:  "0.1",
				BusInfo:     "pci@0000:1a:00.1",
				MacAddress:  "",
			},
			{
				Common: common.Common{
					Oem:         false,
					Description: "ConnectX-4 Lx EN network interface card; 25GbE dual-port SFP28; PCIe3.0 x8; ROHS R6",
					Vendor:      "mellanox",
					Model:       "MCX4121A-ACA_Ax",
					Serial:      "0c:42:a1:3d:bd:e0",
					ProductName: "MT27710 Family [ConnectX-4 Lx]",
					Firmware: &common.Firmware{
						Installed:  "14.29.2002",
						Available:  "14.28.2006",
						SoftwareID: "",
						Previous:   nil,
						Metadata: map[string]string{
							"firmware_pxe_available":  "3.6.0102",
							"firmware_pxe_installed":  "3.6.0204",
							"firmware_uefi_available": "14.21.0017",
							"firmware_uefi_installed": "14.22.0016",
						},
					},
					Status: nil,
					Metadata: map[string]string{
						"driver":   "mlx5_core",
						"duplex":   "full",
						"firmware": "14.27.1016 (MT_2420110034)",
						"link":     "yes",
						"speed":    "10Gbit/s",
					},
				},
				ID:          "",
				Description: "Ethernet interface",
				SpeedBits:   0,
				PhysicalID:  "0",
				BusInfo:     "pci@0000:d8:00.0",
				MacAddress:  "",
			},
		},
		Drives: []*common.Drive{
			{
				Common: common.Common{
					Oem:         false,
					Description: "ATA Disk",
					Vendor:      "micron",
					Model:       "Micron_5200_MTFDDAK960TDN",
					Serial:      "20032613EE70",
					ProductName: "Micron_5200_MTFD",
					LogicalName: "/dev/sda",
					Firmware: &common.Firmware{
						Installed:  "D1MU020",
						Available:  "",
						SoftwareID: "",
						Previous:   nil,
					},
					Status: nil,
				},
				ID:                  "",
				OemID:               "",
				Type:                "Sata-SSD",
				StorageController:   "",
				BusInfo:             "scsi@5:0.0.0",
				WWN:                 "",
				Protocol:            "",
				SmartStatus:         "ok",
				SmartErrors:         nil,
				CapacityBytes:       960197124096,
				BlockSizeBytes:      0,
				CapableSpeedGbps:    0,
				NegotiatedSpeedGbps: 0,
			},
			{
				Common: common.Common{
					Oem:         false,
					Description: "ATA Disk",
					Vendor:      "micron",
					Model:       "Micron_5200_MTFDDAK960TDN",
					Serial:      "20032613D813",
					ProductName: "Micron_5200_MTFD",
					LogicalName: "/dev/sdb",
					Firmware: &common.Firmware{
						Installed:  "D1MU020",
						Available:  "",
						SoftwareID: "",
						Previous:   nil,
					},
					Status: nil,
				},
				ID:                  "",
				OemID:               "",
				Type:                "Sata-SSD",
				StorageController:   "",
				BusInfo:             "scsi@6:0.0.0",
				WWN:                 "",
				Protocol:            "",
				SmartStatus:         "ok",
				SmartErrors:         nil,
				CapacityBytes:       960197124096,
				BlockSizeBytes:      0,
				CapableSpeedGbps:    0,
				NegotiatedSpeedGbps: 0,
			},
			{
				Common: common.Common{
					Oem:         false,
					Description: "NVMe device",
					Vendor:      "toshiba",
					Model:       "KXG60ZNV256G TOSHIBA",
					Serial:      "Z9DF71TJFY3L",
					ProductName: "KXG60ZNV256G TOSHIBA",
					LogicalName: "/dev/nvme0",
					Firmware: &common.Firmware{
						Installed:  "AGGA4104",
						Available:  "",
						SoftwareID: "",
						Previous:   nil,
					},
					Status: nil,
				},
				ID:                  "",
				OemID:               "",
				Type:                "NVMe-PCIe-SSD",
				StorageController:   "",
				BusInfo:             "pci@0000:3b:00.0",
				WWN:                 "",
				Protocol:            "",
				SmartStatus:         "ok",
				SmartErrors:         nil,
				CapacityBytes:       0,
				BlockSizeBytes:      0,
				CapableSpeedGbps:    0,
				NegotiatedSpeedGbps: 0,
			},
			{
				Common: common.Common{
					Oem:         false,
					Description: "NVMe device",
					Vendor:      "toshiba",
					Model:       "KXG60ZNV256G TOSHIBA",
					Serial:      "Z9DF71YGFY3L",
					ProductName: "KXG60ZNV256G TOSHIBA",
					LogicalName: "/dev/nvme1",
					Firmware: &common.Firmware{
						Installed:  "AGGA4104",
						Available:  "",
						SoftwareID: "",
						Previous:   nil,
					},
					Status: nil,
				},
				ID:                  "",
				OemID:               "",
				Type:                "NVMe-PCIe-SSD",
				StorageController:   "",
				BusInfo:             "pci@0000:3c:00.0",
				WWN:                 "",
				Protocol:            "",
				SmartStatus:         "ok",
				SmartErrors:         nil,
				CapacityBytes:       0,
				BlockSizeBytes:      0,
				CapableSpeedGbps:    0,
				NegotiatedSpeedGbps: 0,
			},
			{
				Common: common.Common{
					Oem:         false,
					Description: "ATA Disk",
					Vendor:      "hgst",
					Model:       "HGST HUS728T8TALE6L4",
					Serial:      "VDJL8MVK",
					ProductName: "HGST HUS728T8TAL",
					LogicalName: "/dev/sdk",
					Firmware: &common.Firmware{
						Installed:  "V8GNW460",
						Available:  "",
						SoftwareID: "",
						Previous:   nil,
					},
					Status: nil,
				},
				ID:                  "",
				OemID:               "",
				Type:                "Sata-HDD",
				StorageController:   "",
				BusInfo:             "scsi@4:0.8.0",
				WWN:                 "",
				Protocol:            "",
				SmartStatus:         "ok",
				SmartErrors:         nil,
				CapacityBytes:       8001563222016,
				BlockSizeBytes:      0,
				CapableSpeedGbps:    0,
				NegotiatedSpeedGbps: 0,
			},
			{
				Common: common.Common{
					Oem:         false,
					Description: "ATA Disk",
					Vendor:      "hgst",
					Model:       "HGST HUS728T8TALE6L4",
					Serial:      "VDJHR67K",
					ProductName: "HGST HUS728T8TAL",
					LogicalName: "/dev/sdl",
					Firmware: &common.Firmware{
						Installed:  "V8GNW460",
						Available:  "",
						SoftwareID: "",
						Previous:   nil,
					},
					Status: nil,
				},
				ID:                  "",
				OemID:               "",
				Type:                "Sata-HDD",
				StorageController:   "",
				BusInfo:             "scsi@4:0.9.0",
				WWN:                 "",
				Protocol:            "",
				SmartStatus:         "ok",
				SmartErrors:         nil,
				CapacityBytes:       8001563222016,
				BlockSizeBytes:      0,
				CapableSpeedGbps:    0,
				NegotiatedSpeedGbps: 0,
			},
			{
				Common: common.Common{
					Oem:         false,
					Description: "ATA Disk",
					Vendor:      "hgst",
					Model:       "HGST HUS728T8TALE6L4",
					Serial:      "VDJJ51YK",
					ProductName: "HGST HUS728T8TAL",
					LogicalName: "/dev/sdm",
					Firmware: &common.Firmware{
						Installed:  "V8GNW460",
						Available:  "",
						SoftwareID: "",
						Previous:   nil,
					},
					Status: nil,
				},
				ID:                  "",
				OemID:               "",
				Type:                "Sata-HDD",
				StorageController:   "",
				BusInfo:             "scsi@4:0.10.0",
				WWN:                 "",
				Protocol:            "",
				SmartStatus:         "ok",
				SmartErrors:         nil,
				CapacityBytes:       8001563222016,
				BlockSizeBytes:      0,
				CapableSpeedGbps:    0,
				NegotiatedSpeedGbps: 0,
			},
			{
				Common: common.Common{
					Oem:         false,
					Description: "ATA Disk",
					Vendor:      "hgst",
					Model:       "HGST HUS728T8TALE6L4",
					Serial:      "VDJLNERD",
					ProductName: "HGST HUS728T8TAL",
					LogicalName: "/dev/sdn",
					Firmware: &common.Firmware{
						Installed:  "V8GNW460",
						Available:  "",
						SoftwareID: "",
						Previous:   nil,
					},
					Status: nil,
				},
				ID:                  "",
				OemID:               "",
				Type:                "Sata-HDD",
				StorageController:   "",
				BusInfo:             "scsi@4:0.11.0",
				WWN:                 "",
				Protocol:            "",
				SmartStatus:         "ok",
				SmartErrors:         nil,
				CapacityBytes:       8001563222016,
				BlockSizeBytes:      0,
				CapableSpeedGbps:    0,
				NegotiatedSpeedGbps: 0,
			},
			{
				Common: common.Common{
					Oem:         false,
					Description: "ATA Disk",
					Vendor:      "hgst",
					Model:       "HGST HUS728T8TALE6L4",
					Serial:      "VDJKG9LK",
					ProductName: "HGST HUS728T8TAL",
					LogicalName: "/dev/sdc",
					Firmware: &common.Firmware{
						Installed:  "V8GNW460",
						Available:  "",
						SoftwareID: "",
						Previous:   nil,
					},
					Status: nil,
				},
				ID:                  "",
				OemID:               "",
				Type:                "Sata-HDD",
				StorageController:   "",
				BusInfo:             "scsi@4:0.0.0",
				WWN:                 "",
				Protocol:            "",
				SmartStatus:         "ok",
				SmartErrors:         nil,
				CapacityBytes:       8001563222016,
				BlockSizeBytes:      0,
				CapableSpeedGbps:    0,
				NegotiatedSpeedGbps: 0,
			},
			{
				Common: common.Common{
					Oem:         false,
					Description: "ATA Disk",
					Vendor:      "hgst",
					Model:       "HGST HUS728T8TALE6L4",
					Serial:      "VDJBVPXD",
					ProductName: "HGST HUS728T8TAL",
					LogicalName: "/dev/sdd",
					Firmware: &common.Firmware{
						Installed:  "V8GNW460",
						Available:  "",
						SoftwareID: "",
						Previous:   nil,
					},
					Status: nil,
				},
				ID:                  "",
				OemID:               "",
				Type:                "Sata-HDD",
				StorageController:   "",
				BusInfo:             "scsi@4:0.1.0",
				WWN:                 "",
				Protocol:            "",
				SmartStatus:         "ok",
				SmartErrors:         nil,
				CapacityBytes:       8001563222016,
				BlockSizeBytes:      0,
				CapableSpeedGbps:    0,
				NegotiatedSpeedGbps: 0,
			},
			{
				Common: common.Common{
					Oem:         false,
					Description: "ATA Disk",
					Vendor:      "hgst",
					Model:       "HGST HUS728T8TALE6L4",
					Serial:      "VDJJ7T3K",
					ProductName: "HGST HUS728T8TAL",
					LogicalName: "/dev/sde",
					Firmware: &common.Firmware{
						Installed:  "V8GNW460",
						Available:  "",
						SoftwareID: "",
						Previous:   nil,
					},
					Status: nil,
				},
				ID:                  "",
				OemID:               "",
				Type:                "Sata-HDD",
				StorageController:   "",
				BusInfo:             "scsi@4:0.2.0",
				WWN:                 "",
				Protocol:            "",
				SmartStatus:         "ok",
				SmartErrors:         nil,
				CapacityBytes:       8001563222016,
				BlockSizeBytes:      0,
				CapableSpeedGbps:    0,
				NegotiatedSpeedGbps: 0,
			},
			{
				Common: common.Common{
					Oem:         false,
					Description: "ATA Disk",
					Vendor:      "hgst",
					Model:       "HGST HUS728T8TALE6L4",
					Serial:      "VDJL8U5K",
					ProductName: "HGST HUS728T8TAL",
					LogicalName: "/dev/sdf",
					Firmware: &common.Firmware{
						Installed:  "V8GNW460",
						Available:  "",
						SoftwareID: "",
						Previous:   nil,
					},
					Status: nil,
				},
				ID:                  "",
				OemID:               "",
				Type:                "Sata-HDD",
				StorageController:   "",
				BusInfo:             "scsi@4:0.3.0",
				WWN:                 "",
				Protocol:            "",
				SmartStatus:         "ok",
				SmartErrors:         nil,
				CapacityBytes:       8001563222016,
				BlockSizeBytes:      0,
				CapableSpeedGbps:    0,
				NegotiatedSpeedGbps: 0,
			},
			{
				Common: common.Common{
					Oem:         false,
					Description: "ATA Disk",
					Vendor:      "hgst",
					Model:       "HGST HUS728T8TALE6L4",
					Serial:      "VDJBP7SD",
					ProductName: "HGST HUS728T8TAL",
					LogicalName: "/dev/sdg",
					Firmware: &common.Firmware{
						Installed:  "V8GNW460",
						Available:  "",
						SoftwareID: "",
						Previous:   nil,
					},
					Status: nil,
				},
				ID:                  "",
				OemID:               "",
				Type:                "Sata-HDD",
				StorageController:   "",
				BusInfo:             "scsi@4:0.4.0",
				WWN:                 "",
				Protocol:            "",
				SmartStatus:         "ok",
				SmartErrors:         nil,
				CapacityBytes:       8001563222016,
				BlockSizeBytes:      0,
				CapableSpeedGbps:    0,
				NegotiatedSpeedGbps: 0,
			},
			{
				Common: common.Common{
					Oem:         false,
					Description: "ATA Disk",
					Vendor:      "hgst",
					Model:       "HGST HUS728T8TALE6L4",
					Serial:      "VDJKVBKK",
					ProductName: "HGST HUS728T8TAL",
					LogicalName: "/dev/sdh",
					Firmware: &common.Firmware{
						Installed:  "V8GNW460",
						Available:  "",
						SoftwareID: "",
						Previous:   nil,
					},
					Status: nil,
				},
				ID:                  "",
				OemID:               "",
				Type:                "Sata-HDD",
				StorageController:   "",
				BusInfo:             "scsi@4:0.5.0",
				WWN:                 "",
				Protocol:            "",
				SmartStatus:         "ok",
				SmartErrors:         nil,
				CapacityBytes:       8001563222016,
				BlockSizeBytes:      0,
				CapableSpeedGbps:    0,
				NegotiatedSpeedGbps: 0,
			},
			{
				Common: common.Common{
					Oem:         false,
					Description: "ATA Disk",
					Vendor:      "hgst",
					Model:       "HGST HUS728T8TALE6L4",
					Serial:      "VDJK7E2K",
					ProductName: "HGST HUS728T8TAL",
					LogicalName: "/dev/sdi",
					Firmware: &common.Firmware{
						Installed:  "V8GNW460",
						Available:  "",
						SoftwareID: "",
						Previous:   nil,
					},
					Status: nil,
				},
				ID:                  "",
				OemID:               "",
				Type:                "Sata-HDD",
				StorageController:   "",
				BusInfo:             "scsi@4:0.6.0",
				WWN:                 "",
				Protocol:            "",
				SmartStatus:         "ok",
				SmartErrors:         nil,
				CapacityBytes:       8001563222016,
				BlockSizeBytes:      0,
				CapableSpeedGbps:    0,
				NegotiatedSpeedGbps: 0,
			},
			{
				Common: common.Common{
					Oem:         false,
					Description: "ATA Disk",
					Vendor:      "hgst",
					Model:       "HGST HUS728T8TALE6L4",
					Serial:      "VDJL2NUK",
					ProductName: "HGST HUS728T8TAL",
					LogicalName: "/dev/sdj",
					Firmware: &common.Firmware{
						Installed:  "V8GNW460",
						Available:  "",
						SoftwareID: "",
						Previous:   nil,
					},
					Status: nil,
				},
				ID:                  "",
				OemID:               "",
				Type:                "Sata-HDD",
				StorageController:   "",
				BusInfo:             "scsi@4:0.7.0",
				WWN:                 "",
				Protocol:            "",
				SmartStatus:         "ok",
				SmartErrors:         nil,
				CapacityBytes:       8001563222016,
				BlockSizeBytes:      0,
				CapableSpeedGbps:    0,
				NegotiatedSpeedGbps: 0,
			},
		},
		StorageControllers: []*common.StorageController{
			{
				Common: common.Common{
					Oem:          false,
					Description:  "Serial Attached SCSI controller",
					Vendor:       "Broadcom / LSI",
					Model:        "SAS3008 PCI-Express Fusion-MPT SAS-3",
					Serial:       "dead:beef",
					ProductName:  "SAS3008 PCI-Express Fusion-MPT SAS-3",
					Firmware:     nil,
					Status:       nil,
					PCIVendorID:  "dead",
					PCIProductID: "beef",
				},
				ID:                           "",
				SupportedControllerProtocols: "",
				SupportedDeviceProtocols:     "SAS",
				SupportedRAIDTypes:           "",
				PhysicalID:                   "0",
				BusInfo:                      "pci@0000:3d:00.0",
				SpeedGbps:                    0,
			},
		},
		PSUs: []*common.PSU{
			{
				Common: common.Common{
					Description: "PWS-1K23A-1R",
					Vendor:      "SUPERMICRO",
					Model:       "PWS-1K23A-1R",
					Serial:      "P1K2ACK07HB0670",
					ProductName: "PWS-1K23A-1R",
				},
				ID:                 "1",
				PowerCapacityWatts: 1200,
			},
			{
				Common: common.Common{
					Description: "PWS-1K23A-1R",
					Vendor:      "SUPERMICRO",
					Model:       "PWS-1K23A-1R",
					Serial:      "P1K2ACK07HB0669",
					ProductName: "PWS-1K23A-1R",
				},
				ID:                 "2",
				PowerCapacityWatts: 1200,
			},
		},
		Enclosures: []*common.Enclosure{},
	}
)
