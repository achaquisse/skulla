package internal

import (
	"bytes"
	"errors"
	"github.com/signintech/gopdf"
)

func GenerateCertificate(template string, studentName string, certDescription string) ([]byte, error) {
	if template != "english" && template != "portuguese" {
		return nil, errors.New("invalid certificate template: " + template)
	}

	var pdf = gopdf.GoPdf{}
	pdf.Start(gopdf.Config{
		PageSize: *gopdf.PageSizeA4Landscape,
	})

	pdf.AddPage()
	tpl := pdf.ImportPage("./assets/layouts/"+template+"_certificate_layout.pdf", 1, "/MediaBox")
	pdf.UseImportedTemplate(tpl, 0, 0, 0, 0)

	drawStudentName(&pdf, studentName)
	drawCertDescription(&pdf, certDescription)

	var b bytes.Buffer

	_, err := pdf.WriteTo(&b)
	if err != nil {
		return nil, err
	}
	return b.Bytes(), nil
}

func drawStudentName(pdf *gopdf.GoPdf, studentName string) {
	err := pdf.AddTTFFont("GreatVibes", "./assets/fonts/GreatVibes-Regular.ttf")
	if err != nil {
		panic("Couldn't open GreatVibes font.")
	}

	err = pdf.SetFont("GreatVibes", "", 46)
	if err != nil {
		panic(err)
	}

	pdf.SetXY(gopdf.PageSizeA4Landscape.W/2-202, gopdf.PageSizeA4Landscape.H/2-43)
	pdf.SetTextColor(128, 20, 22)

	rect := gopdf.Rect{W: 400, H: 60}

	err = pdf.CellWithOption(&rect, studentName, gopdf.CellOption{Align: gopdf.Center})
	if err != nil {
		panic("Unable to draw student name.")
	}
}

func drawCertDescription(pdf *gopdf.GoPdf, certDescription string) {
	err := pdf.AddTTFFont("IBMPlexSansCondensed", "./assets/fonts/IBMPlexSansCondensed-Regular.ttf")
	if err != nil {
		panic("Couldn't open IBMPlexSansCondensed font.")
	}

	err = pdf.SetFont("IBMPlexSansCondensed", "", 14)
	if err != nil {
		panic(err)
	}

	var cursorInitX = gopdf.PageSizeA4Landscape.W/2 - 262
	var cursorInitY = gopdf.PageSizeA4Landscape.H/2 + 30

	pdf.SetXY(cursorInitX, cursorInitY)
	pdf.SetTextColor(77, 78, 77)

	var boxWidth float64 = 520
	var boxHeight float64 = 18

	rect := gopdf.Rect{W: boxWidth, H: boxHeight}

	parts, err1 := pdf.SplitTextWithWordWrap(certDescription, boxWidth)
	if err1 != nil {
		panic(err1)
	}

	for i, part := range parts {
		pdf.SetXY(cursorInitX, cursorInitY+(float64(i)*boxHeight))

		err = pdf.CellWithOption(&rect, part, gopdf.CellOption{Align: gopdf.Center})
		if err != nil {
			panic("Unable to cert description.")
		}
	}
}
