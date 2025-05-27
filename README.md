# Local Log Parsing Stack

This project sets up a local log parsing stack using:
- A Go parser app that reads Kubernetes log files
- Logstash for log ingestion and transformation
- Elasticsearch for indexing logs
- Kibana for visualization

---

## 🚀 Getting Started

### 🔧 Requirements

- [Docker](https://docs.docker.com/get-docker/)
- [Docker Compose](https://docs.docker.com/compose/install/)

---

## 📦 Run with Docker Compose

To run the full stack including the parser, Logstash, Elasticsearch, and Kibana:

| Command      | Description                    |
| ------------ | ------------------------------ |
| `make up`    | Build and start all containers |
| `make down`  | Stop and remove containers     |
| `make build` | Build the Go parser binary     |


```bash
make up
```

The above command will:
* Builds the Go parser
* Sets up volumes and directories
* Brings up all services with Docker Compose

## Intended use

This stack was created in response to lack of tools to check logs while in development phase. 

### how to use

* Place a .log file into `/log-stack/logs` directory
* It should be parsed by the Go application and sent to Elasticsearch
* Parsed logs can be visualized with Kibana at `http://localhost:5601`

*side note: follow this standard when creating the .log file: `service-name.env.log` this will add keywords into the Elasticsearch index*

```bash
# assuming you are already authenticated and can fetch containers logs
kubectl -n dev logs deploy/my-service > ~/Downloads/my-service.dev.log

# now copy the log file into the /logs directory where the project was cloned to
cp ~/Downloads/deposit-bitgo.dev.log ~/github/log-stack/logs
```
