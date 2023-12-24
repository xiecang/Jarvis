package conf

// Config interface determines the common methods for parsing configuration from specified resources
type Config interface {
	Parse(any) error
	Get(string) any
	GetBool(string) bool
	GetString(string) string
	GetInt(string) int
}
