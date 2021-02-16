package log

func applyOptions(opts []Option) *options {
	var o options
	o.logger = logger
	for _, opt := range opts {
		opt(&o)
	}
	return &o
}

// Setup returns a new logrus instance.
func Setup(opts ...Option) error {
	conf := applyOptions(opts)
	conf.logger.SetFormatter(conf.logger.Formatter)
	conf.logger.SetLevel(conf.logger.Level)
	return nil
}

func init() {
	logger = newLogger()
	logger.SetReportCaller(false)
}
