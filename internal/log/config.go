package log

// Config defines all the configurable options for the logger.
type Config struct {
	Level   string
	Format  string
	Output  string
	File    string
	Silence bool
}

// SetDefault set sane default for logger's config.
func (conf *Config) SetDefault(level, formater, output, file string, silence bool) {
	conf.Level = level
	conf.Format = formater
	conf.Output = output
	conf.File = file
	conf.Silence = silence
}
