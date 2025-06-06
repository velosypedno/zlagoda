# ZLAGODA

## Install

### Dependencies

Ensure you have the following installed:

- [Go](https://golang.org/doc/install) (>= 1.23.5)
- [Docker](https://docs.docker.com/get-docker/)

### Steps

1. **Clone the repository**:

   ```bash
   git clone git@github.com:velosypedno/zlagoda.git
   ```

2. **Change work directory**:

    ```bash
    cd zlagoda
    ```

3. **Configure environmental variables**:

    Copy `.env.sample`

    ```bash
    cp .env.sample .env
    ```

4. **Build and up services by Docker Compose**:

    ```bash
    docker compose up
    ```

### To Run api without docker

1. **Clone the repository**:

   ```bash
   git clone git@github.com:velosypedno/zlagoda.git
   ```

2. **Change work directory**:

    ```bash
    cd zlagoda
    ```

3. **Configure environmental variables**:

    Copy `.env.sample`

    ```bash
    cp .env.sample .env
    ```

    Change `DB_HOST=localhost` in `.env`

4. **Build and up services by Docker Compose**:

    ```bash
    docker compose up postgres-zlagoda migrator-zlagoda
    ```

5. **Build binary with go**:

    ```bash
    go run cmd/api/*
    ```

6. **WUALYA**
