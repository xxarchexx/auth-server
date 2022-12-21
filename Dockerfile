###delevopment
FROM golang:1.19 as development
RUN go install github.com/go-delve/delve/cmd/dlv@v1.9.0
WORKDIR /build
COPY . .
COPY ./scripts/run.sh .
# RUN ./run.sh

###debug
FROM golang:1.19 as debug
RUN go install github.com/go-delve/delve/cmd/dlv@v1.9.0
WORKDIR /build
COPY . .
COPY ./scripts/dlv.sh .
# RUN ./dlv.sh