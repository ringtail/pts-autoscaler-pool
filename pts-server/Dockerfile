# build stage
FROM golang:1.8-alpine3.6 AS build-env
ADD . /src/github.com/ringtail/pts-autoscaler-pool/pts-server
ENV GOPATH /:/src/github.com/ringtail/pts-autoscaler-pool/pts-server/vendor
WORKDIR /src/github.com/ringtail/pts-autoscaler-pool/pts-server
RUN go build -o app


# test stage
#FROM golang:1.8-alpine3.6
#WORKDIR /src/github.com/ringtail/pts-autoscaler-pool/pts-server
#RUN go test


# release stage
FROM alpine
WORKDIR /app
EXPOSE 8080
COPY --from=build-env /src/github.com/ringtail/pts-autoscaler-pool/pts-server/app /app/
CMD ["./app"]