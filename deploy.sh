#!/bin/bash

. .env.deploy

ssh $USER@$APIDOMAIN -i ~/Desktop/funcy_gcp 'sudo rm -rf ~/scp/*'

ssh $USER@$APIDOMAIN -i ~/Desktop/funcy_gcp 'mkdir -p ~/scp/app'

go build -o ./build/main ./cmd/main.go

cd file-server && go build -o ../build/file-server && cd ..

scp -i ~/Desktop/funcy_gcp \
    ./build/main \
    ./build/file-server \
    $GOPATH/bin/migrate \
    $USER@$APIDOMAIN:~/scp/app/

ssh $USER@$APIDOMAIN -i ~/Desktop/funcy_gcp '
    # Move files to appropriate locations
    sudo rm -rf /go/src/app/*
    sudo cp -r ~/scp/app/* /go/src/app/
    
    # Set permissions
    sudo chmod +x /go/src/app/main
    sudo chmod +x /go/src/app/file-server
    sudo chmod +x /go/src/app/migrate
    
    # Clean up
    rm -rf ~/scp/*
'
