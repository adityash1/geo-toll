A distributed microservice-based system for calculating road tolls based on geolocation data.

#### Architecture

<img src="https://github.com/user-attachments/assets/a84c0bba-3786-488f-aff0-7c3d66300d85" width=500 />

GEO-TOLL uses a modern event-driven microservice architecture:

- **On-Board Unit (OBU)**: Vehicle-mounted device that transmits geolocation data
- **Receiver Service**: Ingests real-time location data from OBUs and publishes to Kafka
- **Distance Calculator Service**: Processes geolocation events to calculate distances traveled
- **Invoice Service**: Manages toll invoicing with DB persistence
- **Invoice Calculator Service**: Computes toll charges based on distance, vehicle type, and toll rates
- **API Gateway**: Unified entry point for client applications

#### Technical Stack

- **Event Streaming**: Kafka for reliable, scalable event processing
- **Microservices**: Independently deployable services with clear domain boundaries
- **Database**: Persistent storage for invoice and toll data
- **RESTful APIs**: Communication between services and external clients
- **Containerization**: Docker for consistent deployment across environments

#### Key Features

- Real-time geolocation data processing
- Scalable event-driven architecture
- Fault-tolerant design with message persistence
- Accurate distance calculation algorithms
- Flexible toll calculation based on multiple parameters

#### Implementation Highlights

- Asynchronous communication patterns for high throughput
- Decoupled services for independent scaling and deployment
- Event sourcing for reliable data processing
- Comprehensive API for integration with external systems
