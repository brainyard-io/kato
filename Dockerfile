FROM golang:alpine AS build-env
ADD . /src
RUN cd /src && go build -ldflags '-w -extldflags "-static"' -o kato

FROM alpine
WORKDIR /opt
COPY --from=build-env /src/kato /opt/
ENTRYPOINT ./kato 