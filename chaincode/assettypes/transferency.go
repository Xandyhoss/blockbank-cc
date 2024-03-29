package assettypes

import "github.com/hyperledger-labs/cc-tools/assets"

var Transferency = assets.AssetType{
	Tag:         "transferency",
	Label:       "Transferency",
	Description: "Transferency",

	Props: []assets.AssetProp{
		{
			Required: true,
			IsKey:    true,
			Label:    "tdId",
			DataType: "string",
			Tag:      "txId",
		},
		{
			Required: true,
			IsKey:    true,
			Tag:      "sender",
			Label:    "Sender",
			DataType: "->holder",
			Writers:  []string{"orgMSP"},
		},
		{
			Required: true,
			IsKey:    true,
			Tag:      "receiver",
			Label:    "Receiver",
			DataType: "->holder",
			Writers:  []string{"orgMSP"},
		},
		{
			Required: true,
			Tag:      "value",
			Label:    "Value",
			DataType: "number",
		},
		{
			Tag:      "date",
			Label:    "Transferency date",
			DataType: "datetime",
			Required: true,
		},
	},
}
