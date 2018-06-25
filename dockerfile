FROM golang:latest
RUN mkdir /app 
ADD . /app/ 
RUN apt-get install -y -qq tesseract-ocr
RUN apt-get install -y -qq tesseract-ocr-eng
RUN go build -o main .
ENTRYPOINT ["/app/teixeiract"]