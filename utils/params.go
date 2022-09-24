package utils

import "github.com/gin-gonic/gin"

type RequestParams struct {
	Filters map[string][]string
}

func ParseParams(c *gin.Context) (params *RequestParams) {
	params = &RequestParams{Filters: make(map[string][]string, 0)}
	
	for i := range c.Params {
		key := c.Params[i].Key
		value := c.Params[i].Value

		params.Filters[key] = append(params.Filters[key], value)
	}

	return
}

func (p *RequestParams) TemFiltro(key string) bool {
	if _, ok := p.Filters[key]; ok {
		return true
	}
	return false
}