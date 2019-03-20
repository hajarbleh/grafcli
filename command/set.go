package command

import (
	"fmt"
	"github.com/hajarbleh/grafcli/config"
	"github.com/pkg/errors"
	"github.com/urfave/cli"
)

type Set struct {
}

func (s *Set) Execute(ctx *cli.Context) error {
	fmt.Println("Example executionsss")
	_, err := config.Read()
	fmt.Println("Example executionsss")
	if err != nil {
		fmt.Println(err)
		return errors.Wrap(err, "Error reading config file")
	}
	//key := ctx.Args()[0]
	//value := ctx.Args()[1]

	return nil
}

func (s *Set) Flags() []cli.Flag {
	return []cli.Flag{}
}
