FROM golang:1.22.0 as builder
RUN go env -w GO111MODULE=on
RUN go env -w GOPROXY=https://goproxy.cn,direct

WORKDIR /app/soup
COPY go.mod .
RUN go mod download

COPY . .
RUN CGO_ENABLED=0 go build -o soup


FROM alpine
ENV TIME_ZONE=Asia/Shanghai
COPY --from=builder /usr/share/zoneinfo/$TIME_ZONE /etc/localtime
RUN echo $TIME_ZONE > /etc/timezone

WORKDIR /dist
COPY --from=builder /app/soup/soup .
COPY --from=builder /app/soup/conf ./conf
EXPOSE 8080

ARG PARAMS
CMD /dist/soup $PARAMS