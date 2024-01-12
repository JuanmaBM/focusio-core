# Build Stage
FROM golang:1.21 AS build

WORKDIR /app

COPY go.mod go.sum ./

# Cache dependencies separately to improve build speed
RUN go mod download

COPY . .

# Build the Go application
RUN CGO_ENABLED=0 GOOS=linux go build -o myapp 

# Final Stage
FROM alpine:latest

WORKDIR /app

# Copy only the necessary files from the build stage
COPY --from=build /app/myapp .

# Set environment variables if needed
# ENV MY_VARIABLE=value

# Expose the port that your application will run on
EXPOSE 8080

# Command to run the application
CMD ["./myapp"]
