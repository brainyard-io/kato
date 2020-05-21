FROM golang:alpine AS builder
ADD . /src
RUN cd /src && go build -ldflags '-w -extldflags "-static"' -o kato

FROM scratch
WORKDIR /opt
COPY --from=builder /src/kato /opt/
ENTRYPOINT ./kato 
