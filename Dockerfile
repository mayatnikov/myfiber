FROM golang:alpine as app-builder
WORKDIR /go/src/app
COPY . .
# RUN apk add git
# Static build required so that we can safely copy the binary over.
# `-tags timetzdata` embeds zone info from the "time/tzdata" package.
# RUN go mod vendor
RUN CGO_ENABLED=0 go build -o /go/bin/app -ldflags '-extldflags "-static"'

FROM scratch
# the test program:
COPY --from=app-builder /go/bin/app /bin/app
COPY ./views/* /var/views/
# the tls certificates:
# NB: this pulls directly from the upstream image, which already has ca-certificates:
# COPY --from=alpine:latest /etc/ssl/certs/ca-certificates.crt /etc/ssl/certs/
ENTRYPOINT ["/bin/app", "/var/views"]
