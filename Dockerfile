FROM golang:1.15.6


WORKDIR /app
COPY . .
EXPOSE 5000
ENTRYPOINT ["./go-gim-api" ] 

