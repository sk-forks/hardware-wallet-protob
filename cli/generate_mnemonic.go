package cli

import (
	gcli "github.com/urfave/cli"

	deviceWallet "github.com/skycoin/hardware-wallet-go/device-wallet"
)

func generateMnemonicCmd() gcli.Command {
	name := "generateMnemonic"
	return gcli.Command{
		Name:        name,
		Usage:       "Ask the device to generate a mnemonic and configure itself with it.",
		Description: "",
		Flags: []gcli.Flag{
			gcli.BoolFlag{
				Name:  "usePassphrase",
				Usage: "Configure a passphrase",
			},
			gcli.IntFlag{
				Name:  "wordCount",
				Usage: "Use a specific (12 | 24) number of words for the Mnemonic",
				Value: 12,
			},
			gcli.StringFlag{
				Name:   "deviceType",
				Usage:  "Device type to send instructions to, hardware wallet (USB) or emulator.",
				EnvVar: "DEVICE_TYPE",
			},
		},
		OnUsageError: onCommandUsageError(name),
		Action: func(c *gcli.Context) {
			usePassphrase := c.Bool("usePassphrase")
			wordCount := uint32(c.Uint64("wordCount"))

			var deviceType deviceWallet.DeviceType
			switch c.String("deviceType") {
			case "USB":
				deviceType = deviceWallet.DeviceTypeUsb
			case "EMULATOR":
				deviceType = deviceWallet.DeviceTypeEmulator
			default:
				log.Error("device type not set")
				return
			}

			deviceWallet.DeviceGenerateMnemonic(deviceType, wordCount, usePassphrase)
		},
	}
}
