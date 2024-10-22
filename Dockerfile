FROM golang:latest as builder

WORKDIR /app

COPY . .

RUN GOOS=linux CGO_ENABLED=0 go build -ldflags="-w -s" -o server ./cmd


#FROM scratch

#COPY --from=builder /app/server .
#COPY --from=builder /app/cmd/config.env .
#COPY --from=builder /app/cmd/run.sh .

#RUN chmod +x ./app/run.sh

#RUN ./app/run.sh

CMD [ "./server" ]