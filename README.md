# graphql
simple graphql client

exp:
```go
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
	`).Val("$entityID", "xxxx").Val("$p", p).
		Header("OAuth", "Bearer eyJhbGciOiJ").Body()
	if err != nil {
		log.Fatalln(err)
	}
	log.Println(string(body))
```