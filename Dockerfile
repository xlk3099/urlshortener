
# Start from a image with the latest version of Go installed
# and a workspace (GOPATH) configured at /go.
FROM golang

# get urlshortener from github
RUN go get github.com/xlk3099/urlshortener

# Build the urlshortener command inside the container.
RUN go install github.com/xlk3099/urlshortener

# Run the urlshortener command by default when the container starts.
ENTRYPOINT /go/bin/urlshortener

# Expose the application on port 8080
EXPOSE 8080
