package main

import (
	"bytes"
	"fmt"
	"html/template"
	"net/http"
	"os"
	"os/exec"
)

func writePDF(
	w http.ResponseWriter,
	tmpl *template.Template,
	data any,
	filename string,
) error{
	var html bytes.Buffer

	if err := tmpl.Execute(&html, data); err != nil{
		return err
	}

	pdf, err := generatePDFWithWkhtmltopdf(html.String())
	if err != nil{
		return err
	}

	w.Header().Set("Content-Type", "application/pdf")
	w.Header().Set(
		"Content-Disposition",
		fmt.Sprintf("attachment; filename=%s", filename),
	)
	_, err = w.Write(pdf)
	return err
}


func generatePDFWithWkhtmltopdf(htmlContent string) ([]byte, error) {
	// Временный HTML
	tmpHTML, err := os.CreateTemp("", "order-*.html")
	if err != nil {
		return nil, err
	}
	defer os.Remove(tmpHTML.Name())

	if _, err := tmpHTML.WriteString(htmlContent); err != nil {
		return nil, err
	}
	tmpHTML.Close()

	// Временный PDF
	tmpPDF, err := os.CreateTemp("", "order-*.pdf")
	if err != nil {
		return nil, err
	}
	tmpPDF.Close()
	defer os.Remove(tmpPDF.Name())

	// Запуск wkhtmltopdf
	cmd := exec.Command(
		"wkhtmltopdf",
		"--enable-local-file-access", // чтобы загружались CSS и изображения
		"--encoding", "utf-8",
		tmpHTML.Name(),
		tmpPDF.Name(),
	)

	if output, err := cmd.CombinedOutput(); err != nil {
		return nil, fmt.Errorf("wkhtmltopdf: %v\n%s", err, string(output))
	}

	return os.ReadFile(tmpPDF.Name())
}