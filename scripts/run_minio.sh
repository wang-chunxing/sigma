#!/bin/bash

docker run -p 9000:9000 -p 9001:9001 \
  --name ximager-minio \
  -e MINIO_ACCESS_KEY=ximager \
  -e MINIO_SECRET_KEY=ximager-ximager \
  -e MINIO_REGION_NAME=cn-north-1 \
  --rm -d \
  --entrypoint "" \
  quay.io/minio/minio:RELEASE.2023-02-22T18-23-45Z \
  sh -c 'mkdir -p /data/ximager && minio server /data --console-address ":9001"'