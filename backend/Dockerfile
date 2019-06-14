FROM golang:latest

WORKDIR /go/src/github.com/henriqueholanda/backend-challenge/backend

# copy the content
COPY . .

# get dependency manager
RUN go get -u github.com/golang/dep/cmd/dep

# install dependencies
RUN dep ensure -v

# build binary
RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o checkout-api .

EXPOSE 80

# execute
CMD ["./checkout-api"]