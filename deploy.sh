#!/bin/bash
# env
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
    $GOPATH/bin/migrate \
    $USER@$APIDOMAIN:~/scp/app/

# deployment and authorisation of binary files
echo "Setting up files on remote server..."
ssh $USER@$APIDOMAIN -i ~/Desktop/funcy_gcp '
    sudo mkdir -p /go/src/app
    sudo rm -f /go/src/app/main /go/src/app/file-server /go/src/app/migrate

    sudo cp -r ~/scp/app/* /go/src/app/

    for file in main file-server migrate; do
        if [ -f "/go/src/app/$file" ]; then
            sudo chmod +x "/go/src/app/$file"
        fi
    done

    # clean up
    rm -rf ~/scp/*
'

echo "Deployment completed successfully!"
