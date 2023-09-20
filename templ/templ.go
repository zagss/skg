package templ

import (
	"log"
	"os"
	"text/template"
)

type CmdConfig struct {
	ConfigPath string
	RouterPath string
}

func ReplaceTempl(filePath, text string, data any) {
	// 根据指定模版文本生成template
	tmpl, err := template.New("").Parse(text)
	if err != nil {
		log.Fatalln("create template failed, err:", err)
	}
	// 模版渲染，并写入文件
	f, err := os.OpenFile(filePath, os.O_WRONLY|os.O_CREATE, 0644)
	if err != nil {
		log.Fatalln(err)
	}
	defer f.Close()
	if err := tmpl.Execute(f, data); err != nil {
		log.Fatalln(err)
	}
}

const ModTemplate = `module {{.Module}}

go {{.Version}}`

const CmdTemplate = `package main

import (
	"{{.}}/config"
	"{{.}}/router"
)

func main() {
	config.InitConfig()
	r := router.NewRouter()
	_ = r.Run(config.Config.System.HttpPort)
}
`
