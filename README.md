# simple-rest-api

[![build](https://github.com/kedoodle/simple-rest-api/actions/workflows/build.yaml/badge.svg)](https://github.com/kedoodle/simple-rest-api/actions/workflows/build.yaml)
[![publish](https://github.com/kedoodle/simple-rest-api/actions/workflows/publish.yaml/badge.svg)](https://github.com/kedoodle/simple-rest-api/actions/workflows/publish.yaml)

An implementation of the Innablr technical challenge using the [Gin web framework](https://github.com/gin-gonic/gin).

## Technologies used
- Code repository: GitHub
- Pipelines: GitHub Actions
- API's programming language: Golang
- Image repository: GitHub Packages

## Pipeline
The repository includes a pair of basic GitHub Actions pipelines defined in `.github/workflows`.  

`build.yaml` is run with each push and pull request. It runs three concurrent stages for a quick feedback loop:
- **`lint`**: runs [`golangci-lint`](https://golangci-lint.run/)
- **`test`**: runs the Go tests
- **`build`**: test builds an image but does not push it to a registry

`publish.yaml` is run with each [GitHub release publish](https://github.com/kedoodle/simple-rest-api/releases/new). It builds a container image and pushes it to `ghcr.io/kedoodle/simple-web-api` with three tags:
- `<commit sha>`
- `<release tag>`
- `latest`

## API endpoints

The example `simple-rest-api` implements the following endpoints:
- `GET /`: Returns a basic "Hello World" message
- `GET /status`: Returns some application metadata provided at build-time
    ```json
    {
        "my-application": [
            {
                "description": "my-application's description.",
                "sha": "fd938c35a6e8073eb6eb5f39bf6b07ba1f1308f3",
                "version": "1.0"
            }
        ]
    }
    ```

## Running the application
The `simple-rest-api` is a static application, packaged into a container image, ready to be served. See the [available images here](https://github.com/kedoodle/simple-rest-api/pkgs/container/simple-rest-api).

Example running the application locally from a pre-built image:  
```console
$ docker run --platform linux/amd64 --rm -p 8080:8080 ghcr.io/kedoodle/simple-rest-api:1.0
[GIN-debug] [WARNING] Creating an Engine instance with the Logger and Recovery middleware already attached.

[GIN-debug] [WARNING] Running in "debug" mode. Switch to "release" mode in production.
 - using env:	export GIN_MODE=release
 - using code:	gin.SetMode(gin.ReleaseMode)

[GIN-debug] GET    /                         --> main.helloWorldHandler (3 handlers)
[GIN-debug] GET    /status                   --> main.statusHandler (3 handlers)
[GIN-debug] [WARNING] You trusted all proxies, this is NOT safe. We recommend you to set a value.
Please check https://pkg.go.dev/github.com/gin-gonic/gin#readme-don-t-trust-all-proxies for details.
[GIN-debug] Environment variable PORT is undefined. Using port :8080 by default
[GIN-debug] Listening and serving HTTP on :8080
```
```console
$ curl -s localhost:8080/
Hello World%

$ curl -s localhost:8080/status | jq
{
  "my-application": [
    {
      "description": "my-application's description.",
      "sha": "fd938c35a6e8073eb6eb5f39bf6b07ba1f1308f3",
      "version": "1.0"
    }
  ]
}
```

## Contributing
- Ensure Go 1.20 is installed and on your PATH
- Clone the repository
- Install dependencies with `go mod download`
- Hack away!

## Rambles
- Versioning: I've adhered to the requirements by supplying the version of the application through a metadata file. As an alternative, given the numerous other GitHub features being used, I think it would make sense to use GitHub releases to handle the versioning of the service (as exposed by the `/status` endpoint). We could softly enforce a particular format for versions (e.g. semantic versioning) by only pushing a container image when the release tag matches that format.
- Branching strategy: given that this was an individual technical test, I developed on the `main` branch for simplicity. On a larger project with multiple contributors, trunk-based development may be prudent.
- Portability: idiomatic GitHub Actions workflows aren't particularly portable in that they lean heavily on the various marketplace actions. I kept it simple in this implementation by using lots of common actions, but this does have a few drawbacks. Notably, it can lead to the "works on my machine" problem, as the pipelines are likely to run different commands to what developers may run locally. Something like a `devcontainer` or even just running tests inside a container (both locally and on CI) could save some future headaches.
- Platform: the solution does not explicitly deal with multi-platform. The organisation should consider where the micro-services are intended to be deployed and ensure the appropriate image(s) are built.
- Propagation: this is just one example implementation. If the organisation were to use this implementation as a guideline, it would propagate any potential pitfalls associated with this particular pattern.
- Maintenance: If the organisation is to continue to use this as a scaffold, this repository in itself must be appropriately maintained to remain relevant.
- Divergence: As the organisation gains experience with developing micro services, the various services will inevitably diverge in patterns. General improvements made in each service should be back-ported to this scaffold for the rest of the organisation to benefit.
