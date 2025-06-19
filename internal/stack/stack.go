package stack

import (
	"context"
	"fmt"
	"log"
	"path/filepath"

	"github.com/aws/aws-sdk-go-v2/service/cloudformation"
	"github.com/aws/aws-sdk-go-v2/service/cloudformation/types"
)

var blacklistedStatuses = []string{
	"CREATE_IN_PROGRESS",
	"DELETE_IN_PROGRESS",
	"DELETE_COMPLETE",
	// TODO add more statuses if needed
}

func FilterStacks(client *cloudformation.Client, pattern string) []string {
	paginator := cloudformation.NewListStacksPaginator(client, &cloudformation.ListStacksInput{})
	stacks := make([]string, 0)

	for paginator.HasMorePages() {
		output, err := paginator.NextPage(context.TODO())
		if err != nil {
			log.Fatalf("Failed to list stacks: %v", err)
		}

		for _, stack := range output.StackSummaries {
			stackName := *stack.StackName

			// Check if the stack name matches the glob pattern
			matched, _ := filepath.Match(pattern, stackName)
			if !matched {
				continue
			}

			// Check if the stack is in any of the denylisted status
			if isBlackListed(stack) {
				fmt.Printf("Skipping stack %s with status %s\n", stackName, stack.StackStatus)
				continue
			}

			stacks = append(stacks, stackName)
		}
	}

	return stacks
}

func isBlackListed(stack types.StackSummary) bool {
	for _, status := range blacklistedStatuses {
		if status == string(stack.StackStatus) {
			return true
		}
	}
	return false
}

func DeleteStacks(client *cloudformation.Client, stacks []string) {
	for _, stack := range stacks {
		fmt.Printf("Deleting stack %s\n", stack)
		_, err := client.DeleteStack(context.TODO(), &cloudformation.DeleteStackInput{
			StackName: &stack,
		})
		if err != nil {
			log.Fatalf("Failed to delete stack %s: %v", stack, err)
		}

		// @TODO STEP 1: Wait for all deletions to start
		// @TODO STEP 2: Poll for the status of the deletions
		// @TODO STEP 3: If stack deletion fails because of a dependency (another stack), retry step 1 and 2 after a delay
	}
}
