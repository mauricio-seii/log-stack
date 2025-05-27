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

By placing a .log file in the `/log-stack/logs` directory, it should be parsed by the Go application and sent to Elasticsearch. Parsed logs can be visualized with Kibana at `http://localhost:5601`.

Ideally, you should run the following command and if you want `env` and `service` keywords in your Elasticsearch indexes, then follow the pattern `service-name.env.log` and Logstash will extract those for you making it a bit easier to filter/search logs later on.

```bash
# assuming you are already authenticated and can fetch containers logs
kubectl -n dev logs deploy/deposit-bitgo > ~/Downloads/deposit-bitgo.dev.log

# now copy the log file into the /logs directory where the project was cloned to
cp ~/Downloads/deposit-bitgo.dev.log ~/github/log-stack/logs
```
