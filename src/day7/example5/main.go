package main

import (
	"archive/tar"
	"os"
	"fmt"
	"io"
)

func main()  {
	fileName := "/Users/abner/workspace/work/Desktop/mysql_群文件/Percona-XtraBackup-2.4.3-r6a46905-el6-x86_64-bundle.tar"
	file, err := os.Open(fileName)
	if err != nil {
		fmt.Println("error")
	}
	tarFiles := tar.NewReader(file)
	for f, err := tarFiles.Next(); err != io.EOF ; f, err = tarFiles.Next(){
		if err != nil {
			fmt.Println("aa", err)
			return
		}
		fileInfo := f.FileInfo()
		newFile, _ := os.OpenFile(f.Name, os.O_WRONLY|os.O_CREATE, fileInfo.Mode())
		io.Copy(newFile, tarFiles)
		newFile.Close()
		fmt.Println(fileInfo.Name(), fileInfo.Mode())

	}


}
