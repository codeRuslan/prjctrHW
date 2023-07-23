package mainProggram

import (
	"fmt"
	fileOrganizer "github.com/userName/otherModule"
	"regexp"
)

func main() {
	createFoldersStrategy := &fileOrganizer.CreateFolderStrategy{}
	deleteFilesStrategy := &fileOrganizer.DeleteFilesStrategy{}
	deleteFilesWithUnderscoreStrategy := &fileOrganizer.RegexDecorator{
		RegexPattern: regexp.MustCompile("_"),
		Strategy:     deleteFilesStrategy,
	}

	createFolderWithNameStrategy := &fileOrganizer.RegexDecorator{
		RegexPattern: regexp.MustCompile("folder\\d+"),
		Strategy:     createFoldersStrategy,
	}

	strategies := []fileOrganizer.FileOperationStrategy{
		createFoldersStrategy,
		deleteFilesStrategy,
		deleteFilesWithUnderscoreStrategy,
		createFolderWithNameStrategy,
	}

	for _, strategy := range strategies {
		// Replace "actionPath" and "pattern" with your desired values
		actionPath := "/path/to/your/target/directory"
		err := strategy.Execute(actionPath)
		if err != nil {
			fmt.Println("Error executing strategy:", err)
		}
	}
}
