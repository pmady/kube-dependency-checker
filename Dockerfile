FROM alpine:3.19

RUN apk add --no-cache ca-certificates

COPY kube-dependency-checker /usr/local/bin/kube-dependency-checker

ENTRYPOINT ["kube-dependency-checker"]
