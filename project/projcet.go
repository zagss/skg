package project

import (
	"bufio"
	"errors"
	"fmt"
	"os"
	"path"
	"path/filepath"
	"runtime"

	"github.com/zagss/skg/templ"
)

func Create(projectPath, projectName string) error {
	if projectPath == "" {
		projectPath = getWorkingDirPath()
	}
	if projectName == "" {
		projectName = "kiss"
	}
	p := path.Join(projectPath, projectName)

	// 删除文件夹
	err := os.RemoveAll(p)
	if err != nil {
		fmt.Println("Error removing file:", err)
		return err
	}

	apiPath := path.Join(p, "api")
	modPath := path.Join(p, "go.mod")
	configPath := path.Join(p, "config")
	routerPath := path.Join(p, "router")
	cmdPath := path.Join(p, "cmd")
	prePaths := []string{apiPath, modPath, configPath, routerPath, cmdPath}

	for _, prePath := range prePaths {
		// 获取文件后缀名
		// fileName, fileExt := filepath.Split(prePath)
		ext := filepath.Ext(prePath)
		if ext == "" {
			if err := os.MkdirAll(prePath, 0755); err != nil {
				fmt.Println("Error creating folder:", err)
				return err
			}
		} else {
			f, err := os.Create(modPath)
			if err != nil {
				fmt.Printf("create %s failed, err: %v", modPath, err)
				return err
			}
			defer f.Close()
		}

		switch prePath {
		case modPath:
			version, _ := getGolangVersion()
			var modCfg struct {
				Module  string
				Version string
			}
			modCfg.Module = projectName
			modCfg.Version = version
			templ.ReplaceTempl(modPath, templ.ModTemplate, modCfg)
		case cmdPath:
			cmdFile := path.Join(cmdPath, "cmd.go")
			templ.ReplaceTempl(cmdFile, templ.CmdTemplate, projectName)
		default:
		}
	}
	return nil
}

func getWorkingDirPath() string {
	dir, err := os.Getwd()
	if err != nil {
		panic(err)
	}
	return dir
}

func getGolangVersion() (string, error) {
	ver := runtime.Version()
	if ver != "" {
		return ver[2:6], nil
	}
	return "", errors.New("golang 环境不存在")
}

func writeFile(path, content string) error {
	file, err := os.OpenFile(path, os.O_WRONLY|os.O_CREATE, 0666)
	if err != nil {
		fmt.Println("文件打开失败", err)
	}
	defer file.Close()
	// 写入文件时，使用带缓存的 *Writer
	write := bufio.NewWriter(file)
	_, err = write.WriteString(content)
	if err != nil {
		return err
	}
	err = write.Flush()

	return err
}
