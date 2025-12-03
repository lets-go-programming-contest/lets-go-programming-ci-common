package wifi

import (
	"net"

	"github.com/mdlayher/wifi"
)

type WiFiHandle interface {
	Interfaces() ([]*wifi.Interface, error)
}

type WiFiService struct {
	WiFi WiFiHandle
}

func New(wifi WiFiHandle) WiFiService {
	return WiFiService{WiFi: wifi}
}

func (service WiFiService) GetAddresses() ([]net.HardwareAddr, error) {
	interfaces, err := service.WiFi.Interfaces()
	if err != nil {
		return nil, err
	}

	var addrs []net.HardwareAddr

	for _, iface := range interfaces {
		addrs = append(addrs, iface.HardwareAddr)
	}

	return addrs, nil
}

func (service WiFiService) GetNames() ([]string, error) {
	interfaces, err := service.WiFi.Interfaces()
	if err != nil {
		return nil, err
	}

	var name_list []string

	for _, iface := range interfaces {
		name_list = append(name_list, iface.Name)
	}

	return name_list, nil
}
