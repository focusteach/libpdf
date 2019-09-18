package main


import (
	"github.com/focusteach/libpdf/creator"
)

func main() {
	var paths []string
	var oPaths string
	ImagesToPdf(paths, oPaths)
}


// ImagesToPdf Images to PDF.
func ImagesToPdf(inputPaths []string, outputPath string) error {
	c := creator.New()

	for _, imgPath := range inputPaths {


		img, err := c.NewImageFromFile(imgPath)
		if err != nil {
			return err
		}
		img.ScaleToWidth(612.0)

		// Use page width of 612 points, and calculate the height proportionally based on the image.
		// Standard PPI is 72 points per inch, thus a width of 8.5"
		height := 612.0 * img.Height() / img.Width()
		c.SetPageSize(creator.PageSize{612, height})
		c.NewPage()
		img.SetPos(0, 0)
		_ = c.Draw(img)
	}

	err := c.WriteToFile(outputPath)
	return err
}