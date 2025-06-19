package prompt

import (
	"fmt"

	"github.com/Dzhuneyt/cf-purge/utils"
)

func ConfirmStackDeletion(stacks []string) bool {
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
