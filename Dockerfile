FROM jakubknejzlik/godog as builder
WORKDIR /go/src/github.com/jakubknejzlik/godog-graphql
COPY . .
RUN apk add --update build-base && go get -t ./...
RUN godog -o /tmp/godog

FROM jakubknejzlik/wait-for as wait-for

FROM alpine
VOLUME [ "/godog/features" ]
WORKDIR /godog
COPY --from=wait-for /usr/local/bin/wait-for /usr/local/bin/wait-for
COPY --from=builder /tmp/godog /usr/local/bin
CMD [ "/bin/sh","-c","wait-for ${GRAPHQL_URL} && godog" ]
