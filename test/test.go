package test

/*
 Test Utility functions.
 */
import (
	"io/ioutil"
	"os"
	"path"
)

func createTestFilePath(fileName string) string {
	testDirPath, _ := os.Getwd()
	return path.Join(testDirPath, fileName)
}

func LoadTestJson(fileName string) string {
	filePath := createTestFilePath(fileName)
	contents, _ := ioutil.ReadFile(filePath)
	return string(contents);
}

