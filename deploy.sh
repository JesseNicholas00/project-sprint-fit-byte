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
  "10.0.6.144"
  "10.0.3.175"
  "10.0.7.171"
  "10.0.1.22"
  "10.0.6.5"
)

# Loop over servers and upload the files
for SERVER in "${SERVERS[@]}"; do
  echo "Uploading files to ${SERVER}..."
  scp -i "${KEY}" -r ${FILES} "ubuntu@${SERVER}:${DEST}"
  echo "Done uploading to ${SERVER}."
  ssh ubuntu@${SERVER} -i ../project-sprint-infra/projectsprint.key
done
