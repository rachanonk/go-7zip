// Shows how to extract an passsword encrypted zip file using 7zip.
// By Larry Battle <https://github.com/LarryBattle>
// Answer to http://stackoverflow.com/questions/20330210/golang-1-2-unzip-password-protected-zip-file
// 7-zip.chm - http://sevenzip.sourceforge.jp/chm/cmdline/switches/index.htm
// Effective Golang - http://golang.org/doc/effective_go.html
package main

import (
	"fmt"

	"github.com/rachanonk/go-7zip/internal/zip"
)

func main() {
	fmt.Println("# Setup")
	zip.CheckFor7Zip()
	zip.SetupDir()
	zip.CreateSampleFile()
	zip.CreateZipWithPassword()
	fmt.Println("Done.")
}
