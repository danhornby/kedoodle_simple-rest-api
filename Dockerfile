FROM golang:1.20 AS build

WORKDIR /go/src/github.com/kedoodle/simple-rest-api
COPY go.* ./
RUN go mod download

COPY . ./
RUN CGO_ENABLED=0 GOOS=linux go build -o simple-rest-api .

FROM gcr.io/distroless/static-debian11:nonroot AS release

COPY --from=build /go/src/github.com/kedoodle/simple-rest-api/simple-rest-api /usr/bin/

ENTRYPOINT ["simple-rest-api"]
