FROM golang:1.20

WORKDIR /genie
COPY . /genie

WORKDIR /genie/cmd/api
RUN go build -o /bin/genie -ldflags '-w -s' -tags netgo -a -installsuffix cgo -v .

EXPOSE 8080
CMD ["/bin/genie"]