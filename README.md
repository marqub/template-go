# template-go

***GO***
go run *.go
ps aux
USER              PID  %CPU %MEM      VSZ    RSS   TT  STAT STARTED      TIME COMMAND
212552621       42521   0.0  0.1 556625424   9560 s000  S     5:43PM   0:00.26 go run doc.go main.go

***DOCKER***
https://blog.golang.org/docker
https://hackernoon.com/golang-docker-microservices-for-enterprise-model-5c79addfa811

docker build -t template-go .
docker run --rm -p 9090:8080 --name template-go-microservice -d template-go
docker ps
    CONTAINER ID        IMAGE               COMMAND             CREATED              STATUS              PORTS                    NAMES
    1bb5e747dd64        template-go         "go-wrapper run"    About a minute ago   Up About a minute   0.0.0.0:9090->8080/tcp   template-go-microservice
docker logs template-go-microservice
docker stats template-go-microservice
    CONTAINER                  CPU %               MEM USAGE / LIMIT    MEM %               NET I/O             BLOCK I/O           PIDS
    template-go-microservice   0.00%               1.98MiB / 1.952GiB   0.10%               3.68kB / 1.14kB     0B / 8.19kB         5
docker stop template-go-microservice


use yml to setup env variqble with docker?
