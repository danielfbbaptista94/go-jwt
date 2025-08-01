FROM golang:1.24

# set working directory
WORKDIR /go/src/app

# Copy the source code
COPY . . 

#EXPOSE the port
EXPOSE 3000

# Build the Go app
RUN go build -o main main.go

# Run the executable
CMD ["./main"]