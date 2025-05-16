


# Create accounts
curl -X POST -H "Content-Type: application/json" \
  -d '{"id":"A1", "name":"Alice", "balance":1000}' \
  http://localhost:8081/accounts

curl -X POST -H "Content-Type: application/json" \
  -d '{"id":"A2", "name":"Abrar", "balance":1000}' \
  http://localhost:8081/accounts  

# Get account
curl http://localhost:8081/accounts/A1
curl http://localhost:8081/accounts/A2


# Debit
curl -X POST -H "Content-Type: application/json" \
  -d '{"amount":200}' \
  http://localhost:8081/accounts/A1/debit


# Credit
curl -X POST -H "Content-Type: application/json" \
  -d '{"amount":300}' \
  http://localhost:8081/accounts/A1/credit
