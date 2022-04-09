FROM golang:latest as builder

RUN apt update && apt install libtesseract-dev -y

WORKDIR /app
COPY . ./

RUN curl https://raw.githubusercontent.com/Shreeshrii/tessdata_shreetest/226419f02431675e24c9937643ce42f3675e2b56/digits.traineddata -o digits.traineddata

RUN go mod download && go build -trimpath -ldflags "-s -w -buildid=" -o healthreport

# image for deployment
FROM debian:stable-slim

RUN apt update && apt install libtesseract4 -y && rm -rf /var/lib/apt/lists/*

COPY --from=builder /app/healthreport /usr/local/bin/
COPY --from=builder /app/digits.traineddata /usr/share/tesseract-ocr/4.00/tessdata/

# listen on port 9000
EXPOSE 9000

ENTRYPOINT ["healthreport"]
