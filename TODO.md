# Redis data modeling excercise

## TODO

    - [ ] Create a simple product data model for redis
        - [ ] Keep a record of relevant queries as per each data model
    - [ ] Translate Redis queries and data model to etcd
    - [ ] Create a simple product data model for postgres
    - [ ] Create a simple product data model for mongo
    - [ ] Create a docker file for gRPC server
        - [ ] Set up alpine, busybox and scratch images
    - [ ] Set up docker compose for gRPC server
    - [ ] Set up docker compose for postgres
    - [ ] Set up docker compose for redis
    - [ ] Set up docker compose for mongo
    - [ ] Set up docker compose grpc gateway
    - [ ] Set up k8s cluster
        - [ ] Set up ingress
        - [ ] Set up database
        - [ ] Set up knative
        - [ ] Set up gRPC gateway
    - [ ] Provision a redis database
        - https://free-for.dev/#/
            - render
            - upstash
            - localhost
    - [ ] Provision a postgres database
        - https://free-for.dev/#/
            - render
            - localhost

## Doing

    - [ ] etcd version
    - [ ] High Level Planning
        - [ ] CI/CD planning
        - [ ] Test Planning and Implementation
            - [ ] Unit Testing Redis get test Events needs to be able to accept elements in any order
    - [ ] Document Everything

## Done

    - [x] Create a simple proto file
    - [x] Create a simple gRPC server
    - [x] Set up justfile
