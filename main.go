package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"
)

func generateFileList(rootPath string) []string {
	// 遍历指定目录下的所有md文件，生成文件路径列表
	var fileList []string
	err := filepath.Walk(rootPath, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}
		if !info.IsDir() && strings.HasSuffix(info.Name(), ".md") {
			fileList = append(fileList, path)
		}
		return nil
	})
	if err != nil {
		fmt.Println(err)
	}
	return fileList
}

func generateFileTree(fileList []string, rootPath string) map[string]interface{} {
	// 将文件路径列表转换为树形结构
	fileTree := make(map[string]interface{})
	for _, filePath := range fileList {
		relPath, err := filepath.Rel(rootPath, filePath)
		if err != nil {
			fmt.Println(err)
			continue
		}
		dirList := strings.Split(filepath.Dir(relPath), string(os.PathSeparator))
		node := fileTree
		for _, dirName := range dirList {
			if _, ok := node[dirName]; !ok {
				node[dirName] = make(map[string]interface{})
			}
			node = node[dirName].(map[string]interface{})
		}
		fileName := strings.TrimSuffix(filepath.Base(relPath), ".md")
		node[fileName] = filePath
	}
	return fileTree
}

func generateFileTxt(fileTree map[string]interface{}, indent int) string {
	// 将文件树形结构转换为txt文件格式
	var fileTxt string
	for name, node := range fileTree {
		if filePath, ok := node.(string); ok {
			fileTxt += fmt.Sprintf("%s- %s: %s\n", strings.Repeat("  ", indent), name, filePath)
		} else {
			subTree := node.(map[string]interface{})
			fileTxt += fmt.Sprintf("%s- %s:\n%s", strings.Repeat("  ", indent), name, generateFileTxt(subTree, indent+1))
		}
	}
	return fileTxt
}

func main() {
	var (
		rootPath   string
		outputPath string
	)
	flag.StringVar(&rootPath, "p", "", "要遍历的文件夹路径")
	flag.StringVar(&outputPath, "o", "file_list.txt", "输出文件的路径和文件名")
	flag.Parse()

	if rootPath == "" {
		fmt.Println("请使用 -p 参数指定要遍历的文件夹路径")
		return
	}

	fileList := generateFileList(rootPath)
	fileTree := generateFileTree(fileList, rootPath)
	fileTxt := generateFileTxt(fileTree, 0)

	err := ioutil.WriteFile(outputPath, []byte(fileTxt), 0644)
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println("已生成文件列表文件")
}
