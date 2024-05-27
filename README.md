# Route Server

Route Server is a microservice responsible for managing the registration of domains and endpoints stored in a database. It provides data for routing configuration and service access.

### Functionality

- Manages registration of domains and endpoints in a PostgreSQL database.
- Generates and updates routing configuration for other services.
- Interacts with the database to fetch current data.

### Interaction with Other Microservices

- **xds**: Route Server provides current domain and endpoint data via gRPC requests.
- **cert-server**: Can interact with Cert Server to obtain TLS certificates.

### Technologies

- Written in Go.
- Uses gRPC for communication with other microservices.
- Interacts with PostgreSQL for data storage.