package lib

import (
	"Novelas/internal/capitulo/domain"
	"fmt"
	"log"
	"net/http"
	"regexp"
	"strconv"
	"strings"
	"sync"
	"time"

	"github.com/gocolly/colly/v2"
)

func ExtrarCapitulo(novela_id string, url string) []domain.Capitulo {

	scrap := colly.NewCollector(
		colly.Async(true),
		colly.MaxDepth(5),
		colly.IgnoreRobotsTxt(),
		// colly.
		// colly.UserAgent("Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/91.0.4472.124 Safari/537.36"),
		// colly.SetRequestTimeou(30*time.Second), // Aumentar el tiempo límite de la solicitud a 30 segundos
	)
	scrap.WithTransport(&http.Transport{
		ResponseHeaderTimeout: 2 * time.Minute,
	})

	// scrap := colly.NewCollector(
	// 	colly.Async(true),
	// 	colly.Async(true),
	// 	colly.MaxDepth(1),
	// 	colly.IgnoreRobotsTxt(),
	// 	colly.UserAgent("Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/91.0.4472.124 Safari/537.36"),
	// )

	var url_novelas []string
	var wg sync.WaitGroup
	var mutex sync.Mutex
	wg.Add(1)

	capitulos := []domain.Capitulo{}

	scrap.OnHTML("[class^=lcp_catlist]", func(h *colly.HTMLElement) {
		defer wg.Done()

		url_novelas = h.ChildAttrs("a", "href")

	})

	scrap.Visit(url)
	wg.Wait() // Esperar a que todas las funciones ForEach hayan terminado

	// leng := len(url_novelas) - 1
	// outerLoop:
	var listErr []string

	for _, urlNovela := range url_novelas {
		wg.Add(1)

		go func(url string, id string) {

			defer wg.Done()

			capitulo, err := extrerContendido(url, id)
			if err != nil {
				// fmt.Printf("Error al extraer capítulo de %s: %v\n", url, err)
				listErr = append(listErr, url)
				return
			}
			// Puedes agregar el capítulo a una lista segura usando un mutex si es necesario
			mutex.Lock()
			capitulos = append(capitulos, capitulo)
			mutex.Unlock()

		}(urlNovela, novela_id)
		// break outerLoop
	}

	wg.Wait() // Esperar a que todas las funciones ForEach hayan terminado

	// for _, v := range capitulos {

	// 	fmt.Printf("%d ", v.Ncap)
	// }

	if len(listErr) > 0 {
		log.Println("################## Extrayendo Error ####################")
		for _, link := range listErr {
			wg.Add(1)
			log.Println(link)

			go func(url string, id string) {

				defer wg.Done()

				capitulo, err := extrerContendido(url, id)
				if err != nil {
					// fmt.Println(err)
					// fmt.Printf("Error al extraer capítulo de %s\n", url)
					listErr = append(listErr, url)
					return
				}
				// Puedes agregar el capítulo a una lista segura usando un mutex si es necesario
				mutex.Lock()
				capitulos = append(capitulos, capitulo)
				mutex.Unlock()

			}(link, novela_id)
			// break outerLoop
		}

		wg.Wait()

	}

	return capitulos
}

func extrerContendido(url string, novela_id string) (domain.Capitulo, error) {
	scrap := colly.NewCollector()

	capitulo := domain.Capitulo{}
	capitulo.NovelaId = novela_id

	scrap.OnHTML("div.main-col-inner", func(container *colly.HTMLElement) {

		container.ForEach("p.SomeClass", func(i int, contenido *colly.HTMLElement) {
			capitulo.Contenido += contenido.Text + "\n" + "<br/><br/>"
		})

		container.ForEach("p.SomeClass:nth-child(1)", func(i int, data *colly.HTMLElement) {

			capitulo.Titulo = strings.TrimSpace(data.Text)
			capitulo.Cap = strings.TrimSpace(data.Text)
			// fmt.Println(capitulo.Titulo)

		})

		if capitulo.Titulo == "" || capitulo.Cap == "" {
			container.ForEach("h2, h3", func(_ int, elem *colly.HTMLElement) {

				if capitulo.Titulo == "" || capitulo.Cap == "" {
					capitulo.Titulo = strings.TrimSpace(elem.Text)
					capitulo.Cap = strings.TrimSpace(elem.Text)
				}

			})

		}

		container.ForEach("[class^=item-title_]", func(i int, title *colly.HTMLElement) {

			data := title.Text
			re := regexp.MustCompile(`\d+`)
			match := re.FindString(data)
			ncap, err := strconv.Atoi(match)
			if err != nil {
				fmt.Println("Errro de conversion de string a numero ")
			}
			capitulo.Ncap = ncap
			if capitulo.Titulo == "" || capitulo.Cap == "" {

				capitulo.Titulo = strings.TrimSpace(data)
				capitulo.Cap = strings.TrimSpace(data)
			}

		})

		if capitulo.Contenido == "" {
			container.ForEach("div p", func(i int, cont *colly.HTMLElement) {
				capitulo.Contenido += cont.Text + "\n" + "<br/><br/>"

			})

		}

	})

	scrap.OnError(func(r *colly.Response, err error) {

		err = r.Request.Retry()

	})

	if err := scrap.Visit(url); err != nil {
		return capitulo, err
	}

	return capitulo, nil
}
