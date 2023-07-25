FROM golang:alpine

RUN apk update && apk add --no-cache git \
curl

WORKDIR /app

COPY . .

# RUN curl -fLo install.sh https://raw.githubusercontent.com/cosmtrek/air/master/install.sh \
#     && chmod +x install.sh && sh install.sh && cp ./bin/air /bin/air

RUN go install github.com/cosmtrek/air@latest

RUN go mod tidy

RUN go build -o binary

# ENTRYPOINT ["/app/binary"]

CMD ["air", "-c", "air.toml"]