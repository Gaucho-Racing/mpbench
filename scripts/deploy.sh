#!/bin/bash

# Get the version from config.go
VERSION=$(grep 'var Version =' config/config.go | cut -d '"' -f 2)

# Check if version was successfully extracted
if [ -z "$VERSION" ]
  then
    echo "Error: Could not extract version from config/config.go"
    exit 1
fi

# Check if docker is installed
if ! [ -x "$(command -v docker)" ]; then
  echo 'Error: docker is not installed.' >&2
  exit 1
fi

echo "Building container for MPBench v$VERSION"
# Build the docker container
docker build -t gauchoracing/mpbench:"$VERSION" -t gauchoracing/mpbench:latest --platform linux/amd64,linux/arm64 --push --progress=plain .

echo "Container deployed successfully"

# Check if GitHub CLI is installed
if ! command -v gh &> /dev/null
then
    echo "GitHub CLI (gh) is not installed. Please install it to proceed."
    exit 1
fi

# Create a release tag
git tag -s v$VERSION -m "Release version $VERSION"
git push origin v$VERSION

# Create a release
gh release create v$VERSION --generate-notes

echo "Package released successfully for version $VERSION"

echo "Restarting service on gr-verstappen ec2 instance"
sudo ssh -i ~/.ssh/bharat-verstappen.pem ec2-user@verstappen-ec2.gauchoracing.com "cd verstappen-infra && git pull && docker compose -f mpbench.yml down && docker compose -f mpbench.yml pull && docker compose -f mpbench.yml up -d"