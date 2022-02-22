
## install solc windows / linux / WSL - ubuntu

	windows 
		
		downlowd solc package https://github.com/ethereum/solidity/releases


	linux / wsl 


		`solc --version`  -Display the currently solc versions:
		
		`solc-select versions` -Display the currently installed versions:
		
		
		`solc-select install` - check Available version
			Available versions to install:
				0.3.6
				0.4.0
				...
				0.8.0
				0.8.1
		
		`solc-select install 0.8.10` - install version 0.8.10
				Installing '0.8.10'...
				Version '0.8.10' installed.
		
		
		`solc-select use 0.8.10` - change solc version to use 0.8.10
	


## install Go on WSL / ubuntu 

	download go v1.17.3   from https://go.dev/doc/install
	
	in wsl: 
	
	sudo apt-get remove golang-go
	
	** move or copy go package to $home dir
	cd $home
	cp /mnt/c/User/xdotk/downloads/xxxx.tar.gz ./
	sudo apt autoremove
	sudo tar -xvf go1.17.3.linux-amd64.tar.gz
	sudo mv go /usr/local
	export GOROOT=/usr/local/go
	export GOPATH=$HOME/go
	export PATH=$GOPATH/bin:$GOROOT/bin:$PATH
	source ~/.profile
	go version 

## install abigen 1.10.14 on wsl / ubuntu 

	https://geth.ethereum.org/docs/install-and-build/installing-geth#install-on-ubuntu-via-ppas
	
	sudo add-apt-repository -y ppa:ethereum/ethereum
	sudo apt-get update
	sudo apt-get install ethereum
	
	/usr/bin/abigen --version
	
	usage:
	
	 /usr/bin/abigen --abi=../build/test.abi --bin=../build/test.bin --pkg=api --out=../deploy/test/test.go

