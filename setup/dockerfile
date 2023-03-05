# build image
FROM golang:alpine as builder

WORKDIR /go/src/go-gin

COPY . .

# RUN go get .

RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o app .

# deploy image
FROM scratch

WORKDIR /bin/

COPY --from=builder /go/src/go-gin/app .
COPY --from=builder /go/src/go-gin/config.json .
COPY --from=builder /go/src/go-gin/cert cert/

CMD [ "./app" ]

EXPOSE 4444