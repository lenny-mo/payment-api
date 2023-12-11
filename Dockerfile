FROM alpine
ADD payment-api /payment-api
ENTRYPOINT [ "/payment-api" ]
EXPOSE 8091