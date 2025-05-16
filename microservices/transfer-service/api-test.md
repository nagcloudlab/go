


### Transfer Service API Test
curl -X POST http://localhost:8082/transfer \
  -H "Content-Type: application/json" \
  -d '{
    "from_account": "A1",
    "to_account": "A2",
    "amount": 200
}'

curl -X POST http://localhost:8082/transfer \
  -H "Content-Type: application/json" \
  -d '{
    "from_account": "A2",
    "to_account": "A1",
    "amount": 200
}'