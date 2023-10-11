FROM golang:1.20 as modules

COPY go.mod go.sum /m/
RUN cd /m && go mod download

FROM golang:1.20
COPY --from=modules /go/pkg /go/pkg

WORKDIR /app
COPY . .

RUN go install github.com/swaggo/swag/cmd/swag@v1.16.2

CMD ["go", "run", "cmd/discipline-api/main.go"]
