# Documentation


## Installation

to instal package use vendor

```bash
go mod vendor
```

## How To Run It

to started it

```bash
go run ./cmd/main.go
```

# Explanation
### a and b answer.
login created with jwt and md5 encryption for password
jwt will expired in 15 minutes

to login use api 


```
{baseUrl}/login
```

with body 
```
{
  "username" : "admin1"
  "password" : "admin1"
}
```

curl for postman
```
curl --location --request POST 'http://0.0.0.0:14045/api/v1/login' \
--header 'Content-Type: application/json' \
--data-raw '{
    "user_name" : "admin1",
    "password" : "admin1"
}'
```

### c, d, e, f answer.
report will only show to user that login with jwt an

report can access in api

```
{baseUrl}/report
```
with body 
```
{
  "month" : "",
  "page"  : "",
  "limit" : ""

}
```
header
```
Authirization : Bearer {access_token}
```
access token is provided after login in login responses 

### curl for postman 

```
curl --location --request POST 'http://0.0.0.0:14045/api/v1/report' \
--header 'Authorization: Bearer {access_token}' \
--header 'Content-Type: application/json' \
--data-raw '{
    "month" : "11",
    "limit" : "2",
    "page" : "2"
}'
```

## g answer
that is optimal erd because that is included to normal data without redundant field in each table

## h answer
query for the report  : 

```
select merchants.merchant_name,transactions.id,sum(transactions.bill_total) as bill_total" 
		group by transactions.created_at.
		left join merchants on users.id = merchants.user_id"
		left join transactions on merchants.id = transactions.merchant_id"
		users.id = 1 where month(transactions.updated_at) = "11" limit 10 , 1
```

## number 3 answer
![answer number 3](https://github.com/prasetiyo28/test-case-majoo/blob/main/nomer4.jpg)

link
[pseudocode](https://go.dev/play/p/Yjqwzs2BBq2)

## number 4 answer

link
[pseudocode](https://go.dev/play/p/uRirUSNnaUX)