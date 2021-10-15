# GPS Data Collector

GPS Data Collector is a microservice for gathering geolocation data.
A gRPC service built on [GoKit](https://gokit.io) that receives data from GPS devices saves it to the [PostgreSQL](https://www.postgresql.org) database. GPS device uses REST API.

## Installation

### Automated
Build image using Docker
```bash
docker build . -t gps_data_collector
```
Run Docker image
```bash
docker run gps_data_collector
```
>**NOTE:** Attached Dockerfile still has some issue. Check [Known issues]() section for a details.

### Manual
Use [protoc](https://grpc.io/docs/protoc-installation/) to recompile service and message .proto file *(optional)*
```bash
protoc --go_out=. --go-grpc_out=. api/v1/pb/gpsDataCollectorSvc.proto
```
Start database
```bash
docker-compose -f ./config/database/docker-compose.yaml up
```
Run an application
```bash
go run ./cmd/gpsDataCollector/main.go
```

## Usage
Checking microservice status
```bash
curl http://localhost:8081/health
{"status":200}
```
Sending data to server
```bash
curl -H "Content-Type: application/json" -X POST http://localhost:8081/addCoords -d '{"coordinates": {"latitude": 103.54, "longitude": -77.45}, "device_id": "handy gps"}' --user "user:password"
{"insertedId":"Success. ID: "}
```
>**NOTE:** The server has basic authentication middleware. Request with invalid credentials is always denied.

>**NOTE:** Database settings and authentication credentials are stored in the appropriate config files (/config/*.yaml)

## Known issues
### Testing
The main functionality has unit tests, but much of the code is not due to lack of time.\
Integration and E2E tests are still missing - they will be added later if time permits.
### Containerization
Dockerfile is practically almost ready, but still has a problem getting the docker daemon up in alpine. This is a mandatory step before starting the database. Consider replacing alpine with another alternative image.
### Other
Lack of:
* metrics collection
* better logging
* .dockerignore file
* .gitignore file
* other nice features :-)

## Project structure

**api**\
Stores the versions of the APIs swagger files and also the proto and pb files for the gRPC protobuf interface.

**cmd**\
Contains the entry point (main.go) files for all the services and also any other container images if any

**doc**\
This will contain the documentation for the project

**config**\
All the sample files or any specific configuration files should be stored here

**deploy**\
This directory will contain the deployment files used to deploy the application

**internal**\
This package is the conventional internal package identified by the Go compiler. It contains all the packages which need to be private and imported by its child directories and immediate parent directory. All the packages from this directory are common across the project

**pkg**\
This directory will have the complete executing code of all the services in separate packages.

**tests**\
It will have all the integration and E2E tests

**vendor**\
This directory stores all the third-party dependencies locally so that the version doesn’t mismatch later


## Contributing
Pull requests are welcome. For major changes, please open an issue first to discuss what you would like to change.

Please make sure to update tests as appropriate.

## License
[MIT](https://choosealicense.com/licenses/mit/)