  ## Overview ##
  
  This project exists to support other microservices written in Go. The project name is horrible, and will almost certainly change. Here is a breakdown of the current folder structure:

  * **database** - lightweight support for MSSQL [todo: rename/refactor to call out being mssql specific to support other RDBMS and nosql systems (postgres, bbolt, etc)]
  * **encryption** - support for encrypted values. Expects the appropriate .x509 cert and file with the password to be found at the path `/secrets/internal-encryption-certificate/`
  * **environment** - simple support for working with multiple environment files such as `.Dev.env` and `.Prod.env` 
  * **errors** - basic, reusable error definitions with the intent of capturing issues in processing data in a non-implementation specific way [not currently implemented in consuming microservice, but will be]
  * **grpc** - interceptors for easily and consistently adding cross-cutting functionality such as logging, tracing, and security within gRPC implementations
  * **logging** - basic, lightweight, simple support for integrating zap logging into splunk
  * **proto** - various useful bits for supporting protobuf
    * ***third_party*** - collection of common & useful `.proto` definitions. Meant to simplify build definitions
    * ***typehelpers*** - useful functions for translating to and from google's well-known types [WKTs]


### Upcoming/Planned functionality ### 
* support for more database systems
* extend out grpc interceptors
* some sort of eventing support (possibly with abstraction?) likely RabbitMQ, Azure Service Bus, nats.io
* centralized healthcheck/monitoring `.proto` definition?
* common bits to support rate limiting