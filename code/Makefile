gen.keypair:
	mkdir -p ${JWT_KEYPAIR_DIRECTORY}
	openssl genpkey -algorithm ed25519 -out ${JWT_KEYPAIR_DIRECTORY}/${JWT_PRIVKEY_FILENAME}
	openssl pkey -in ${JWT_KEYPAIR_DIRECTORY}/${JWT_PRIVKEY_FILENAME} -pubout -out ${JWT_KEYPAIR_DIRECTORY}/${JWT_PUBKEY_FILENAME}

test:
	go test -v ./...
