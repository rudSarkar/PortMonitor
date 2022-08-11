package options

import "flag"

type Options struct {
	File string
	Host string
}

func ParseOptions() *Options {

	options := &Options{}

	flag.StringVar(&options.File, "f", "", "File containing domains")
	flag.StringVar(&options.Host, "h", "", "Single host")

	flag.Parse()

	return options
}
