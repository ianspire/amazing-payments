# amazing-payments
## About this project
A golang project for integrating with Stripe and capturing payments

## Prerequisites
#### Golang
This project was built using golang 1.13, and requires an install of golang to build the source.
Earlier versions of golang may work (such as 1.11), however you may need to set the `GO111MODULE`
environment variable, and/or clone the repository outside of your `GOPATH` to get the desired results.
See here for more details: https://github.com/golang/go/wiki/Modules#gomod

#### Postgres
The service also requires a Postgres database.

#### Stripe
A valid Stripe account should already be setup, and you will need to provide the API key as described
in **Configuration** below.

#### Python, Yoyo
Database migrations are managed with [yoyo](https://pypi.org/project/yoyo-migrations/) (a Python package).
Python will need to be installed, as well as the Yoyo package in order to run the migrations.


## Configuration
The service is configured using environment variables, which are specified in `pkg/config.go` file.
Defaults are already specified in the Config struct tags.  Feel free to provide overrides to the defaults
via environment variables though.

#### Stripe Client
The one required config variable that doesn't have a default configuration is `STRIPE_KEY`.  You will
need to provide a Stripe API key in order to run the service.

#### Yoyo
The file at `db/yoyo.ini` does not come with a default configuration.  Please set the variables to the
appropriate values to configure the database.

#### Service Ports
Default ports for the payment service are:
`8080`: grpc-gateway (where HTTP/1.1 calls can be made using REST)
`8081`: grpc (where HTTP/2 RPC calls can be made)

However, these ports can be configured using the following environment variables.
`JSON_PORT`: grpc-gateway
`RPC_PORT` : grpc


## Installation and Execution
1. Clone this repository
2. Run `make` at the root of the repository
3. Ensure `db/yoyo.ini` has been properly configured, `cd` into the `db` directory, and run `yoyo apply ./migrations`
3. `cd` back to the repository root, export `STRIPE_KEY` environment variable (and any other overrides to the defaults)
4. Run `go run cmd/amazing-payments/main.go` to start the service


## API Documentation
This project uses [Protocol Buffers](https://developers.google.com/protocol-buffers) to define the API, which
help generate some of the code used to produce the service and the request handlers.  It also provides a plugin to
produce [swagger docs](https://swagger.io/tools/swagger-ui/download/), which can be used to provide easy-to-use
API documentation and tools for manual testing.

Currently there are only three methods available, so we can provide information on them here:

| `GET` | `/v1/healthcheck` | 
Description: Pings the database, and returns true if the connection is still active
Parameters: None
Response: 

| `POST` | `/v1/customer` | 
Description:
Parameters:
Response:

| `GET` | `/v1/customer/{customerID}` | 
Description:
Parameters:
Response:

More methods are planned to fulfill the requirements of the project, as well:
| `POST` | `/v1/product` |


## System Architecture

<img src="https://imgur.com/a/F2SLJ5F" width="480">

The intended architecture is to provision the payment service either in AWS EC2 instances that scale with an
Auto-Scaling Group, or in a managed container orchestration tool such as ECS (Docker) or EKS (Kubernetes)
which can scale the service as necessary.  The service is provided behind a generic AWS ELB for network
load balancing.  Each service can connect to another cluster (either EC2, ECS or EKS again) that hosts
pgbouncer or pgpool II as outlined in this [article](https://aws.amazon.com/blogs/database/a-single-pgpool-endpoint-for-reads-and-writes-with-amazon-aurora-postgresql/)

The Postgres load balancer can provide a single endpoint for both read and write connections, allowing
for some smart optimization using segregation of read and write requests when transaction isolation is less
of a concern.

Finally, messaging will be sent from the service to an SQS queue, where a lambda can pick up
the messages to be sent to an email service provider, allowing customers to receive notifications as
payments are sent.  In case of failure, using a re-drive policy, the SQS Queue can be backed by a deadletter queue
where the messages can be stored for future review.
