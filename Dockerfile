FROM golang:1.20

WORKDIR /genie
COPY . /genie

WORKDIR /genie/cmd/api
RUN go build -o /bin/genie -ldflags '-w -s' -tags netgo -a -installsuffix cgo -v .

EXPOSE 8081
CMD ["/bin/genie"]