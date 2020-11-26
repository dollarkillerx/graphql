package graphql

import (
	"encoding/json"
	"log"
	"strings"
	"testing"
)

func TestJson(t *testing.T) {
	p := H{"columnID": "columnID", "isDesc": true, "age": 18}
	marshal, err := json.Marshal(p)
	if err != nil {
		log.Fatalln(err)
	}
	log.Println(string(marshal))

	p = H{"columnID": "columnID"}

	pMarshal, err := Marshal(p)
	if err != nil {
		log.Fatalln(err)
	}
	log.Println(pMarshal)

	p = H{"columnID": "columnID", "isDesc": true, "age": 18}
	pMarshal, err = Marshal(p)
	if err != nil {
		log.Fatalln(err)
	}
	log.Println(pMarshal)

	p2 := []interface{}{
		p, p,
	}
	pMarshal, err = Marshal(p2)
	if err != nil {
		log.Fatalln(err)
	}
	log.Println(pMarshal)
}

var p = `query ivnestorsdsdsd {\n  InstitutionalInvestor(entityID: \"2000383468\", entityType: ORGANIZATION) {\n    exits(first: 1, orders: [{columnID: \"closed_on\", isDesc: true}]) {\n      totalCount\n    }\n  }\n}\n\nquery ivnestor2 {\n  InstitutionalInvestor(entityID: \"2000383468\", entityType: ORGANIZATION) {\n    funds(first: 10) {\n      totalCount\n      edges {\n        node {\n          fundTypes\n          identifier {\n            avatarURI\n            description\n          }\n          size {\n            display\n          }\n          dryPowder {\n            display\n          }\n        }\n      }\n    }\n  }\n}\n"}`

func TestSp(t *testing.T) {
	index := strings.Index(p, "query")
	index2 := strings.Index(p, "{")
	log.Println(strings.TrimSpace(p[index+5 : index2]))

	log.Println(getQueryID(p))
}

func TestSend(t *testing.T) {
	client := NewClient("http://xxxx:9998/api/graphql")
	body, err := client.NewRequest(`
	query ivnestor {
  InstitutionalInvestor(entityID: "$entityID", entityType: ORGANIZATION) {
    exits(first: 1, orders: [{columnID: "closed_on", isDesc: true}]) {
      totalCount
    }
  }
}

query xxxx {
  InstitutionalInvestor(entityID: "$entityID", entityType: ORGANIZATION) {
    funds(first: 10) {
      totalCount
      edges {
        node {
          fundTypes
          identifier {
            avatarURI
            description
          }
          size {
            display
          }
          dryPowder {
            display
          }
        }
      }
    }
  }
}
	`).Val("$entityID", "xxxx").
		Header("OAuth", "Bearer eyJ").Body()
	if err != nil {
		log.Fatalln(err)
	}
	log.Println(string(body))
}

func TestSend2(t *testing.T) {
	client := NewClient("http://xxxx:9998/api/graphql")

	// [{columnID: "closed_on", isDesc: true}]
	p := []interface{}{
		H{"columnID": "closed_on", "isDesc": true},
	}

	body, err := client.NewRequest(`
	query ivnestor {
  InstitutionalInvestor(entityID: "$entityID", entityType: ORGANIZATION) {
    exits(first: 1, orders: $p) {
      totalCount
    }
  }
}

query xxxx {
  InstitutionalInvestor(entityID: "$entityID", entityType: ORGANIZATION) {
    funds(first: 10) {
      totalCount
      edges {
        node {
          fundTypes
          identifier {
            avatarURI
            description
          }
          size {
            display
          }
          dryPowder {
            display
          }
        }
      }
    }
  }
}
	`).Val("$entityID", "xxx").Val("$p", p).
		Header("OAuth", "Bearer eyJhbGciOiJ").Body()
	if err != nil {
		log.Fatalln(err)
	}
	log.Println(string(body))
}
