package main

import (
	"archive/zip"
	"io"
	"os"
	"path"
)

type BundleConfig struct {
	Modules []string `yaml:"modules"`
}

func createBundle(config *BundleConfig, modulesDir string, outputPath string) error {
	// Create the output file
	file, err := os.Create(outputPath)
	if err != nil {
		return err
	}
	defer file.Close()
	//
	zipWriter := zip.NewWriter(file)
	defer zipWriter.Close()

	// Add the modules to the zip file
	for _, module := range config.Modules {
		zipFolder(zipWriter, path.Join(modulesDir, module))
	}
	return nil
}

func zipFolder(zipWriter *zip.Writer, folderPath string) error {
	// Open the folder
	folder, err := os.Open(folderPath)
	if err != nil {
		return err
	}
	defer folder.Close()
	// Get the list of files in the folder
	files, err := folder.Readdir(-1)
	if err != nil {
		return err
	}
	// Add each file to the zip archive
	for _, file := range files {
		if file.IsDir() {
			// Recursively add subfolders
			err = zipFolder(zipWriter, path.Join(folderPath, file.Name()))
			if err != nil {
				return err
			}
		} else {
			// Add the file to the zip archive
			err = zipFile(zipWriter, path.Join(folderPath, file.Name()))
			if err != nil {
				return err
			}
		}
	}
	return nil
}

func zipFile(zipWriter *zip.Writer, filePath string) error {
	// Open the file
	file, err := os.Open(filePath)
	if err != nil {
		return err
	}
	defer file.Close()
	// Create a new file in the zip archive
	zipFile, err := zipWriter.Create(filePath)
	if err != nil {
		return err
	}
	// Copy the file to the zip file
	_, err = io.Copy(zipFile, file)
	if err != nil {
		return err
	}
	return nil
}
