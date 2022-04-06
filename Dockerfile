FROM golang:1.18.0-alpine

WORKDIR /usr/scr/app

COPY . .
RUN go get github.com/bwmarrin/discordgo
RUN go get github.com/joho/godotenv
CMD ["go","run","main.go"]