FROM golang:1.10.5-alpine3.8

# Make the source code path
RUN mkdir -p /go/src/idris/addition-go

# Copy all source code
COPY . /go/src/idris/addition-go

# Run the Go installer
RUN cd /go/src/idris/addition-go && go install

# Indicate the binary as our entrypoint
ENTRYPOINT [ "/go/bin/addition-go" ]

# Expose Go port
EXPOSE 8080