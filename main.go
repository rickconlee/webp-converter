package main

import (
	"fmt"
	"image"
	_ "image/gif"
	_ "image/jpeg"
	_ "image/png"
	"os"
	"path/filepath"
	"strings"
	"time"

	_ "golang.org/x/image/tiff"

	"github.com/chai2010/webp"
	"github.com/schollz/progressbar/v3"
)

func main() {
	if len(os.Args) < 2 {
		fmt.Println("Usage: $ webpconvert <directory_path>")
		os.Exit(1)
	}

	directoryPath := os.Args[1]
	fmt.Println("Scanning directory for images...")

	// Count total images first
	totalImages, err := countImages(directoryPath)
	if err != nil {
		fmt.Printf("Error scanning directory: %v\n", err)
		os.Exit(1)
	}

	if totalImages == 0 {
		fmt.Println("No images found for conversion.")
		os.Exit(0)
	}

	fmt.Printf("Found %d images for conversion. Starting...\n", totalImages)
	bar := progressbar.Default(int64(totalImages))

	// Start the conversion process with timer
	start := time.Now()
	err = convertDirectoryToWebP(directoryPath, bar)
	if err != nil {
		fmt.Printf("Error converting images: %v\n", err)
		os.Exit(1)
	}

	elapsed := time.Since(start)
	fmt.Printf("\nAll conversions successful! Time taken: %s\n", elapsed)
}

func countImages(dirPath string) (int, error) {
	var count int
	err := filepath.Walk(dirPath, func(path string, info os.FileInfo, err error) error {
		if err == nil && !info.IsDir() && isSupportedFormat(path) {
			count++
		}
		return err
	})
	return count, err
}

func convertDirectoryToWebP(dirPath string, bar *progressbar.ProgressBar) error {
	return filepath.Walk(dirPath, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}
		if !info.IsDir() && isSupportedFormat(path) {
			err := convertToWebP(path)
			if err != nil {
				fmt.Printf("Failed to convert %s: %v\n", path, err)
			}
			bar.Add(1)
		}
		return nil
	})
}

func isSupportedFormat(filePath string) bool {
	ext := strings.ToLower(filepath.Ext(filePath))
	return ext == ".jpg" || ext == ".jpeg" || ext == ".png" || ext == ".gif" || ext == ".tiff" || ext == ".tif"
}

func convertToWebP(inputPath string) error {
	file, err := os.Open(inputPath)
	if err != nil {
		return err
	}
	defer file.Close()

	img, _, err := image.Decode(file)
	if err != nil {
		return err
	}

	outputFileName := strings.TrimSuffix(inputPath, filepath.Ext(inputPath)) + ".webp"
	outputFile, err := os.Create(outputFileName)
	if err != nil {
		return err
	}
	defer outputFile.Close()

	return webp.Encode(outputFile, img, &webp.Options{Lossless: true})
}
