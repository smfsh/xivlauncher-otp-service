# Start with the official Go image to create a build artifact.
FROM golang:latest as builder

# Set the Current Working Directory inside the container
WORKDIR /app

# Copy everything from the current directory to the PWD inside the container
COPY . .

# Compile the application.
# Disable CGO and recompile the Go source files to create a static binary.
RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o main .

# Use a scratch (empty) container to run the application
FROM scratch

# Copy the pre-built binary file from the previous stage.
COPY --from=builder /app/main .

# Command to run the executable
CMD ["./main"]