FROM golang:1.20 as modules

COPY go.mod go.sum /m/
RUN cd /m && go mod download

FROM golang:1.20
COPY --from=modules /go/pkg /go/pkg

WORKDIR /app
COPY . .

CMD ["go", "run", "cmd/discipline-api/main.go"]
