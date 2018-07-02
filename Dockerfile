FROM ubuntu:18.04

COPY ./fruitserver2 /

RUN apt-get update
RUN apt-get install -y curl
EXPOSE 8082
RUN chmod 777 /fruitserver2
RUN chmod +x /fruitserver2
CMD ["/fruitserver2"]
