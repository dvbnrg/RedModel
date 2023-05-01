# CI/CD Flow

    - This is an overview of the CI/CD flow for this project

## Steps

    - go test ./... -v
    - go build
    - docker build
    - push to registry
    - kubernetes manifest with nginx
        - nginx
        - postgres
        - redis
    - knative
