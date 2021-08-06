// Shows how to extract an passsword encrypted zip file using 7zip.
// By Larry Battle <https://github.com/LarryBattle>
// Answer to http://stackoverflow.com/questions/20330210/golang-1-2-unzip-password-protected-zip-file
// 7-zip.chm - http://sevenzip.sourceforge.jp/chm/cmdline/switches/index.htm
// Effective Golang - http://golang.org/doc/effective_go.html
package main

import (
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
	"strings"
)

var (
	txt_content     = "Sample file created."
	txt_filename    = "name.txt"
	zip_filename    = "sample.zip"
	zip_password    = "42"
	zip_encryptType = "AES256"
	base_path       = "./"

	test_path          = filepath.Join(base_path, "test")
	src_path           = filepath.Join(test_path, "src")
	extract_path       = filepath.Join(test_path, "extracted")
	extracted_txt_path = filepath.Join(extract_path, txt_filename)
	txt_path           = filepath.Join(src_path, txt_filename)
	zip_path           = filepath.Join(src_path, zip_filename)
)
var txt_fileSize int64

func checkError(e error) {
	if e != nil {
		panic(e)
	}
}
func setupTestDir() {
	fmt.Printf("Removing `%s`\n", test_path)
	var e error
	os.Remove(test_path)
	fmt.Printf("Creating `%s`,`%s`\n", extract_path, src_path)
	e = os.MkdirAll(src_path, os.ModeDir|os.ModePerm)
	checkError(e)
	e = os.MkdirAll(extract_path, os.ModeDir|os.ModePerm)
	checkError(e)
}
func createSampleFile() {
	fmt.Println("Creating", txt_path)
	file, e := os.Create(txt_path)
	checkError(e)
	defer file.Close()
	_, e = file.WriteString(txt_content)
	checkError(e)
	fi, e := file.Stat()
	txt_fileSize = fi.Size()
}
func createZipWithPassword() {
	fmt.Println("Creating", zip_path)
	commandString := fmt.Sprintf(`7za a %s %s -p"%s" -mem=%s`, zip_path, txt_path, zip_password, zip_encryptType)
	commandSlice := strings.Fields(commandString)
	fmt.Println(commandString)
	c := exec.Command(commandSlice[0], commandSlice[1:]...)
	e := c.Run()
	checkError(e)
}
func extractZipWithPassword() {
	fmt.Printf("Unzipping `%s` to directory `%s`\n", zip_path, extract_path)
	commandString := fmt.Sprintf(`7za e %s -o%s -p"%s" -aoa`, zip_path, extract_path, zip_password)
	commandSlice := strings.Fields(commandString)
	fmt.Println(commandString)
	c := exec.Command(commandSlice[0], commandSlice[1:]...)
	e := c.Run()
	checkError(e)
}
func checkFor7Zip() {
	_, e := exec.LookPath("7za")
	if e != nil {
		fmt.Println("Make sure 7zip is install and include your path.")
	}
	checkError(e)
}
func checkExtractedFile() {
	fmt.Println("Reading", extracted_txt_path)
	file, e := os.Open(extracted_txt_path)
	checkError(e)
	defer file.Close()
	buf := make([]byte, txt_fileSize)
	n, e := file.Read(buf)
	checkError(e)
	if !strings.Contains(string(buf[:n]), strings.Fields(txt_content)[0]) {
		panic(fmt.Sprintf("File`%s` is corrupted.\n", extracted_txt_path))
	}
}
func main() {
	fmt.Println("# Setup")
	checkFor7Zip()
	setupTestDir()
	createSampleFile()
	createZipWithPassword()
	fmt.Println("# Answer to question...")
	extractZipWithPassword()
	checkExtractedFile()
	fmt.Println("Done.")
}
