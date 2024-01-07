package application

import (
	capitulo "Novelas/internal/capitulo/domain"
	"Novelas/internal/capitulo/domain/lib"
	novela "Novelas/internal/novela/domain"
	"fmt"
	"log"
	"net/url"
	"strconv"
	"sync"
)

type CaseCapitulointerface interface {
	FindAllCapitulo(novelaId string) Minovela
	Paginate(novela_id string, page int) []capitulo.Capitulo
	GetPage(novela_id string, page int) []capitulo.Capitulo
	FindCapitulo(id int, novela_id string, numero_cap int) capitulo.Capitulo
	AddCapitulo(novela_id string, url string)
}

type CaseCapitulo struct {
	NovelaRepository   novela.NovelaRepository
	CapituloRepository capitulo.CapituloRepository
}

type Minovela struct {
	Novela    novela.Novela
	Capitulos []capitulo.Capitulo
}

func NewCaseCApitulo(NovelaRepo novela.NovelaRepository, CapituloRepo capitulo.CapituloRepository) *CaseCapitulo {

	return &CaseCapitulo{NovelaRepository: NovelaRepo, CapituloRepository: CapituloRepo}

}
func (self *CaseCapitulo) FindAllCapitulo(novelaId string) Minovela {

	var data Minovela

	novela := self.NovelaRepository.GetById(novelaId)
	capitulos := self.CapituloRepository.GetCapituloById(novelaId)
	// Procesar los resultados obtenidos

	data.Novela = novela
	data.Capitulos = capitulos

	return data
}

func (self *CaseCapitulo) Paginate(novela_id string, page int) []capitulo.Capitulo {

	capitlos := self.CapituloRepository.GetCapituloPaginated(novela_id, page)

	return capitlos
}

func (this *CaseCapitulo) GetPage(novela_id string, page int) []capitulo.Capitulo {

	capitulo := this.CapituloRepository.GetPage(novela_id, page)

	return capitulo

}
func (this *CaseCapitulo) FindCapitulo(capitulo_id int, novela_id string, numero_cap int) capitulo.Capitulo {

	capitulo := this.CapituloRepository.FindCapitulo(capitulo_id, novela_id, numero_cap)

	return capitulo
}

func (this *CaseCapitulo) AddCapitulo(novela_id string, novela_url string) {

	parserUrl, err := url.Parse(novela_url)
	if err != nil {
		log.Println("url incorrecta")
	} else {

		params := parserUrl.Query().Get("lcp_page0")

		ultimaPage, err := strconv.Atoi(params)
		if err != nil {
			log.Println("erro en el parceo ")
		}

		func() {
			var wg sync.WaitGroup
			var mutex sync.Mutex

			for i := ultimaPage; i >= ultimaPage-1; i-- {

				wg.Add(1)
				go func(page int) {
					defer wg.Done()
					// Adquirir un "permiso" del semáforo
					// sema <- struct{}{}
					// defer func() { <-sema }() // Liberar el "permiso" al finalizar

					modifiedURL := *parserUrl
					query := modifiedURL.Query()
					query.Set("lcp_page0", strconv.Itoa(page))
					modifiedURL.RawQuery = query.Encode()

					fmt.Println("Página:", modifiedURL.String())

					capitulos := lib.ExtrarCapitulo(novela_id, modifiedURL.String())
					mutex.Lock()
					defer mutex.Unlock()
					this.CapituloRepository.AddCapitulo(capitulos)
				}(i)

			}
			wg.Wait()

		}()

		// Cerrar el canal de "done" en una gorutina después de que todas las otras hayan terminado

		modifiedURL := *parserUrl
		query := modifiedURL.Query()
		if ultimaPage == 1 {
			query.Set("lcp_page0", strconv.Itoa(ultimaPage))

		} else {

			query.Set("lcp_page0", strconv.Itoa(ultimaPage-2))
		}
		modifiedURL.RawQuery = query.Encode()
		this.NovelaRepository.UpdateUrlNovela(novela_id, modifiedURL.String())

		// go func() {
		// 	close(done)
		// }()

		// // Esperar a que todas las gorutinas hayan terminado y el canal de "done" se cierre
		// <-done
	}
	log.Println("Fin")
}
