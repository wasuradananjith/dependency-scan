package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"path/filepath"

	jsoniter "github.com/json-iterator/go"
	"github.com/xuri/excelize/v2"
)

// APIDefinition represents an API artifact in APIM
type PackageJson struct {
	Dependencies map[string]string `json:"dependencies,omitempty"`
}

func main() {
	f := excelize.NewFile()
	columnA := "A"
	columnB := "B"
	columnC := "C"
	cell := 1

	f.SetCellValue("Sheet1", fmt.Sprintf("%s%d", columnA, cell), "Dependency")
	f.SetCellValue("Sheet1", fmt.Sprintf("%s%d", columnB, cell), "Version")
	f.SetCellValue("Sheet1", fmt.Sprintf("%s%d", columnC, cell), "FileName")

	cell++

	filePath := filepath.Join("/home/wasura/Desktop/security-scan/files")

	files, err := ioutil.ReadDir(filePath)
	if err != nil {
		fmt.Println(err)
	}

	for _, file := range files {
		fileName := file.Name()

		content, err := ioutil.ReadFile(filePath + "/" + fileName)
		if err != nil {
			fmt.Println(err)
		}

		packageJson := PackageJson{}

		// Read from yaml definition
		err = jsoniter.Unmarshal(content, &packageJson)
		if err != nil {
			fmt.Println(err)
		}

		for key, version := range packageJson.Dependencies {
			f.SetCellValue("Sheet1", fmt.Sprintf("%s%d", columnA, cell), key)
			f.SetCellValue("Sheet1", fmt.Sprintf("%s%d", columnB, cell), version)
			f.SetCellValue("Sheet1", fmt.Sprintf("%s%d", columnC, cell), fileName)
			cell++
		}

	}

	// Save spreadsheet by the given path.
	if err := f.SaveAs("Book1.xlsx"); err != nil {
		fmt.Println(err)
	}

	//fmt.Println(packageJson)
}

func callGitHubAPI() {
	fmt.Println("Calling API...")
	client := &http.Client{}
	req, err := http.NewRequest("GET", "https://api.github.com/search/code", nil)
	if err != nil {
		fmt.Print(err.Error())
	}
	req.Header.Add("Accept", "application/json")
	req.Header.Add("Content-Type", "application/json")
	resp, err := client.Do(req)
	if err != nil {
		fmt.Print(err.Error())
	}
	defer resp.Body.Close()
	bodyBytes, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Print(err.Error())
	}
	var responseObject interface{}
	json.Unmarshal(bodyBytes, &responseObject)
	fmt.Printf("API Response as struct %+v\n", responseObject)
}
