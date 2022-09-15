# Flight task

## How to run

`
go run main.go
`

It will be run on endpoint:

http://localhost/flight


Send request by following example in json:

`
[["IND", "EWR"], ["SFO", "ATL"], ["GSO", "IND"], ["ATL", "GSO"]]
`

Response will be:

`
[
"SFO",
"EWR"
]
`

Also, you may try test, just run service_test.go