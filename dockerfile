FROM golang:latest

RUN mkdir /app 
COPY . .
ADD . /app/ 
RUN apt-get update
RUN apt-get install
RUN apt-get install -y -qq tesseract-ocr
RUN apt-get install -y -qq tesseract-ocr-eng
RUN go build -o main .

