package config

// Config represents the application configuration
type Config struct {
	Timeout    int
	MaxRetries int
	Headers    map[string]string
}

// NewConfig creates a new configuration with default values
// Returns:
// - *Config: A new configuration instance with default values
func NewConfig() *Config {
	return &Config{
		Timeout:    30,
		MaxRetries: 3,
		Headers:    map[string]string{},
	}
}

// SetTimeout sets the request timeout
// Parameters:
// - timeout: int - Timeout in seconds
func (c *Config) SetTimeout(timeout int) {
	c.Timeout = timeout
}

// SetMaxRetries sets the maximum number of retries
// Parameters:
// - retries: int - Maximum number of retries
func (c *Config) SetMaxRetries(retries int) {
	c.MaxRetries = retries
}

// AddHeader adds a header to the configuration
// Parameters:
// - key: string - Header key
// - value: string - Header value
func (c *Config) AddHeader(key, value string) {
	c.Headers[key] = value
}
