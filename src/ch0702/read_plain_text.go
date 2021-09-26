package ch0702

import (
	"bytes"
	"fmt"
	"github.com/ledongthuc/pdf"
)

func main() {
	pdf.DebugOn = true
	filePath := "F:\\100-work\\01-code\\go\\src\\lagouGoLangStudy\\src\\ch0702\\test.pdf"
	content, err := readPdf(filePath) // Read local pdf file
	if err != nil {
		panic(err)
	}
	fmt.Println(content)
	return
}

func readPdf(path string) (string, error) {
	f, r, err := pdf.Open(path)
	// remember close file
	defer f.Close()
	if err != nil {
		return "", err
	}
	var buf bytes.Buffer
	b, err := r.GetPlainText()
	if err != nil {
		return "", err
	}
	buf.ReadFrom(b)
	return buf.String(), nil
}
