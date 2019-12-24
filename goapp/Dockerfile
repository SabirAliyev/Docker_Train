#Virtual Machine to run Go-application
FROM golang:latest
LABEL maintainer="saab01@training.com"
RUN apt-get update
RUN apt-get install -y wget

RUN mkdir test

# Set the Current Working Directory inside the container.
WORKDIR /app

#VOLUME vol1

#Downloading files.
#CMD ["go", "mod", "init", "github.com/callicoder/go-docker"]

# Using Go modules for dependency management.
#CMD ["go mod init"]

# Download all dependencies. Dependencies will be cached if the go.mod and go.sum files are not changed.
#CMD ["go mod download"]

# Copy the source from the current directory to the Working Directory inside the container.
COPY hello_world.go .

# Build the Go app
#CMD ["go", "build"]
#CMD ["go", "build" "-o", "hello_world.go"]
RUN go build -o hello_world .

EXPOSE 8001

# Command to run the executable.
CMD ["./hello_world"]
#CMD ["./main"]


