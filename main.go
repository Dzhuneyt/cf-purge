package main

import (
	"context"
	"fmt"
	"log"
	"os"
	"path/filepath"

	"github.com/Dzhuneyt/cf-purge/utils"
	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/service/cloudformation"
	"github.com/urfave/cli/v2"
)

// @TODO Cleanup this file and split it into multiple files

func deleteStacks(client *cloudformation.Client, stacks []string) {
	for _, stack := range stacks {
		fmt.Printf("Deleting stack %s\n", stack)
		_, err := client.DeleteStack(context.TODO(), &cloudformation.DeleteStackInput{
			StackName: &stack,
		})
		if err != nil {
			log.Fatal(err)
		}

		// @TODO STEP 1: Wait for all deletions to start
		// @TODO STEP 2: Poll for the status of the deletions
		// @TODO STEP 3: If stack deletion fails because of a dependency (another stack), retry step 1 and 2 after a delay
	}
}

func newCloudFormationClient() (*cloudformation.Client, error) {
	cfg, err := config.LoadDefaultConfig(context.TODO())
	if err != nil {
		return nil, fmt.Errorf("failed to load AWS configuration: %w", err)
	}

	// Check if AWS credentials are available
	if _, err = cfg.Credentials.Retrieve(context.TODO()); err != nil {
		return nil, fmt.Errorf("failed to retrieve AWS credentials: %w", err)
	}

	return cloudformation.NewFromConfig(cfg), nil
}

var blacklistedStatuses = []string{
	"CREATE_IN_PROGRESS",
	"DELETE_IN_PROGRESS",
	"DELETE_COMPLETE",
	// TODO add more statuses if needed
}

func filterStacks(client *cloudformation.Client, pattern string) []string {
	paginator := cloudformation.NewListStacksPaginator(client, &cloudformation.ListStacksInput{})
	stacks := make([]string, 0)

	for paginator.HasMorePages() {
		output, err := paginator.NextPage(context.TODO())
		if err != nil {
			log.Fatal(err)
			return nil
		}
		for _, stack := range output.StackSummaries {
			stackName := *stack.StackName

			// Check if the stack name matches the glob pattern
			matched, _ := filepath.Match(pattern, stackName)
			if !matched {
				continue
			}

			// Check if the stack is in any of the denylisted status
			isBlackListed := false
			for _, status := range blacklistedStatuses {
				if status == string(stack.StackStatus) {
					fmt.Printf("Skipping stack %s with status %s\n", stackName, stack.StackStatus)
					isBlackListed = true
					break
				}
			}

			if isBlackListed {
				continue
			}

			stacks = append(stacks, stackName)
		}
	}

	return stacks
}

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
	cloudformationClient, err := newCloudFormationClient()
	if err != nil {
		log.Fatal(err)
	}

	stacks := filterStacks(cloudformationClient, globalPattern)

	isConfirmed := confirmStackDeletion(stacks)
	if !isConfirmed {
		log.Fatal("Operation cancelled")
		return
	}

	deleteStacks(cloudformationClient, stacks)
}

func confirmStackDeletion(stacks []string) bool {
	fmt.Println("Glob pattern matched the following stacks:")
	fmt.Println("----------------------------------------")
	for _, stack := range stacks {
		fmt.Println(stack)
	}
	fmt.Println("----------------------------------------")

	fmt.Printf("This operation will delete %d stacks\n", len(stacks))
	fmt.Println("THIS OPERATION IS DESTRUCTIVE AND IRREVERSIBLE")
	fmt.Println("Please, confirm that you want to delete these stacks irreversibly.")

	return utils.AskForConfirmation()
}
