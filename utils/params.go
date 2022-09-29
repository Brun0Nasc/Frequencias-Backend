package utils

import (
	"log"
	"strconv"

	"github.com/gin-gonic/gin"
)

type RequestParams struct {
	Filters map[string][]string
}

func ParseParams(c *gin.Context) (params *RequestParams) {
	params = &RequestParams{Filters: make(map[string][]string, 0)}

	for key, value := range c.Request.URL.Query() {
		params.Filters[key] = append(params.Filters[key], value...)
	}

	return
}

func (p *RequestParams) TemFiltro(key string) bool {
	if _, ok := p.Filters[key]; ok {
		return true
	}
	return false
}

// TemFiltroBool verifica se um existe um valor para o filtro informado e retorna o valor do filtro booleano e um OK para informar se o filtro existe
func (p *RequestParams) TemFiltroBool(key string) (bool, bool) {
	if value1, ok := p.Filters[key]; ok {
		value, err := strconv.ParseBool(value1[0])
		if err != nil {
			log.Fatalln(err)
		}

		return value, ok
	}

	return false, false
}
