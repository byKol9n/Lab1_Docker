FROM golang:1.21.2

WORKDIR /noname

COPY go.mod ./
COPY go.sum ./
RUN go mod download

COPY . ./

RUN go build -o ./bin/noname.exe -v ./main.go

EXPOSE 4040

CMD [ "./bin/noname.exe" ]