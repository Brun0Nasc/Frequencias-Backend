package utils

import (
	"log"
	"strconv"

	sq "github.com/Masterminds/squirrel"
	"github.com/gin-gonic/gin"
)

// Constantes para identificar uma operação a ser realizada
const (
	FlagFilterEq = 1 << iota
	FlagFilterNotEq
	FlagFilterIn
	FlagFilterNotIn
	FlagFilterGt
	FlagFilterGtOrEq
	FlagFilterLt
	FlagFilterLtOrEq
)

type RequestParams struct {
	Filters map[string][]string
}

type Filter struct {
	name string
	data string
	flag int
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

func TransformFilterInQuery(p *RequestParams, builder *sq.SelectBuilder, filters ...Filter) {
	for i := range filters {
		for key, value := range p.Filters {
			if key == filters[i].name {
				switch filters[i].flag {
				case FlagFilterEq:
					*builder = builder.Where(sq.Eq{filters[i].data: value})
				case FlagFilterNotEq:
					*builder = builder.Where(sq.NotEq{filters[i].data: value})
				case FlagFilterIn:
					*builder = builder.Where(sq.Eq{filters[i].data: value})
				case FlagFilterNotIn:
					*builder = builder.Where(sq.NotEq{filters[i].data: value})
				case FlagFilterGt:
					*builder = builder.Where(sq.Gt{filters[i].data: value[0]})
				case FlagFilterGtOrEq:
					*builder = builder.Where(sq.GtOrEq{filters[i].data: value[0]})
				case FlagFilterLt:
					*builder = builder.Where(sq.Lt{filters[i].data: value[0]})
				case FlagFilterLtOrEq:
					*builder = builder.Where(sq.LtOrEq{filters[i].data: value[0]})
				}
			}
		}
	}
}

func BuildFilter(filterName, field string, flag int) (f Filter) {
	f = Filter{}
	f.name = filterName
	f.data = field
	f.flag = flag

	return
}

// AddFilter adiciona um filtro com chave e valores ao params
func (p *RequestParams) AddFilter(key string, values ...string) {
	p.Filters[key] = values
}
