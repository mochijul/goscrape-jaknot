package scrape

import (
	"fmt"
	b "goscrape/mysql"
	"strings"

	"github.com/gocolly/colly"
)

type IndexProduk struct {
	link string
	nama string
}

func GetIndexProduct(c *colly.Collector) {
	b.RunExec("CREATE TABLE IF NOT EXISTS goscrape.index_produk (id INT AUTO_INCREMENT PRIMARY KEY, nama VARCHAR(255), link VARCHAR(255) UNIQUE, created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP, updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP, last_scraped_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,ERR VARCHAR(150));", "Create Table index_produk")

	//init array index produk
	var indexProduk []IndexProduk

	c.OnHTML("tr td:nth-child(1)", func(e *colly.HTMLElement) {
		e.ForEach("a", func(_ int, el *colly.HTMLElement) {
			indexProduk = append(indexProduk, IndexProduk{
				link: el.Attr("href"),
				nama: strings.ReplaceAll(el.Text, "'", ""),
			})
		})
	})

	c.Visit("https://www.jakartanotebook.com/index-product")

	//insert into index produk batch
	batchInsert := ""
	countDataInsert := 0
	indexProdukChunk := chunk(indexProduk, 100)
	for _, v := range indexProdukChunk {
		for _, v2 := range v {
			//if last index v2
			if v2 != v[len(v)-1] {
				batchInsert = batchInsert + "('" + v2.nama + "', '" + v2.link + "'),"
			} else {
				batchInsert = batchInsert + "('" + v2.nama + "', '" + v2.link + "')"
			}
		}
		b.RunExec("INSERT INTO goscrape.index_produk (nama, link) VALUES "+batchInsert+" AS new ON DUPLICATE KEY UPDATE nama=new.nama, link=new.link", "Insert into index produk")

		batchInsert = ""
		countDataInsert = countDataInsert + len(v)
	}
	fmt.Println("Total data index produk inserted : ", countDataInsert)

}

func chunk(slice []IndexProduk, chunkSize int) [][]IndexProduk {
	var chunks [][]IndexProduk
	for i := 0; i < len(slice); i += chunkSize {
		end := i + chunkSize

		if end > len(slice) {
			end = len(slice)
		}

		chunks = append(chunks, slice[i:end])
	}
	return chunks
}
