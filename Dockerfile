ARG BASE_IMAGE=golang:1.19-alpine
ARG CONTAINER=alpine:3.16

FROM ${BASE_IMAGE}

ARG APP_NAME=detfes

RUN apk update
RUN apk add --update upx g++ musl-dev git lapack-dev blas-dev libjpeg-turbo-dev \
    && apk add dlib --repository=http://dl-cdn.alpinelinux.org/alpine/edge/testing \
    && wget -O- -nv https://raw.githubusercontent.com/golangci/golangci-lint/master/install.sh | sh -s

WORKDIR /src

COPY go.* ./
RUN go mod download -x

COPY . .
RUN go vet . && golangci-lint run
RUN CGO_ENABLED=1 go build -ldflags="-s -w" -o ${APP_NAME} && upx -q --best --lzma ${APP_NAME}

FROM ${CONTAINER}
ARG APP_NAME=detfes

RUN apk add --update libstdc++ libgcc libjpeg-turbo-dev lapack-dev blas-dev tzdata ca-certificates \
    && apk add dlib --repository=http://dl-cdn.alpinelinux.org/alpine/edge/testing

WORKDIR /app

COPY --from=0 /src/${APP_NAME} .
COPY config.yaml .
COPY models models/
RUN adduser -D groot -G users && chown -R groot.users /app

EXPOSE 2804

USER groot

ENTRYPOINT ["./detfes"]
