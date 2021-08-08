FROM python:3.7

COPY ./gcloud/app /app
WORKDIR /app

RUN curl https://sdk.cloud.google.com | bash && /root/google-cloud-sdk/install.sh
