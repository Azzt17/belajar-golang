/*kode ini bertumpu pada object konkret (Loggel FileLogger)
* sehingga firewall hanya bisa mengenali FileLogger
* sehingga kode ini tidak SOLID dan harus membongkar
* Firewall jika ingin mengganti cara logging
 */
package main

import "fmt"

type FileLogger struct{}

func (f FileLogger) LogDroppedPacket(ipAddress string) {
	fmt.Println("System Log: Koneksi di tolak dari IP:", ipAddress)
}

type Firewall struct {
	Logger FileLogger
}

func (fw Firewall) BlockTraffic(ipAddress string) {
	fmt.Println("Firewall mendeteksi traffic mencurigakan..")
	fw.Logger.LogDroppedPacket(ipAddress)
}

func main() {
	fmt.Println("--Testing Bad Firewall--")

	logger := FileLogger{}
	myFirewall := Firewall{
		Logger: logger,
	}
	myFirewall.BlockTraffic("192.168.0.20")
}
