# E-wallet

## How to start

Rename ".env_example" to ".env"

Build and run service
```yaml
docker-compose up --build
```
# Queries

## Send
```yaml
POST /api/send
```
Body
```yaml
{
  "from": string,
  "to": string,
  "amount": float
}
```

## GetLast
```yaml
GET /api/transactions
```
Body
```yaml
{
  "count" integer
}
```

## GetBalance
```yaml
GET /api/wallet/{address}/balance
```

## GetWallets
```yaml
GET /api/wallets
```
