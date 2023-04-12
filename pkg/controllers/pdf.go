package controllers

import (
	"fmt"
	"io"
	"net/http"
	"os"

	"github.com/gofiber/fiber/v2"
	"github.com/signintech/gopdf"
)

func PdfTester(c *fiber.Ctx) error {
	var err error

	// // Download a Font
	// fontUrl := "https://github.com/google/fonts/raw/master/ofl/daysone/DaysOne-Regular.ttf"
	// if err = DownloadFile("example-font.ttf", fontUrl); err != nil {
	// 	panic(err)
	// }

	// // Download a PDF
	// fileUrl := "https://tcpdf.org/files/examples/example_012.pdf"
	// if err = DownloadFile("example-pdf.pdf", fileUrl); err != nil {
	// 	panic(err)
	// }

	pdf := gopdf.GoPdf{}
	pdf.Start(gopdf.Config{PageSize: gopdf.Rect{W: 612, H: 792}}) //595.28, 841.89 = A4

	pdf.AddPage()
	pdf.AddPage()

	err = pdf.AddTTFFont("arial", "/Users/g.tan/Projects/Dashboards/public/arial.ttf")
	if err != nil {
		fmt.Print(err.Error())
	}

	err = pdf.SetFont("arial", "", 8.5)
	if err != nil {
		panic(err)
	}

	// Color the page
	// pdf.SetLineWidth(0.5)
	// pdf.SetFillColor(124, 252, 0) //setup fill color
	// pdf.RectFromUpperLeftWithStyle(50, 100, 400, 600, "FD")
	// pdf.SetFillColor(0, 0, 0)

	// Import page 1
	tpl1 := pdf.ImportPage("CARMELA ePN Template.pdf", 1, "/MediaBox")

	// Draw pdf onto page
	pdf.UseImportedTemplate(tpl1, 0, 0, 0, 0)

	// Write on top of
	pdf.SetX(355)
	pdf.SetY(132.1)
	pdf.MultiCellWithOption(&gopdf.Rect{
		W: 185,
		H: 50,
	}, "PN No. 1234567890123456789023456789023456789023456789023456789 \n asdasfada", gopdf.CellOption{
		Align:                  gopdf.Right,
		CoefLineHeight:         200,
		CoefUnderlinePosition:  1,
		CoefUnderlineThickness: 1,
	})

	err = pdf.WritePdf("test.pdf")
	if err != nil {
		panic(err)
	}

	return c.SendString("Success")
}

// DownloadFile will download a url to a local file. It's efficient because it will
// write as it downloads and not load the whole file into memory.
func DownloadFile(filepath string, url string) error {
	// Get the data
	resp, err := http.Get(url)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	// Create the file
	out, err := os.Create(filepath)
	if err != nil {
		return err
	}
	defer out.Close()

	// Write the body to file
	_, err = io.Copy(out, resp.Body)
	return err
}

// func PdfTester2(c *fiber.Ctx) error {
// 	r := u.NewRequestPdf("")

// 	//html template path
// 	templatePath := "templates/sample.html"

// 	//path for download pdf
// 	outputPath := "storage/example.pdf"

// 	//html template data
// 	templateData := struct {
// 		Title       string
// 		Description string
// 		Company     string
// 		Contact     string
// 		Country     string
// 	}{
// 		Title:       "HTML to PDF generator",
// 		Description: "This is the simple HTML to PDF file.",
// 		Company:     "Jhon Lewis",
// 		Contact:     "Maria Anders",
// 		Country:     "Germany",
// 	}

// 	if err := r.ParseTemplate(templatePath, templateData); err == nil {
// 		ok, _ := r.GeneratePDF(outputPath)
// 		fmt.Println(ok, "pdf generated successfully")
// 	} else {
// 		fmt.Println(err)
// 	}

// 	return c.SendString("success")
// }
