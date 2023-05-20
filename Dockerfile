FROM golang:1.20 AS build
ARG VERSION
ARG DESCRIPTION
ARG COMMIT

WORKDIR /go/src/github.com/kedoodle/simple-rest-api
COPY go.* ./
RUN go mod download

COPY . ./
RUN CGO_ENABLED=0 GOOS=linux go build \
    -ldflags "-s -w \
        -X main.Version=${VERSION} \
        -X main.Description=${DESCRIPTION} \
        -X main.Commit=${COMMIT}" \
    -o simple-rest-api .

FROM gcr.io/distroless/static-debian11:nonroot AS release

COPY --from=build /go/src/github.com/kedoodle/simple-rest-api/simple-rest-api /usr/bin/

ENTRYPOINT ["simple-rest-api"]
