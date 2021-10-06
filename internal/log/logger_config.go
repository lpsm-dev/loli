package log

// Config defines all the configurable options for the logrus logger.
type Config struct {
	Level   string
	Format  string
	Output  string
	File    string
	Verbose bool
}

// SetDefault set default values for logrus logger configurable options.
func (config *Config) SetDefault(level, formater, output, file string, verbose bool) {
	config.Level = level
	config.Format = formater
	config.Output = output
	config.File = file
	config.Verbose = verbose
}
