# Compose

`docker-compose` deployment of microservices based document management service.

## Services

There are two main services in the system, `Gateway` and `Storage`. The services 
built built with `Golang` by using `Rabbitmq` and `Protobuf`. Gateway service 
act as microservices api-gateway. Storage service do actual document management 
functions(ex: create documents in database). `Protobuf` use to serialized structured 
data(request response between microservices). `Rabbitmq` uses as message broker.

Client submit `JSON` documents to gateway service via REST api. Gateway service 
create `Protobuf` message based on the `JSON` data and pass it to storage service 
via `Rabbitmq`. Then storage service do what ever the operation(ex: create documents 
on a database) and sends the status back to gateway service as `Protobuf` message. 
Finally gateway service convert `Protobuf` message which receives from storage 
service into `JSON` object and send it back to client. Gateway service listen to 
Rabbitmq `gateway` queue and Storage service listen to `storage` queue.

