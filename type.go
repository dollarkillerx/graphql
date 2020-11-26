package graphql

import (
	"encoding/json"
	"fmt"
	"strings"
)

type H map[string]interface{}

func Marshal(h interface{}) (string, error) {
	marshal, err := json.Marshal(h)
	if err != nil {
		return "", err
	}
	var resp string

	sp1 := strings.Split(string(marshal), ",")
	for _, v := range sp1 {
		sp2 := strings.Split(v, ":")
		if len(sp2) == 2 {
			replace := strings.Replace(sp2[0], `"`, "", -1)
			resp += replace + ":"
			resp += sp2[1] + ","
		}
	}

	if len(resp) != 0 {
		return resp[:len(resp)-1], nil
	}
	return resp, fmt.Errorf("marshal error")
}

func getQueryID(scheam string) string {
	i1 := strings.Index(scheam, "query")
	i2 := strings.Index(scheam, "{")

	return strings.TrimSpace(scheam[i1+5 : i2])
}
