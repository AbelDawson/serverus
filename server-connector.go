// This file is heavily inspired by goftp file-driver. 
//package driver
package main

import (
	"errors"
	"fmt"
	"os"
	"io"
	//"strings"
)


func deleteDir (path string) error {
	f, err := os.Lstat(path)
	if err != nil {
		return err
	}
	if f.IsDir() {
		return os.Remove(path)
	}
	return errors.New("Not a directory")
}

func deleteFile(path string) error {
	f, err := os.Lstat(path)
	if err != nil {
		return err
	}
	if !f.IsDir(){
		return os.Remove(path)
	}
	return errors.New("Not a file")
}

func rename(oldPath string, newPath string) error {
	return os.Rename(oldPath, newPath)
}

func getFile (path string) (io.ReadCloser, error) {
	f, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	return f, nil
}

func putFile (path string, data io.Reader) error {
	fmt.Println(path, data)
	f, err := os.Lstat(path)
	if err == nil {
		if f.IsDir() {
			return errors.New("A dir has the same name")
		}
	} else if os.IsExist(err) {
		return errors.New("Put file error")
	}

	of, err := os.Create(path)
	if err != nil {
		return err
	}
	defer of.Close()
	/*bytes, err := io.Copy(of, data)
	if err != nil {
		return err
	}*/
	return nil
}
/*
// test main function
func main(){
	path := "/home/abel/programming/serverusSnape//server-connector.go"
	_, err := getFile(path)
	if err != nil {
		fmt.Printf("Error: $s\n", err)
	}
	err = putFile("/home/abel/programming/serverusSnape/hello.txt", strings.NewReader("Hello from hello.txt\n"))
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println("File recieved")
}
*/
