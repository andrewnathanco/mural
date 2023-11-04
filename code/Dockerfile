# build stage
FROM golang:alpine AS builder
RUN apk add --no-cache git gcc libc-dev
WORKDIR /go/src/app
COPY . .
RUN go mod edit -module mural
RUN go get -d -v ./...
RUN go install -v ./...
RUN go build -o /go/bin/mural main.go

# final stage
FROM alpine:latest
ARG VERSION
ENV VERSION=$VERSION
LABEL Name=mural Version=$VERSION
RUN apk --no-cache add ca-certificates
COPY --from=builder /go/bin/mural /mural
RUN touch .env 
EXPOSE 1323
ENTRYPOINT ./mural
