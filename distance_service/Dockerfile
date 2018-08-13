FROM golang:1.10.3-alpine as build
RUN apk --update add git
RUN apk --update add ca-certificates
WORKDIR $GOPATH/src/github.com/andrewdruzhinin/travel-cost-app/distance_service
COPY . .

RUN go version && go get -u -v golang.org/x/vgo
RUN CGO_ENABLED=0 vgo install ./...

FROM scratch
COPY --from=build /etc/ssl/certs/ca-certificates.crt /etc/ssl/certs/ca-certificates.crt
COPY --from=build /go/bin/distance_service /
CMD ["/distance_service"]