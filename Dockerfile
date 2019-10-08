FROM golang

WORKDIR /app

RUN go get github.com/gorilla/mux
RUN go get go.mongodb.org/mongo-driver/mongo

ENV SRC_DIR=/go/src/github.com/Software-Architecture-2019-2/phets-event-ms/

ADD . $SRC_DIR

RUN cd $SRC_DIR; go build -o phets-event-ms; cp phets-event-ms /app/

EXPOSE 8000

CMD ["./phets-event-ms"]
