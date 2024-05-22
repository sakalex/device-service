# BlockRockSouth Test Assignment
This is a test assignment for the Senior Golang Engineer position at [BlockRockSouth](https://blackrocksouth.com/).

## Task Description


```
Нужно разработать приложение, которое через веб предоставляет список 
устройств в формате protobuf, включая тип (мышь, клавиатура, прочее), 
device path, vendorid, productid.
```


## Clarification Questions and Answers
|Questions | Answers |
|----------|---------|
| __1 разработать протобаф схему__ |  |
|1.1 Тип устройства должен быть перечислением или текстовым полем? | Перечисление |
|1.2 device path - строка или структура описывающая пусть (если структура то какие поля в ней пресутствуют) | трока |
|1.3 vendorid, productid - целые числа или uuid | uuid |
|1.4 какой должен быть канал одиночный или стримы, если стримы то в каком направлении | одиночный |
|1.5 Должны ли быть пагинация или другие служебные данные в теле запроса | нет |
| __2 Реализовать протобаф сервер с поддержкой этого api__ | |
|2.1 где должны храниться данные которые будут отдавться по protobuf api ?  файл? база данных? память? | память |
|2.2 Должно ли программа как-то особенно обрабатывать ошибки? | нет |
|2.3 надо ли контейнеризировать приложение? | нет |
|2.4 Ожидается ли покрытие тестами логики приложения в рамках тестового задания? | нет |
| В каком виде ожидается сдача задания? гит репозитои? zip архив? иное?| В виде гита |

## Project Structure

```
.
|-- Dockerfile
|-- Makefile
|-- README.md
|-- cli
|   |-- device_consumer
|   |   `-- main.go
|   `-- device_provider
|       `-- main.go
|-- go.mod
|-- go.sum
|-- proto
|   `-- device-service.proto
`-- src
    |-- devices
    |   |-- provider.go
    |   `-- service.go
    |-- proto-gen
    |   `-- proto
    |       |-- device-service.pb.go
    |       `-- device-service_grpc.pb.go
    `-- repos
        |-- in_memory_repos.go
        `-- repos.go
```
- [__Dockerfile__](./Dockerfile): Dockerfile for dev container.
- [__Makefile__](./Makefile): The makefile for building and running project.
- [__README.md__](./README.md): The README file for the project, which contains instructions and other information.
- [__cli/device_consumer/main.go__](./cli/device_consumer/main.go): The entry point for the device consumer CLI client.
- [__cli/device_provider/main.go__](./cli/device_provider/main.go): The entry point for the device provider server.
- __go.mod and go.sum__: The Go module files.
- [__proto/device-service.proto__](./proto/device-service.proto): The Protocol Buffers file defining the gRPC service.
- [__src/devices/provider.go__](./src/devices/provider.go): The implementation of the gRPC service.
- [__src/devices/service.go__](./src/devices/service.go): The server logic for starting the gRPC service.
- [__src/proto-gen/proto/device-service.pb.go and src/proto-gen/proto/device-service_grpc.pb.go__](./src/proto-gen/proto/): The Go code generated from the .proto file.
- [__src/repos/in_memory_repos.go__](./src/repos/in_memory_repos.go): The implementation of the device repository using an in-memory data store.
- [__src/repos/repos.go__](./src/repos/repos.go): The interface definition for the device repository.

## Usefull commands
1. Run server
    ```bash
    make run-server &
    ```
2. Run client with list command
    ```bash
    make list-device
    ```
3. cli client help
    ```bash
    go run ./cli/device_consumer/main.go -address localhost:50051 -help
    ```
4. Regenerate protobuf
    ```bash
    make proto
    ```

## Example 
```bash
$ make run-server &
[1] 73766
go run ./cli/device_provider/main.go -port 50051 -logtostderr
I0522 07:04:08.629708   73885 service.go:21] Starting the server on port 50051...
I0522 07:04:08.629882   73885 service.go:42] Server is running on port 50051

$ 
$ make list-device
go run ./cli/device_consumer/main.go -address localhost:50051 list
I0522 07:04:17.325361   73885 service.go:35] Request: /device_service.DeviceProviderService/ListDevices, Time: 1.8µs, Response: devices:{id:"mouse_1"  type:MOUSE  device_path:"/dev/input/mouse0"  vendor_id:"uuid1"  product_id:"uuid2"}  devices:{id:"keyboard_1"  type:KEYBOARD  device_path:"/dev/input/keyboard1"  vendor_id:"uuid3"  product_id:"uuid4"}
2024/05/22 07:04:17 Device: id:"mouse_1"  type:MOUSE  device_path:"/dev/input/mouse0"  vendor_id:"uuid1"  product_id:"uuid2"
2024/05/22 07:04:17 Device: id:"keyboard_1" type:KEYBOARD device_path:"/dev/input/keyboard1" vendor_id:"uuid3" product_id:"uuid4"
$ 
$ kill %1
$ make: *** [Makefile:10: run-server] Terminated

[1]+  Terminated              make run-server
$
```

### Client
```bash
$ go run ./cli/device_consumer/main.go -help
Usage of /tmp/go-build3458314594/b001/exe/main:
  -address string
        The server address (default "localhost:50051")
  -help
        Display help

Commands:
  list    List all devices
  add     Add a new device
  delete  Delete a device

Usage of list:
Usage of add:
  -id string
        The device id (default "d3c1defb-e595-4c44-9a11-aefa6c7a03bf")
  -path string
        The device path
  -product_id string
        The product id (default "7f7edc22-c9e6-4eae-83b6-bd84fa9fec9b")
  -type value
        The device type (MOUSE, KEYBOARD, OTHER). Default is MOUSE. (default MOUSE)
  -vendor_id string
        The vendor id (default "3c9e26a1-f019-40af-b500-3890238d93f7")
Usage of delete:
  -id string
        The device id (default "de0f79af-d0fc-487f-a59c-a4bd2c6b4dc8")
```

1. Create
```bash
$ go run ./cli/device_consumer/main.go -address localhost:50051 add
I0522 09:01:20.541698    6444 service.go:35] gRPC Request: /device_service.DeviceProviderService/AddDevice, Time: 1.4µs, Response: success:true
$ go run ./cli/device_consumer/main.go -address localhost:50051 list
I0522 09:01:34.387775    6444 service.go:35] gRPC Request: /device_service.DeviceProviderService/ListDevices, Time: 2µs, Response: devices:{id:"mouse_1"  type:MOUSE  device_path:"/dev/input/mouse0"  vendor_id:"uuid1"  product_id:"uuid2"}  devices:{id:"keyboard_1"  type:KEYBOARD  device_path:"/dev/input/keyboard1"  vendor_id:"uuid3"  product_id:"uuid4"}  devices:{id:"a061277a-3727-41b1-b211-73634b5f4d05"  type:MOUSE  vendor_id:"c84a1e52-da9b-44cf-a22a-9c165cc616e5"  product_id:"bf2ac3e1-f572-45b5-b51d-6109e6e5dfc6"}
2024/05/22 09:01:34 Device: id:"mouse_1"  type:MOUSE  device_path:"/dev/input/mouse0"  vendor_id:"uuid1"  product_id:"uuid2"
2024/05/22 09:01:34 Device: id:"keyboard_1"  type:KEYBOARD  device_path:"/dev/input/keyboard1"  vendor_id:"uuid3"  product_id:"uuid4"
2024/05/22 09:01:34 Device: id:"a061277a-3727-41b1-b211-73634b5f4d05"  type:MOUSE  vendor_id:"c84a1e52-da9b-44cf-a22a-9c165cc616e5"  product_id:"bf2ac3e1-f572-45b5-b51d-6109e6e5dfc6"
```

2. Delete
```bash
$ go run ./cli/device_consumer/main.go -address localhost:50051 delete
2024/05/22 08:47:36 could not delete device: rpc error: code = Unavailable desc = connection error: desc = "transport: Error while dialing: dial tcp 127.0.0.1:50051: connect: connection refused"
exit status 1
$ go run ./cli/device_consumer/main.go -address localhost:50051 delete -id a061277a-3727-41b1-b211-73634b5f4d05
I0522 09:02:02.229278    6444 service.go:35] gRPC Request: /device_service.DeviceProviderService/DeleteDevice, Time: 3.9µs, Response: success:true
$ go run ./cli/device_consumer/main.go -address localhost:50051 list
I0522 09:02:18.991637    6444 service.go:35] gRPC Request: /device_service.DeviceProviderService/ListDevices, Time: 1.2µs, Response: devices:{id:"mouse_1"  type:MOUSE  device_path:"/dev/input/mouse0"  vendor_id:"uuid1"  product_id:"uuid2"}  devices:{id:"keyboard_1"  type:KEYBOARD  device_path:"/dev/input/keyboard1"  vendor_id:"uuid3"  product_id:"uuid4"}
2024/05/22 09:02:18 Device: id:"mouse_1"  type:MOUSE  device_path:"/dev/input/mouse0"  vendor_id:"uuid1"  product_id:"uuid2"
2024/05/22 09:02:18 Device: id:"keyboard_1"  type:KEYBOARD  device_path:"/dev/input/keyboard1"  vendor_id:"uuid3"  product_id:"uuid4"
```