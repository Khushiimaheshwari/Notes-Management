FROM golang:1.24

WORKDIR /app

# Install MySQL client
RUN apt-get update && apt-get install -y default-mysql-client

COPY go.mod ./
RUN go mod download

COPY . .

RUN go build -o main .

COPY wait-for-mysql.sh /wait-for-mysql.sh
RUN chmod +x /wait-for-mysql.sh

CMD ["sh", "-c", "/wait-for-mysql.sh mysql_db ./main"]
