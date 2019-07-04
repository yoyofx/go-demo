FROM golang


RUN mkdir -p /root/gopath
RUN mkdir -p /root/gopath/src
RUN mkdir -p /root/gopath/pkg
RUN mkdir -p /root/gopath/bin
RUN mkdir -p /root/code
ENV GOPATH /root/gopath


 WORKDIR /godocker

 ADD . /godocker

 RUN go build application.go

 EXPOSE 8090

 ENTRYPOINT ["./application"]