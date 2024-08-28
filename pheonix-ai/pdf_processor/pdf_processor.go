package pdf_processor

import (
	"pheonix/pheonix"
)

func ProcessPDF(pdfPath string, ai *pheonix.PhoenixIA) (string, error) {
	err := ai.Ingest(pdfPath)
	if err != nil {
		return "", err
	}

	return "PDF procesado exitosamente", nil
}
