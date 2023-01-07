package reports

import (
	"os"
	"os/exec"
	"text/template"
)

type TitleFields struct {
	WorkType  string
	Title     string
	Author    string
	Teacher   string
	Group     string
	Course    string
	LabNumber string
}

const titleTemplate = `
\documentclass{iu9lab}
\worktype{{{.WorkType}}}
\title{{{.Title}}}
\author{{{.Author}}}
\teacher{{{.Teacher}}}
\group{{{.Group}}}
\course{{{.Course}}}
\labnumber{{{.LabNumber}}}
\begin{document}
\maketitle
\end{document}
`

var templatePath string

func init() {
	templatePath = os.Getenv("LAB8_TEMPLATE")
}

type ReportGenerator struct {
	tmpl *template.Template
}

func NewReportGenerator() (*ReportGenerator, error) {
	tmpl, err := template.New("titlepage").Parse(titleTemplate)
	if err != nil {
		return nil, err
	}

	return &ReportGenerator{
		tmpl: tmpl,
	}, nil
}

func (g *ReportGenerator) GenerateTitle(fields *TitleFields, resPath string) ([]byte, error) {
	resFile, err := os.Create(resPath)
	if err != nil {
		return nil, err
	}
	defer resFile.Close()

	err = g.tmpl.Execute(resFile, fields)
	if err != nil {
		return nil, err
	}

	cmd := exec.Command("lualatex",
		"-interaction",
		"nonstopmode",
		"resPath",
	)
	return cmd.CombinedOutput()
}

func CompileMarkdown(mdPath, resPath string) ([]byte, error) {
	cmd := exec.Command("pandoc",
		"-s",
		"--template",
		templatePath,
		"--pdf-engine=lualatex",
		"--to",
		"pdf",
		mdPath,
		"-o",
		resPath,
	)
	return cmd.CombinedOutput()
}
