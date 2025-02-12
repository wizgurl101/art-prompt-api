# Art Prompt API

## Dev Setup

### Add .env file with the following variables:

- MONGODB_URI={get from mongo db}
- OLLAMA_URL={host:port_number}

Note: If there is a connection issue with Ollama or Redis service
with Docker Compose, use host.docker.internal instead of localhost

### Run Docker Compose Build

```bash
    docker-compose up --build
```

### If the containers been build, use the following commands to run the API

```bash
    docker-compose up backend
```

```bash
    docker-compose up ollama
```
