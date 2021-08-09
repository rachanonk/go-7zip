package zip

import (
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
	"strings"

	"io/ioutil"
)

var (
	txt_content     = "1 | 2 | 3 | 4"
	txt_filename    = "content.txt"
	zip_filename    = "sample.zip"
	result_filename = "result.zip"
	zip_password    = "1234"
	zip_encryptType = "AES256"
	base_path       = "../"

	file_path          = filepath.Join(base_path, "files")
	src_path           = filepath.Join(file_path, "src")
	extract_path       = filepath.Join(file_path, "extracted")
	zip_path           = filepath.Join(file_path, "zip")
	sameple_zip_path   = filepath.Join(file_path, zip_filename)
	sameple_txt_path   = filepath.Join(src_path, txt_filename)
	extracted_txt_path = filepath.Join(extract_path, txt_filename)
	dst_txt_path       = filepath.Join(zip_path, txt_filename)
	dst_zip_path       = filepath.Join(zip_path, result_filename)
)

var txt_fileSize int64

func CheckError(e error) {
	if e != nil {
		panic(e)
	}
}
func SetupZipDir() {
	var e error
	fmt.Printf("Removing `%s`\n", zip_path)
	os.Remove(zip_path)
	fmt.Printf("Creating `%s`\n", zip_path)
	e = os.MkdirAll(zip_path, os.ModeDir|os.ModePerm)
	CheckError(e)
}
func ConvertContent() {
	fmt.Printf("Converting `%s` to `%s`\n", extracted_txt_path, dst_txt_path)
	input, e := ioutil.ReadFile(extracted_txt_path)
	CheckError(e)

	lines := strings.Split(string(input), "\n")

	for i := range lines {
		lines[i] = strings.Replace(lines[i], "|", ",", 4)
	}
	output := strings.Join(lines, "\n")
	e = ioutil.WriteFile(dst_txt_path, []byte(output), 0644)
	CheckError(e)
}
func ReCreateZipWithPassword() {
	fmt.Println("Creating", dst_zip_path)
	commandString := fmt.Sprintf(`7za a %s %s -p"%s" -mem=%s`, dst_zip_path, dst_txt_path, zip_password, zip_encryptType)
	commandSlice := strings.Fields(commandString)
	fmt.Println(commandString)
	c := exec.Command(commandSlice[0], commandSlice[1:]...)
	e := c.Run()
	CheckError(e)
}

func SetupDir() {
	var e error
	fmt.Printf("Removing `%s`\n", file_path)
	os.Remove(file_path)
	fmt.Printf("Creating `%s`,`%s`\n", extract_path, src_path)
	e = os.MkdirAll(src_path, os.ModeDir|os.ModePerm)
	CheckError(e)
	e = os.MkdirAll(extract_path, os.ModeDir|os.ModePerm)
	CheckError(e)
}
func CreateSampleFile() {
	fmt.Println("Creating", sameple_txt_path)
	file, e := os.Create(sameple_txt_path)
	CheckError(e)
	defer file.Close()
	_, e = file.WriteString(txt_content)
	CheckError(e)
	fi, e := file.Stat()
	txt_fileSize = fi.Size()
}
func CreateZipWithPassword() {
	fmt.Println("Creating", sameple_zip_path)
	commandString := fmt.Sprintf(`7za a %s %s -p"%s" -mem=%s`, sameple_zip_path, sameple_txt_path, zip_password, zip_encryptType)
	commandSlice := strings.Fields(commandString)
	fmt.Println(commandString)
	c := exec.Command(commandSlice[0], commandSlice[1:]...)
	e := c.Run()
	CheckError(e)
}
func ExtractZipWithPassword() {
	fmt.Printf("Unzipping `%s` to directory `%s`\n", sameple_zip_path, extract_path)
	commandString := fmt.Sprintf(`7za e %s -o%s -p"%s" -aoa`, sameple_zip_path, extract_path, zip_password)
	commandSlice := strings.Fields(commandString)
	fmt.Println(commandString)
	c := exec.Command(commandSlice[0], commandSlice[1:]...)
	e := c.Run()
	CheckError(e)
}
func CheckFor7Zip() {
	_, e := exec.LookPath("7za")
	if e != nil {
		fmt.Println("Make sure 7zip is install and include your path.")
	}
	CheckError(e)
}

func VerifyZip() {
	fmt.Printf("Unzipping `%s` to directory `%s`\n", dst_zip_path, extract_path)
	commandString := fmt.Sprintf(`7za e %s -o%s -p"%s" -aoa`, dst_zip_path, extract_path, zip_password)
	commandSlice := strings.Fields(commandString)
	fmt.Println(commandString)
	c := exec.Command(commandSlice[0], commandSlice[1:]...)
	e := c.Run()
	CheckError(e)
}
