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


![Alt text](image/rabbit.png?raw=true "Architecture")

## Compile protobuf

```
# install protoc
brew install protobuf

# install go protobuf
go get github.com/golang/protobuf
go get github.com/golang/protobuf/protoc-gen-go

# complie protobuf
protoc spec/*.proto --go_out=:$GOPATH/src
```

## Deploy services

Change `host.docker.local` field in `.env` file to local machines ip or add
a host entry to `/etc/hosts` file by overriding `host.docker.local` with local
machines ip. In mac-os host.docker.local will route to localhost, so nothing to
change no mac-os.

```
docker-compose up -d rabbitmq
docker-compose up -d gateway
docker-compose up -d storage
```

## Test services

Gateway service expose document REST api on `/api/v1/documents:7654`. `JSON`
documents can be published to this api via `curl`.

```
// request
curl -v -XPOST "http://localhost:7654/api/v1/documents" \
--header "Content-Type: application/json" \
--data '
{
  "id": "1111",
  "name": "dot",
  "timestam": 89211232
}
'

// output
Note: Unnecessary use of -X or --request, POST is already inferred.
*   Trying ::1...
* TCP_NODELAY set
* Connected to localhost (::1) port 7654 (#0)
> POST /api/v1/documents HTTP/1.1
> Host: localhost:7654
> User-Agent: curl/7.54.0
> Accept: */*
> Content-Type: application/json
> Content-Length: 63
>
* upload completely sent off: 63 out of 63 bytes
< HTTP/1.1 201 Created
< Date: Sun, 20 Oct 2019 01:11:21 GMT
< Content-Length: 7
< Content-Type: text/plain; charset=utf-8
<
* Connection #0 to host localhost left intact
Created
```

Rabbitmq admin console can be access via `http://localhost:15672/#`. The queue 
informations, broker informations etc can be viewd from this admin console.

```
// rabbitmq queues can be viewed from here
//  username - admin
//  password - admin
http://localhost:15672/#/queues
```
