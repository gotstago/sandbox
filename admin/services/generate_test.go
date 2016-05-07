package services

import (
	"fmt"
	"os"
	"path/filepath"
	"testing"
	// "strings"
)

func TestGenerateFromTemplate(t *testing.T) {
    //create entities
    //get list of entities that are tagged as resources
	packageName := "models"
	goPath, err := GetGopath()
	if err != nil {
		t.Fatal("should be a GOPATH")
	}
	projectName := "myProject"

	//targetRoot := goPath + "/src/github.com/gotstago/qor-projects/" + projectName
	// pkgDir, err := packageDir(packageName)
	// if err != nil {
	// 	panic(err)
	// }
	targetRoot := goPath + "/src/github.com/gotstago/qor-projects/" + projectName
	targetController := targetRoot + "/app/controllers"
	targetModels := targetRoot + "/app/models"
	targetViews := targetRoot + "/app/views"
	targetAdmin := targetRoot + "/config/admin"
	targetRoutes := targetRoot + "/config/routes"
	targetDb := targetRoot + "/db"
	targetDist := targetRoot + "/public/dist"
	targetFonts := targetRoot + "/public/fonts"
	targetImages := targetRoot + "/public/images"
	targetJavascripts := targetRoot + "/public/javascripts"
	targetStylesheets := targetRoot + "/public/stylesheets"
	targetVendors := targetRoot + "/public/vendors"
	err = CreateProjectFolders(
		targetRoot,
		targetController,
		targetModels,
		targetViews,
		targetAdmin,
		targetRoutes,
		targetDb,
		targetDist,
		targetFonts,
		targetImages,
		targetJavascripts,
		targetStylesheets,
		targetVendors,
	)
	if err != nil {
		t.Fatal(err)
	}

	typeName := "Product"
	typeCategory := "entity"

	typeData := map[string]string{
		"Name":        "string",
		"Description": "string",
	}

	outputFile := FormatFileName(typeName, typeCategory)
	writer, err := os.Create(filepath.Join(targetModels, outputFile))
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
