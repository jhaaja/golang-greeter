FROM golang:1.15 as build

WORKDIR /go/src/app
ADD . /go/src/app
RUN go build -o /go/bin/app

FROM gcr.io/distroless/static
COPY --from=build /go/bin/app /
CMD ["/app"]
