# cached-api

Simple Go-based API service that fetches data from an external API and caches the results using a Valkey. The service is designed to improve performance by reducing redundant external API calls through caching.

## Getting Started

To run the project, ensure the following tools are installed:

- **Go 1.23+** ([official documentation](https://go.dev/))
- **Valkey Database** ([official documentation](https://valkey.io/))

### Configuration

The application uses default configuration values but allows customization via a `.env` file. Use the following structure:

```.env
CACHE_URL=<Cache host>
CACHE_PASSWORD=<Cache password>
CACHE_DATABASE=<Cache database>
CACHE_TTL=<Cache TTL in seconds>
```

## How it works

- Fetches data from an external API (https://jsonplaceholder.typicode.com/todos/{id})
- Caches responses in a Valkey key-value store to optimize performance
- Implements a cache-first strategy to serve data quickly when available

### Workflow Diagram

```bash
Client
  |
  v
GET /todo/{id}
  |
  v
+-------------------+
|   ApiHandler      |
| - Check Cache     |
| - Fetch if Miss   |
+-------------------+
  |
  v
+-------------------+
|   ApiService      |
| - Cache Logic     |
| - External API    |
+-------------------+
  |
  v
+-------------------+
|   Valkey Cache    |
| - Get/Set Data    |
+-------------------+
```