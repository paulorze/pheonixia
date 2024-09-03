package main

import (
	"bytes"
	"log"
	"phoenixia/engine"

	"github.com/joho/godotenv"
	"github.com/jung-kurt/gofpdf"
)

func GeneratePDF() ([]byte, error) {
	pdf := gofpdf.New("P", "mm", "A4", "")
	pdf.AddPage()

	// Set font
	pdf.SetFont("Arial", "", 12)

	// Add Lorem Ipsum text
	loremIpsum := "Las llamas, que pertenecen a la familia de los camélidos y están estrechamente relacionadas con las alpacas, guanacos y vicuñas, son animales nativos de las regiones andinas de América del Sur, particularmente en países como Perú, Bolivia, Ecuador y Chile; domesticadas hace aproximadamente 4,000 a 5,000 años por las culturas precolombinas de los Andes para servir como animales de carga, ya que pueden transportar hasta 25-30% de su peso corporal en condiciones de altitud y clima difíciles, y también son apreciadas por su lana suave y cálida que se utiliza en la confección de ropa y textiles tradicionales; además, su comportamiento social y su capacidad para adaptarse a diferentes tipos de terrenos y climas han hecho que sean animales valiosos tanto para la economía local como para la conservación del medio ambiente andino, donde juegan un papel importante en el pastoreo y el mantenimiento de los ecosistemas de alta montaña; las llamas tienen una estructura física adaptada a su entorno, con un sistema respiratorio y cardiovascular eficiente para soportar las bajas concentraciones de oxígeno en altitudes elevadas, y se caracterizan por su pelaje largo y espeso que varía en colores, desde blanco y negro hasta marrón y gris, además de tener una dieta herbívora que consiste principalmente en pasto, hojas y corteza; socialmente, las llamas viven en grupos jerárquicos y tienen una comunicación compleja que incluye vocalizaciones y lenguaje corporal para coordinarse y mantener la cohesión dentro del grupo, y a lo largo de la historia y la cultura andina, las llamas han sido símbolo de la vida y la supervivencia en las duras condiciones de los Andes, siendo representadas en arte, mitos y ceremonias que subrayan su importancia en la vida cotidiana y espiritual de las comunidades que las han domesticado y cuidado a lo largo de los siglos."
	pdf.MultiCell(0, 10, loremIpsum, "", "", false)

	// Output to a buffer
	var buf bytes.Buffer
	err := pdf.Output(&buf)
	if err != nil {
		return nil, err
	}

	return buf.Bytes(), nil
}

func main() {
	if err := godotenv.Load(); err != nil {
		log.Fatalf("Error loading .env file")
	}

	llm, err := engine.NewOllamaAI()
	if err != nil {
		log.Fatal(err)
	}

	// response := llm.Ask("Las llamas, que pertenecen a la familia de los camélidos y están estrechamente relacionadas con las alpacas, guanacos y vicuñas, son animales nativos de las regiones andinas de América del Sur, particularmente en países como Perú, Bolivia, Ecuador y Chile; domesticadas hace aproximadamente 4,000 a 5,000 años por las culturas precolombinas de los Andes para servir como animales de carga, ya que pueden transportar hasta 25-30% de su peso corporal en condiciones de altitud y clima difíciles, y también son apreciadas por su lana suave y cálida que se utiliza en la confección de ropa y textiles tradicionales; además, su comportamiento social y su capacidad para adaptarse a diferentes tipos de terrenos y climas han hecho que sean animales valiosos tanto para la economía local como para la conservación del medio ambiente andino, donde juegan un papel importante en el pastoreo y el mantenimiento de los ecosistemas de alta montaña; las llamas tienen una estructura física adaptada a su entorno, con un sistema respiratorio y cardiovascular eficiente para soportar las bajas concentraciones de oxígeno en altitudes elevadas, y se caracterizan por su pelaje largo y espeso que varía en colores, desde blanco y negro hasta marrón y gris, además de tener una dieta herbívora que consiste principalmente en pasto, hojas y corteza; socialmente, las llamas viven en grupos jerárquicos y tienen una comunicación compleja que incluye vocalizaciones y lenguaje corporal para coordinarse y mantener la cohesión dentro del grupo, y a lo largo de la historia y la cultura andina, las llamas han sido símbolo de la vida y la supervivencia en las duras condiciones de los Andes, siendo representadas en arte, mitos y ceremonias que subrayan su importancia en la vida cotidiana y espiritual de las comunidades que las han domesticado y cuidado a lo largo de los siglos.", "", "Que es una llama?")
	// log.Println(response)

	// pdfData, err := GeneratePDF()
	// if err != nil {
	// 	log.Fatal(err)
	// }

	// name, err := llm.PDFLoader(pdfData)
	// if err != nil {
	// 	log.Fatal(err)
	// }
	// log.Println(name)

	docs, err := llm.DocumentsRetriever("e14034b548364c23a6639b86416239c5", "Que simbolizan las llamas")
	if err != nil {
		log.Fatal(err)
	}

	log.Println(docs)

	response, err := llm.Ask(docs, "", "Que simbolizan las llamas")
	if err != nil {
		log.Fatal(err)
	}

	log.Println(response)
}
