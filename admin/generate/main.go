package main

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"
)

// func init() {
// 	flag.Usage = func() {
// 		fmt.Fprintln(os.Stderr, "Usage of wordsmith:")
// 		fmt.Fprintln(os.Stderr, "")
// 		fmt.Fprintln(os.Stderr, "  wordsmith -pointer -type=<type> -format=json")
// 		fmt.Fprintln(os.Stderr, "")
// 		flag.PrintDefaults()
// 	}

// 	flag.CommandLine.Init("", flag.ExitOnError)
// }

func main() {
	// typePointer := flag.Bool("pointer", false, "Determines whether a type is a pointer or not")
	// typeName := flag.String("type", "", "Type that hosts io.WriterTo interface implementation")
	// packageName := flag.String("package", "", "Package name")
	// format := flag.String("format", "json", "Encoding format")

	// flag.Parse()

	// if *typeName == "" || *format != "json" {
	// 	flag.Usage()
	// 	return
	// }
	packageName := "models"

	pkgDir, err := packageDir(packageName)
	if err != nil {
		panic(err)
	}

	typeName := "Product"

	typeData := map[string]string{
		"Name":        "string",
		"Description": "string",
	}

	outputFile := formatFileName(typeName)
	writer, err := os.Create(filepath.Join(pkgDir, outputFile))
	if err != nil {
		panic(err)
	}
	defer writer.Close()

	generator := &Generator{Format: JSON}

	m := metadata(typeName, typeData, packageName)
	if err := generator.Generate(writer, m); err != nil {
		panic(err)
	}

	fmt.Printf("Generated %s\n", outputFile)
}

func formatFileName(typeName string) string {
	return fmt.Sprintf("%s_entity.go", strings.ToLower(typeName))
}

func packageDir(packageName string) (string, error) {
	if packageName == "" {
		return os.Getwd()
	}

	// path := os.Getenv("GOPATH")
	// if path == "" {
	// 	return "", errors.New("GOPATH is not set")
	// }

	currentWd, _ := os.Getwd()

	workDir := filepath.Join(currentWd, packageName)
	if _, err := os.Stat(workDir); err != nil {
		return "", err
	}

	return workDir, nil
}

func metadata(typeName string, typeData map[string]string, packageName string) (m Metadata) {
	m.TypeData = typeData
	m.TypeName = typeName
	m.PackageName = packageName

	// if pointerType {
	// 	m.MarshalObject = m.Object
	// } else {
	// 	m.MarshalObject = fmt.Sprintf("&%s", m.Object)
	// }

	return m
}
