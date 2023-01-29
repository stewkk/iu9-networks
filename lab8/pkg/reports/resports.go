package reports

import (
	"os"
	"os/exec"
	"path"
	"text/template"
)

type Fields struct {
	WorkType  string `json:"jobType"`
	Title     string `json:"jobName"`
	Author    string `json:"author"`
	Teacher   string `json:"teacher"`
	Group     string `json:"group"`
	Course    string `json:"course"`
	LabNumber string `json:"number"`
	Body      string `json:"report"`
	Year	  string
}

const titleTemplate = `
\documentclass{iu9lab}
\worktype{$$.WorkType$$}
\title{$$.Title$$}
\author{$$.Author$$}
\teacher{$$.Teacher$$}
\group{$$.Group$$}
\course{$$.Course$$}
\labnumber{$$.LabNumber$$}
\myyear{$$.Year$$}
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

func NewReportGenerator() (ReportGenerator, error) {
	tmpl, err := template.New("titlepage").Delims("$$", "$$").Parse(titleTemplate)
	if err != nil {
		return ReportGenerator{}, err
	}

	return ReportGenerator{
		tmpl: tmpl,
	}, nil
}

func (g *ReportGenerator) GenerateTitle(fields *Fields, basename string) ([]byte, error) {
	tex := basename+".tex"
	texFile, err := os.Create(tex)
	if err != nil {
		return nil, err
	}
	defer texFile.Close()

	err = g.tmpl.Execute(texFile, fields)
	if err != nil {
		return nil, err
	}

	cmd := exec.Command("lualatex",
		"-interaction",
		"nonstopmode",
		"--output-directory="+path.Dir(basename),
		tex,
	)
	return cmd.CombinedOutput()
}

func CompileMarkdown(mdPath, resPath string) ([]byte, error) {
	cmd := exec.Command("pandoc",
		"-s",
		"--template",
		templatePath,
		"--pdf-engine=lualatex",
		"--from",
		"markdown",
		"--to",
		"pdf",
		mdPath,
		"-o",
		resPath,
	)
	return cmd.CombinedOutput()
}

func MergePdfs(lhs string, rhs string, res string) ([]byte, error) {
	cmd := exec.Command("pdfunite",
		lhs,
		rhs,
		res,
	)
	return cmd.CombinedOutput()
}
