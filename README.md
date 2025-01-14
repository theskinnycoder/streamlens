# Streamlens

Streamlens is a data platform that allows you to collect, store, and query data from your data in natural language.

## Tech Stack

- Go APIs using Chi Router
- PostgreSQL (for structured data)
- MongoDB (for unstructured data like logs)
- Kafka (for streaming data)
- Clickhouse (for real-time analytics)
- Redis (for caching, sessions, rate limiting etc.)
- Docker (for containerization)
- Docker Compose (for running services locally)
- Langchain (for natural language processing)
- OpenAI GPT-4o (for LLM)
- PGVector (for vector search)

## Services

### Auth

- Provides Authentication and Authorization for the platform using JWT, Cookies, Redis.
- Tech Stack:
  - Go Chi Router
  - Redis
  - JWT

### Ingestor

- Produces events to Kafka.
- Tech Stack:
  - Go
  - Kafka

### Processor

- Consumes, and enriches events from Kafka and stores them in Clickhouse.
- Adds metadata like field names, column names to MongoDB.
- Tech Stack:
  - Go
  - Kafka
  - Clickhouse
  - MongoDB

### Queryer (RAG)

- A RAG (Retrieval Augmented Generation) service that queries the data from Clickhouse and returns it to the user.
- Tech Stack:
  - Python
  - Clickhouse
  - Langchain with OpenAI GPT-4o as the LLM
  - PGVector (for vector search)
