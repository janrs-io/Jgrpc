##-----
# BUILDING
##-----
FROM golang:1.19-alpine AS builder

# Create the working directory for compiling the image
RUN mkdir -p /opt/service/

# Setting up a working directory
WORKDIR /opt/service/

# Copy project
COPY src ./

WORKDIR authservice

# Set environment variables
ENV GO111MODULE=on
ENV GOPROXY=https://goproxy.cn

# Download Dependency Packages
RUN go mod download

# Compiling Projects
RUN CGO_ENABLED=0 go build -ldflags "-s -w" -o auth cmd/main.go

##-----
# RUN
##-----
# Using the busybox image
FROM busybox AS runner

# Create a working directory of project executables
RUN mkdir -p /opt/service/authservice

# Use the same working directory
WORKDIR /opt/service/authservice

# Copy the compiled file
COPY --from=builder /opt/service/authservice .

# Set the port on which to run the auth service
EXPOSE 9001
EXPOSE 50051

CMD ["./auth"]
