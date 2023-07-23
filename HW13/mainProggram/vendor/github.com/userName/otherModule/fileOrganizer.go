package fileOrganizer

import (
	"fmt"
	"os"
	"path/filepath"
	"regexp"
)

type FileOperationStrategy interface {
	Execute(path string) error
}

type CreateFolderStrategy struct{}

func (s *CreateFolderStrategy) Execute(path string) error {
	return os.MkdirAll(path, os.ModePerm)
}

type DeleteFilesStrategy struct{}

func (s *DeleteFilesStrategy) Execute(path string) error {
	err := filepath.Walk(path, func(p string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}
		if !info.IsDir() {
			return os.Remove(p)
		}
		return nil
	})
	if err != nil {
		fmt.Println("Error deleting files:", err)
	}
	return err
}

type FileOperationDecorator interface {
	FileOperationStrategy
}

type RegexDecorator struct {
	RegexPattern *regexp.Regexp
	Strategy     FileOperationStrategy
}

func (rd *RegexDecorator) Execute(path string) error {
	matchedPaths := []string{}

	err := filepath.Walk(path, func(p string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}
		if !info.IsDir() && rd.RegexPattern.MatchString(info.Name()) {
			matchedPaths = append(matchedPaths, p)
		}
		return nil
	})
	if err != nil {
		fmt.Println("Error walking the path:", err)
		return err
	}

	for _, matchedPath := range matchedPaths {
		if err := rd.Strategy.Execute(matchedPath); err != nil {
			fmt.Println("Error executing Strategy:", err)
		}
	}

	return nil
}
