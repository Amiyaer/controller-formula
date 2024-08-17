package config

type Config struct {
    NodeID string
	Version string
	// your implements
}

func Default() *Config {
	return &Config{
		NodeID: "MyNodeID",
		Version: "MyVersion",
	}
}