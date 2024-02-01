package endpointbuilder

import "strings"

type EndPointBuilder struct {
	baseUrl string
	paths   []string
	params  map[string]string
}

func SetBaseUrl(s string) EndPointBuilder {
	return EndPointBuilder{
		baseUrl: s,
		params:  make(map[string]string),
	}
}

func (epb EndPointBuilder) AddPath(k string) EndPointBuilder {
	epb.paths = append(epb.paths, k)
	return epb
}
func (ebp EndPointBuilder) AddParams(k string, v string) EndPointBuilder {
	ebp.params[k] = v

	return ebp
}
func (ebp EndPointBuilder) MakeUrl() string {
	var finalResult string
	finalResult = ebp.baseUrl + "/"
	for index, _ := range ebp.paths {
		finalResult = finalResult + ebp.paths[index] + "/"
	}
	finalResult = strings.TrimRight(finalResult, "/")
	finalResult = finalResult + "?"
	for index, _ := range ebp.params {
		finalResult = finalResult + index + "=" + ebp.params[index] + "&"
	}
	finalResult = strings.TrimRight(finalResult, "&")
	return finalResult
}
