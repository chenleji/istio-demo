FROM golang:1.10 as builder
MAINTAINER chenleji@gmail.com

COPY ./ /go/src/github.com/chenleji/istio-demo/
WORKDIR /go/src/github.com/chenleji/istio-demo/
RUN go build -o istio-demo

FROM centos:7
MAINTAINER chenleji@gmail.com
RUN mkdir -p /conf && \
    mkdir -p /views && \
    cp /usr/share/zoneinfo/Asia/Shanghai /etc/localtime
WORKDIR /
COPY --from=builder /go/src/github.com/chenleji/istio-demo/istio-demo .
COPY --from=builder /go/src/github.com/chenleji/istio-demo/conf/app.conf ./conf/
COPY --from=builder /go/src/github.com/chenleji/istio-demo/views/get.tpl ./views/

CMD /istio-demo