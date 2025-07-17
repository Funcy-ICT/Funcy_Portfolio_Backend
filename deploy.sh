#!/bin/bash
# env

# Specify the SSH key path with the -i option (adjust the funcy-gcp part as needed)

if [ ! -f .env.deploy ]; then
    echo "Error: .env.deploy file not found"
    exit 1
fi
source .env.deploy

mkdir -p ./build

echo "Preparing remote directory..."
ssh $USER@$APIDOMAIN -i ~/Desktop/funcy_gcp 'sudo rm -rf ~/scp/* && mkdir -p ~/scp/app'

# main application build
echo "Building main application..."
GOOS=linux GOARCH=amd64 go build -o ./build/main ./cmd/main.go
if [ $? -ne 0 ]; then
    echo "Error: Failed to build main application"
    exit 1
fi

# file server application build
if [ -d "file-server" ]; then
    echo "Building file server..."
    cd file-server
    GOOS=linux GOARCH=amd64 go build -o ../build/file-server
    if [ $? -ne 0 ]; then
        echo "Error: Failed to build file server"
        exit 1
    fi
    cd ..
fi

# transfer files to remote server
echo "Copying files to remote server..."
scp -i ~/Desktop/funcy_gcp \
    ./build/main \
    ./build/file-server \
    $USER@$APIDOMAIN:~/scp/app/

# transfer db directory recursively
echo "Copying db directory to remote server..."
scp -r -i ~/Desktop/funcy_gcp \
    ./db/migration \
    $USER@$APIDOMAIN:~/scp/app/

# deployment and setup of files
echo "Setting up files on remote server..."
ssh $USER@$APIDOMAIN -i ~/Desktop/funcy_gcp '
    sudo mkdir -p /var/www/funcy-backend/app
    sudo rm -rf /var/www/funcy-backend/app/*

    sudo cp -r ~/scp/app/* /var/www/funcy-backend/app

    # Set execute permissions only for binary files
    for file in main file-server; do
        if [ -f "/var/www/funcy-backend/$file" ]; then
            sudo chmod +x "/var/www/funcy-backend/$file"
        fi
    done

    # clean up
    rm -rf ~/scp/*
'

echo "Deployment completed successfully!"
