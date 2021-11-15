package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"os/exec"
	"strings"
)

func main() {
	RunAbiGen()
	//RunAbiGenWithoutBin()
}

// RunAbiGenWithoutBin 只通过abi序列化
func RunAbiGenWithoutBin() {
	abiPath := "abi"
	outPath := "go"

	// 遍历abi目录
	abiFiles, _ := GetAllFile(abiPath)
	fmt.Println(abiFiles)

	for _, abi := range abiFiles {
		abiArgs := abiPath + "/" + abi
		pkgArgs := strings.ToLower(strings.Replace(abi, ".abi", "", 1))
		outArgs := outPath + "/" + pkgArgs + "/" + pkgArgs + ".go"

		fmt.Println(outArgs)
		fmt.Println(PathExists(outPath + "/" + pkgArgs))
		fmt.Println(abi)

		cmd := exec.Command("abigen", "--abi", abiArgs, "--pkg", pkgArgs, "--out", outArgs)
		// 执行命令，并返回结果
		out, err := cmd.Output()
		if err != nil {
			fmt.Println(err.Error())
		}
		fmt.Println(string(out))
	}
}

// RunAbiGen 通过abi和bin序列化
func RunAbiGen() {
	// 通过exec.Command函数执行命令或者shell
	abiPath := "abi"
	binPath := "bin"
	outPath := "go"

	// 遍历abi目录
	abiFiles, _ := GetAllFile(abiPath)
	fmt.Println(abiFiles)

	for _, abi := range abiFiles {
		abiArgs := abiPath + "/" + abi
		binArgs := binPath + "/" + strings.Replace(abi, ".abi", ".bin", 1)
		pkgArgs := strings.ToLower(strings.Replace(abi, ".abi", "", 1))
		outArgs := outPath + "/" + pkgArgs + "/" + pkgArgs + ".go"

		fmt.Println(outArgs)
		fmt.Println(PathExists(outPath + "/" + pkgArgs))
		fmt.Println(abi)

		cmd := exec.Command("abigen", "--abi", abiArgs, "--bin", binArgs, "--pkg", pkgArgs, "--out", outArgs)
		// 执行命令，并返回结果
		out, err := cmd.Output()
		if err != nil {
			panic(err.Error())
		}
		fmt.Println(string(out))
	}
}

//PathExists 判断文件夹是否存在
func PathExists(path string) (bool, error) {
	_, err := os.Stat(path)
	if err == nil {
		return true, nil
	}
	if os.IsNotExist(err) {
		// 创建文件夹
		err := os.MkdirAll(path, os.ModePerm)
		if err != nil {
			fmt.Printf("mkdir failed![%v]\n", err)
		} else {
			return true, nil
		}
	}
	return false, err
}

// GetAllFile 遍历目录
func GetAllFile(pathname string) ([]string, error) {
	var files []string
	rd, err := ioutil.ReadDir(pathname)
	for _, fi := range rd {
		if fi.IsDir() {
			files = append(files, pathname+"/"+fi.Name())
			fmt.Printf("[%s]\n", pathname+"/"+fi.Name())
			tmpFiles, _ := GetAllFile(pathname + fi.Name() + "/")
			for _, f := range tmpFiles {
				files = append(files, f)
			}

		} else {
			files = append(files, fi.Name())
		}
	}
	return files, err
}
