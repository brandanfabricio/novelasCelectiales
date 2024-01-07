package lib

import (
	"Novelas/internal/novela/domain"

	"regexp"
	"strings"
	"sync"

	"github.com/gocolly/colly/v2"
	"github.com/google/uuid"
)

func ExtrarNovle(url string) domain.Novela {

	uuidWithHyphen := uuid.New()
	uuid := strings.Replace(uuidWithHyphen.String(), "-", "", -1)

	scrap := colly.NewCollector()

	var wg sync.WaitGroup
	wg.Add(1)

	novela := domain.Novela{}
	novela.Id = uuid

	scrap.OnHTML("div.site-content", func(h *colly.HTMLElement) {

		defer wg.Done()

		h.ForEach("div.post-title", func(_ int, titulo *colly.HTMLElement) {

			// fmt.Println(titulo.Text)
			// titulo = titu.Text
			novela.Titulo = strings.TrimSpace(titulo.Text)

		})

		h.ForEach("div.summary_image", func(_ int, img *colly.HTMLElement) {

			imgSrc := img.ChildAttr("img", "src")

			novela.Imagen = imgSrc

			// if imgSrc != "" {
			// 	fmt.Println(imgSrc)
			// } else {
			// 	fmt.Println("No se encontró ninguna imagen dentro del div.")
			// }

		})

		h.ForEach("div.summary__content", func(i int, des *colly.HTMLElement) {

			conte := des.DOM.Find("p").Text()
			novela.Description = conte

			// fmt.Println(conte)

		})

		h.ForEach("[class^=lcp_catlist]  li:nth-child(1)", func(i int, pagina *colly.HTMLElement) {

			// ultimaPagina := pagina.ChildAttr("li","title")
			expresion := regexp.MustCompile(`(\d+)`)

			ultimoCapitulo := expresion.FindStringSubmatch(pagina.Text)
			if len(ultimoCapitulo) == 0 {

				novela.Paginas = "-----"

			} else {

				novela.Paginas = ultimoCapitulo[0]
			}

			// fmt.Println(ultimoCapitulo[0])

		})

		h.ForEach("ul.lcp_paginator li:nth-child(8)", func(i int, pagina *colly.HTMLElement) {

			url := pagina.ChildAttr("a", "href")
			// fmt.Println(pagina.Text)
			// fmt.Println(url)
			novela.Url = url

		})

	})

	scrap.Visit(url)

	wg.Wait() // Esperar a que todas las funciones ForEach hayan terminado

	return novela
}
