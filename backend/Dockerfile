FROM golang:1.23

WORKDIR /app

COPY . .

RUN go mod download && go mod verify && go mod tidy && go mod vendor

RUN go build -o /bin/app 

EXPOSE 8380

CMD ["/bin/app"]
