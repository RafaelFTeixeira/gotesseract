FROM golang:alpine
LABEL maintainer="Rafael Teixeira <rafaelteixeiradev@gmail.com>"
COPY . .
RUN apk update && apk add tesseract-ocr && apk add tesseract-ocr-data-por