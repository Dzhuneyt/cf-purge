package main

import (
	"fmt"
	"log"
	"os"

	"github.com/Dzhuneyt/cf-purge/internal/aws"
	"github.com/Dzhuneyt/cf-purge/internal/prompt"
	"github.com/Dzhuneyt/cf-purge/internal/stack"
	"github.com/urfave/cli/v2"
)

func main() {
	app := &cli.App{
		Name:  "cf-purge",
		Usage: "Purge one more CloudFormation stacks that match a pattern",
		Action: func(cCtx *cli.Context) error {
			glob := cCtx.String("glob")
			fmt.Println("glob:", glob)

			purgeStacks(glob)
			return nil
		},
		Flags: []cli.Flag{
			&cli.StringFlag{
				Name:     "glob",
				Value:    "",
				Usage:    "a glob pattern that matches a list of CloudFormation stacks, or a specific stack name",
				Required: true,
			},
		},
	}
	_ = app.Run(os.Args)
}

func purgeStacks(globalPattern string) {
	client, err := aws.NewCloudFormationClient()
	if err != nil {
		log.Fatal(err)
	}

	stacks := stack.FilterStacks(client, globalPattern)

	if !prompt.ConfirmStackDeletion(stacks) {
		log.Fatal("Operation cancelled")
	}

	stack.DeleteStacks(client, stacks)
}
