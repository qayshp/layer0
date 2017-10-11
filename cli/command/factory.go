package command

import (
	"github.com/quintilesims/layer0/cli/printer"
	"github.com/quintilesims/layer0/client"
)

type CommandFactory struct {
	client   client.Client
	printer  printer.Printer
	resolver Resolver
}

func NewCommandFactory(c client.Client, p printer.Printer, r Resolver) *CommandFactory {
	return &CommandFactory{
		client:   c,
		printer:  p,
		resolver: r,
	}
}

func (f *CommandFactory) SetClient(c client.Client) {
	f.client = c
}

func (f *CommandFactory) SetPrinter(p printer.Printer) {
	f.printer = p
}

func (f *CommandFactory) SetResolver(r Resolver) {
	f.resolver = r
}