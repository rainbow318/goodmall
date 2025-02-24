FROM golang:1.23

WORKDIR /usr/src/goodmall

# 设置代理（for China）
ENV GOPROXY=https://goproxy.cn

# pre-copy/cache go.mod for pre-downloading dependencies and only redownloading them in subsequent builds if they change
COPY app/frontend/go.mod app/frontend/go.sum ./app/frontend/
COPY rpc_gen rpc_gen
COPY common common

RUN cd app/frontend && go mod download

COPY app/frontend app/frontend

RUN cd app/frontend && go build -v -o /opt/goodmall/frontend/server 

COPY app/frontend/conf /opt/goodmall/frontend/conf
COPY app/frontend/static /opt/goodmall/frontend/static
COPY app/frontend/template /opt/goodmall/frontend/template

WORKDIR /opt/goodmall/frontend

EXPOSE 8080

CMD ["./server"]
