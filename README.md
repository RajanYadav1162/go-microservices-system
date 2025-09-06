# Go Micro-Masterclass 🎟️

> Build a **production-grade, event-driven microservices stack** in a single weekend – **zero cost, zero ops, maximum résumé impact**.  

[![Go](https://img.shields.io/badge/Go-1.23-00ADD8?style=flat&logo=go)](https://golang.org/) 
[![License](https://img.shields.io/badge/License-MIT-green.svg)](https://opensource.org/licenses/MIT)

---

## 📊 Current State (Step 5 ✅)

| Feature             | Tool                       |
|---------------------|----------------------------|
| REST Edge           | [Gin](https://github.com/gin-gonic/gin) |
| JSON Logs           | [Zerolog](https://github.com/rs/zerolog) |
| OpenAPI Docs        | [Swagger UI](https://swagger.io/tools/swagger-ui/) |
| Database            | [PostgreSQL](https://www.postgresql.org/) + [GORM](https://gorm.io/) (auto-migrate) |
| Events              | [NATS JetStream](https://nats.io/) (publisher side) |
| CI                  | [GitHub Actions](https://github.com/features/actions) (coming soon) |

---

## 🚀 Quick Start (Local Setup)

### Prerequisites
- [Docker](https://www.docker.com/) (for PostgreSQL and NATS)
- [Go 1.23](https://golang.org/dl/) or later

### Steps
1. **Clone the Repository**
   ```bash
   git clone https://github.com/YOUR_GITHUB/go-micro-masterclass.git
   cd go-micro-masterclass
   ```

2. **Start Infrastructure**
   ```bash
   # PostgreSQL
   docker run -d --name postgres-dev -p 5432:5432 \
     -e POSTGRES_USER=tickets -e POSTGRES_PASSWORD=tickets -e POSTGRES_DB=ticketsdb \
     postgres:16-alpine

   # NATS JetStream
   docker run -d --name nats-dev -p 4222:4222 -p 8222:8222 \
     nats:2.10-alpine -js
   ```

3. **Run the API**
   ```bash
   go mod tidy
   go run cmd/api/main.go
   ```

4. **Test the API**
   ```bash
   curl -X POST localhost:8080/orders \
     -H "Content-Type: application/json" \
     -d '{"user_id":"u1","concert_id":"c1","qty":2,"amount":99.99}'
   ```

5. **Browse API Docs**
   Open [http://localhost:8080/swagger/index.html](http://localhost:8080/swagger/index.html) in your browser.

---

## 📂 Project Layout

```
cmd/api/           # Monolith entry point (will split into services soon)
internal/
  model/           # GORM entities
  db/              # Database connection + migrations
  event/           # Domain events (DTOs)
  msg/             # NATS JetStream publisher
docs/              # Swagger JSON/YAML (auto-generated)
tools/             # Local binaries (air, swag) – no global installs needed
```

---

## 🛤️ Roadmap (Next Steps)

| Step | Target                                      | Status |
|------|---------------------------------------------|--------|
| 6    | Consumer microservice `svc-payment` (reacts to `ORDERS.created`) | 🚧 In Progress |
| 7    | Circuit-breaker & retry ([Sentinel](https://github.com/alibaba/Sentinel)) | ⏳ Planned |
| 8    | Service discovery ([Consul](https://www.consul.io/)) | ⏳ Planned |
| 9    | Tracing ([OpenTelemetry](https://opentelemetry.io/)) + Grafana stack | ⏳ Planned |
| 10   | Split into Kubernetes manifests & [Skaffold](https://skaffold.dev/) dev loop | ⏳ Planned |
| 11   | GitHub Actions CI → build, test, push to [ghcr.io](https://ghcr.io) | ⏳ Planned |
| 12   | Add `svc-inventory` & `svc-notification` (full flow) | ⏳ Planned |

---

## 🤔 Why NATS JetStream Instead of Kafka?

- **Lightweight**: Single 40 MB binary vs. JVM + ZooKeeper.
- **Durable & Replayable**: Same durability and replay capabilities with 1% of the operational complexity.
- **Go-Friendly**: Idiomatic Go client with one-line publishing.
- **Perfect for Weekend Projects**: Enterprise-grade features with minimal setup.

---

## 🛠️ Tech Stack (All Free/Open-Source)

- **Go 1.23** – Fast, simple, and modern.
- **Gin** – Lightweight HTTP framework.
- **GORM** – ORM for PostgreSQL.
- **PostgreSQL** – Robust relational database.
- **NATS JetStream** – Event streaming and messaging.
- **Swagger** – API documentation.
- **Zerolog** – Fast JSON logging.
- **Docker** – Containerized infrastructure.
- **GitHub Actions** – CI/CD (coming soon).
- **k3d** – Lightweight Kubernetes (coming soon).

---

## 🙌 Contribute

Feel free to open issues or submit PRs to help build the rest of this project! Let’s create something awesome together. 🚀

> **Note**: Replace `YOUR_GITHUB` in the clone URL with your GitHub username.
