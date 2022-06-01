# Go-Commerce

Simple e-commerce server created using go language

# Guides

## Deployment in Local

- Install Go https://go.dev/doc/install
- Copy `config/example.yml` to `config/local.yml` and modify the values accordingly
- If you aren't setup yet database, use `docker compose up` to set postgresql db on your docker
- Run these commands

```
go mod download # get package dependecy

go run main.go -env=local (flag is base yaml name you used)
```

- now the rest api is available at `localhost:8080/api/v1` (unless you change the port)

## Rest API brief

1. For Buyer & Seller
   - Login: `POST` - `/authentications/login`
2. For Seller
   - View product list: `GET` - `/products/seller`
   - Add new product: `POST` - `/products`
   - View order list: `GET` - `/orders`
   - Accept order: `PATCH` - `/orders/accept/:id`
3. For Buyer
   - View list of products: `GET` - `/products`
   - Order product: `POST` - `/orders`
   - View order lists: `GET` - `/orders`

for detail usage, import collection and its environment from folder `docs` to your `postman`
