FROM golang:1.24.5-alpine AS prod-deps

WORKDIR /usr/src/app

COPY app/go.mod app/go.sum ./
RUN go mod download && go mod verify


FROM prod-deps AS builder

WORKDIR /usr/src/app

COPY --from=prod-deps /usr/src/app /usr/src/app
COPY app/. .

RUN go build -v -o /usr/local/bin/app ./

FROM prod-deps AS runner
WORKDIR /usr/src/app

COPY --from=builder /usr/local/bin/app /usr/local/bin/app

CMD ["/usr/local/bin/app"]