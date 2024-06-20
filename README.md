# Online Store

## Description

This is an online store application with functionalities such as viewing products by category, adding products to the cart, viewing cart items, removing items from the cart, and checking out.

## Technologies

- Gin (Web Framework)
- GORM (ORM)
- PostgreSQL (Database)
- Redis (Cache)
- Docker
- Docker Compose

## Setup

### Prerequisites

- Go 1.16+
- Docker
- Docker Compose

### Steps

1. Clone the repository
git clone https://github.com/yourusername/online-store.git
cd online-store

2. Create `.env` file
#PostgreSQL
export DB_HOST=your_db_host
export DB_USER=your_db_user
export DB_PASSWORD=your_db_password
export DB_NAME=your_db_name
export DB_PORT=your_db_port
#Redis
REDIS_HOST=your_redis_host
REDIS_PORT=your_redis_port
REDIS_PASSWORD=your_redis_password
#JWT Secret Key
JWT_SECRET_KEY=your_jwt_secret_key_here

3. Run Docker Compose
docker-compose up --build

4. Access the application at `http://localhost:8080`

## API Endpoints

- `POST /cart/add`: Add product to cart
- `GET /cart/`: View cart items
- `DELETE /cart/remove/:product_id`: Remove product from cart
- `DELETE /cart/clear`: Clear cart
- `POST /order/checkout`: Checkout
- `GET /products/category/:category_id`: Get products by category

## Deployment

### Docker

1. Build the Docker image
docker build -t yourusername/online-store .

2. Push to Docker Hub
docker push yourusername/online-store

3. Deploy to a cloud provider (e.g., Heroku, AWS, GCP)