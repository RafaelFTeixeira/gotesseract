FROM golang:latest
COPY . .
RUN apt-get update && apt-get install
RUN apt-get install -y -qq tesseract-ocr
RUN apt-get install -y -qq tesseract-ocr-eng