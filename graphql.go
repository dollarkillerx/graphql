package graphql

import (
	"encoding/json"
	"github.com/dollarkillerx/urllib"
	"log"
	"strconv"
	"strings"
)

type ReqData struct {
	OperationName string   `json:"operationName"`
	Variables     []string `json:"variables"`
	Query         string   `json:"query"`
}

type Client struct {
	addr string

	name           string // search name
	scheam         string
	keys           map[string]interface{}
	keysFormatting map[string]string
	header         map[string]string

	err  error
	data []byte
}

func NewClient(addr string) *Client {
	return &Client{
		addr:           addr,
		keys:           map[string]interface{}{},
		keysFormatting: map[string]string{},
		header:         map[string]string{},
		data:           []byte{},
	}
}

func (g *Client) NewRequest(schema string) *Client {
	g.scheam = schema
	return g
}

func (g *Client) Val(key string, val interface{}) *Client {
	g.keys[key] = val
	return g
}

func (g *Client) Header(key string, val string) *Client {
	g.header[key] = val
	return g
}

func (g *Client) formatting() {
	// str
	// int
	// float
	// => to string

	// struct
	// slice
	// map
	// => to json

	for k, v := range g.keys {
		switch v.(type) {
		case string:
			g.keysFormatting[k] = v.(string)
		case int32:
			g.keysFormatting[k] = strconv.Itoa(int(v.(int32)))
		case int:
			g.keysFormatting[k] = strconv.Itoa(v.(int))
		case int64:
			g.keysFormatting[k] = strconv.Itoa(int(v.(int64)))
		case float64:
			g.keysFormatting[k] = strconv.FormatFloat(v.(float64), 'E', -1, 64)
		case float32:
			g.keysFormatting[k] = strconv.FormatFloat(float64(v.(float32)), 'E', -1, 32)
		default:
			marshal, err := Marshal(v)
			if err != nil {
				log.Println(err)
				continue
			}
			g.keysFormatting[k] = marshal
		}
	}

	for k, v := range g.keysFormatting {
		g.scheam = strings.ReplaceAll(g.scheam, k, v)
	}

}

func (g *Client) send() *Client {
	g.formatting()
	req := ReqData{
		OperationName: getQueryID(g.scheam),
		Query:         g.scheam,
	}

	marshal, err := json.Marshal(req)
	if err != nil {
		g.err = err
		return g
	}

	log.Println(string(marshal))
	base := urllib.Post(g.addr).
		SetHeader("Cache-Control", "no-cache").
		SetHeader("Content-Type", "application/json")

	for k, v := range g.header {
		base = base.SetHeader(k, v)
	}

	httpCode, body, err := base.SetJson(marshal).ByteRetry(3)
	if err != nil {
		g.err = err
		return g
	}

	g.data = body
	if httpCode != 200 {
		log.Printf("HTTPCODE: %d resp: %s \n", httpCode, string(body))
	}

	return g
}

func (g *Client) Body() ([]byte, error) {
	g.send()
	return g.data, g.err
}

func (g *Client) BindJson(i interface{}) error {
	g.send()
	return json.Unmarshal(g.data, i)
}
