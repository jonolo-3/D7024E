
FROM larjim/kademlialab:latest

COPY main /home/go/src/main/main



WORKDIR /home/go/src/main

CMD ["./main"]
