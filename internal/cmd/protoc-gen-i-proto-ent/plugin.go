package main

import (
	"flag"

	"google.golang.org/protobuf/compiler/protogen"
)

type Plugin struct {
	opts *protogen.Options

	flg            flag.FlagSet
	flgSkipPackage *bool
}

func NewPlugin() *Plugin {
	p := Plugin{}

	p.InitFlags()
	p.opts = &protogen.Options{
		ParamFunc: p.flg.Set,
	}
	return &p
}

func (p *Plugin) InitFlags() {
	p.flgSkipPackage = p.flg.Bool("skip_package", true, "if true, files are generated as '<i-proto-ent_out>/<entity>/<entity>.gen.go'")
}

func (p *Plugin) Run() {
	p.opts.Run(func(gen *protogen.Plugin) error {
		for _, f := range gen.Files {
			if !f.Generate {
				continue
			}

			for _, msg := range f.Messages {
				p.genEntity(gen, f, msg)
			}
		}
		return nil
	})
}
