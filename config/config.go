package config
import "fmt"
import "net"

type Config struct {
	DB *DBConfig
}

type DBConfig struct {
	Dialect  string
	Host     string
	Port     int
	Username string
	Password string
	Name     string
	Charset  string
}

func GetConfig() (*Config, error) {
	ifaces, err := net.Interfaces()
	if err != nil {
		return nil, err
	}

	// Find the first non-loopback interface with a valid IP address
	var hostIP string
	for _, iface := range ifaces {
		if iface.Flags&net.FlagLoopback != 0 {
			continue
		}

		addrs, err := iface.Addrs()
		if err != nil {
			return nil, err
		}

		for _, addr := range addrs {
			ip, _, err := net.ParseCIDR(addr.String())
			if err != nil {
				continue
			}

			if ip.To4() != nil {
				hostIP = ip.String()
				break
			}
		}

		if hostIP != "" {
			break
		}
	}

	if hostIP == "" {
		return nil, fmt.Errorf("unable to determine host IP address")
	}

	config := &Config{
		DB: &DBConfig{
			Dialect:  "mysql",
			Host:     hostIP,
			Port:     3306,
			Username: "guest",
			Password: "Guest0000!",
			Name:     "todoapp",
			Charset:  "utf8",
		},
	}

	return config, nil
}