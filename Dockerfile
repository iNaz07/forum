FROM golang:1.12.10
RUN mkdir /app
ADD . /app/
WORKDIR /app
RUN go build
CMD ["/app/forum"]

EXPOSE 8080