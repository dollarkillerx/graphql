# graphql
simple graphql client

exp:
```go
client := graphql.NewClient("http://xxxx:9998/api/graphql")

	// [{columnID: "closed_on", isDesc: true}]
	p := []interface{}{
		graphql.H{"columnID": "closed_on", "isDesc": true},
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

exp2:
```go
    client := graphql.NewClient("http://xxxx:9998/api/graphql")
	client = client.Debug()

	token := "Bearer xxxxx"

	client = client.NewRequest(`
		query InstitutionalInvestor(
		  $entityID: String!
			$entityType: EntityType!
		  $first: MaybeInt32
		  $after: MaybeString
		  $last: MaybeInt32
		  $before: MaybeString
		  $orders: [OrderColumn]
		) {
		   InstitutionalInvestor(
			entityID:  $entityID
			entityType: $entityType
		  ){
			investments(
			  first: $first
			  after: $after
			  last: $last
			  before: $before
			  orders: $orders
			) {
			  totalCount
			  pageInfo{
				hasPreviousPage
				hasNextPage
				startCursor
				endCursor
			  }
			  edges{
				cursor
				node{
				  closedOn
				  investedAmount{
					display
				  }
				  equityPercentage{
					display
				  }
				  relatedDeal{
					identifier{
					  entity{
						entityID
					  }
					  description
					}
					dealType
				  }
				  fundedOrg{
					identifier{
					  entity{
						entityID
					  }
					  description
					}
				  }
				  otherInvestors{
					identifier{
					  entity{
						entityID
					  }
					  description
					}
					amount{
					  amountIn10K
					}
				  }
				}
			  }
			}
		  }
		}
	`)

	client.SetVariables(map[string]interface{}{
		"entityID":   "2000435437",
		"entityType": "ORGANIZATION",
		"first":      10,
		"orders": []interface{}{
			map[string]interface{}{
				"columnID": "closed_on",
				"isDesc":   true,
			},
		},
	})


	body, err := client.
		Header("OAuth", token).
		Body()
	if err != nil {
		log.Fatalln(err)
	}
	log.Println(string(body))
```