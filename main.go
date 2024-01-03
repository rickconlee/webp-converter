package main

import (
    "fmt"
    "image"
    // Import with underscore to register formats without direct usage
    _ "image/jpeg"
    _ "image/png"
    "os"
    "path/filepath"
    "strings"

    _ "image/gif" // Register GIF format
    "github.com/chai2010/webp"
)

func main() {
    if len(os.Args) < 2 {
        fmt.Println("Usage: go run main.go <directory_path>")
        os.Exit(1)
    }

    directoryPath := os.Args[1]
    err := convertDirectoryToWebP(directoryPath)
    if err != nil {
        fmt.Printf("Error converting images: %v\n", err)
        os.Exit(1)
    }

    fmt.Println("All conversions successful!")
}

func convertDirectoryToWebP(dirPath string) error {
    return filepath.Walk(dirPath, func(path string, info os.FileInfo, err error) error {
        if err != nil {
            return err
        }
        if !info.IsDir() && isSupportedFormat(path) {
            err := convertToWebP(path)
            if err != nil {
                fmt.Printf("Failed to convert %s: %v\n", path, err)
            } else {
                fmt.Printf("Converted %s successfully\n", path)
            }
        }
        return nil
    })
}

func isSupportedFormat(filePath string) bool {
    ext := strings.ToLower(filepath.Ext(filePath))
    return ext == ".jpg" || ext == ".jpeg" || ext == ".png" || ext == ".gif"
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
