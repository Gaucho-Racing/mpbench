#!/bin/bash

# Check if deploy script exists
if [ ! -f "scripts/deploy.sh" ]; then
    echo "Error: scripts/deploy.sh not found"
    exit 1
fi


# Get the current version
CURRENT_VERSION=$(cat config/config.go | grep "var Version" | awk -F'"' '{print $2}')
if [ -z "$CURRENT_VERSION" ]; then
    echo "Error: Could not extract version from config/config.go"
    exit 1
fi

# Get the latest version
LATEST_VERSION=$(git describe --tags --abbrev=0 | sed 's/^v//')
if [ -z "$LATEST_VERSION" ]; then
    echo "Error: Could not extract latest version from git"
    exit 1
fi

# Increment the version
NEW_VERSION=$(echo $CURRENT_VERSION | awk -F. '{print $1"."$2"."$3+1}')
if [ -z "$NEW_VERSION" ]; then
    echo "Error: Could not increment version"
    exit 1
fi

echo "Current version: $CURRENT_VERSION"
echo "Latest version: $LATEST_VERSION"
echo "New version: $NEW_VERSION"
echo ""

# Compare versions and warn if not latest
if [ "$CURRENT_VERSION" != "$LATEST_VERSION" ]; then
    echo "Warning:  Current version ($CURRENT_VERSION) does not match latest tag ($LATEST_VERSION)"
    echo "          This may indicate the version was manually changed or tags are out of sync"
    exit 1
fi

# Update the version in config/config.go
sed -i '' "s/var Version = \"$CURRENT_VERSION\"/var Version = \"$NEW_VERSION\"/" config/config.go

# Commit the changes
git add config/config.go
git commit -m "version bump"
git push origin main

# Deploy the new version
chmod +x scripts/deploy.sh
./scripts/deploy.sh