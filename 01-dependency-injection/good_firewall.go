/*
Kode ini bertumpu pada kontrak Auditor SecurityAuditor
*sehingga Firewall tidak peduli apapun yang di gunakan asalkan
punya method LogDroppedPacket
*ini mematuhi Dependency Inversion Principle
*tidak perlu menyentuh struct Firewall sama sekali untuk mengganti logger
*/
package main

import "fmt"

type SecurityAuditor interface {
	LogDroppedPacket(ipAddress string)
}

type FileLogger struct{}

func (f FileLogger) LogDroppedPacket(ipAddress string) {
	fmt.Println("System Log: Koneksi ditolak dari IP:", ipAddress)
}

type MockLogger struct {
	BlockedHistory []string
}

func (m *MockLogger) LogDroppedPacket(ipAddress string) {
	m.BlockedHistory = append(m.BlockedHistory, ipAddress)
}

type Firewall struct {
	Auditor SecurityAuditor
}

func (fw Firewall) BlockTraffic(ipAddress string) {
	fmt.Println("Firewall mendeteksi trafik anomali...")
	fw.Auditor.LogDroppedPacket(ipAddress)
}

func main() {
	fmt.Println("--Good Firewall (Production)--")

	prodLogger := FileLogger{}
	prodFirewall := Firewall{
		Auditor: prodLogger,
	}
	prodFirewall.BlockTraffic("10.10.2.65")

	fmt.Println("--Good Firewall (Testing)--")

	// testing:
	testLogger := &MockLogger{}
	testFirewall := Firewall{
		Auditor: testLogger,
	}

	testFirewall.BlockTraffic("192.168.1.99")
	testFirewall.BlockTraffic("172.16.10.21")

	fmt.Println("Hasil Audit di dalam memori: ", testLogger.BlockedHistory)
}
