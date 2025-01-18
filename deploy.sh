#!/usr/bin/env bash

echo "Building project"

env GOOS=linux GOARCH=arm64 go build -o fit-byte.out

echo "Finished building project"

# Path to your private key
KEY="../project-sprint-infra/projectsprint.key"

# Files/folders to upload
FILES="fit-byte.out .env migrations"

# Destination directory on the remote servers
DEST="/home/ubuntu/"

# Servers to deploy to
SERVERS=(
  "10.0.6.220"
  "10.0.0.11"
  "10.0.7.90"
)

# Loop over servers and upload the files
for SERVER in "${SERVERS[@]}"; do
  echo "Uploading files to ${SERVER}..."
  scp -i "${KEY}" -r ${FILES} "ubuntu@${SERVER}:${DEST}"
  echo "Done uploading to ${SERVER}."
  echo
done
