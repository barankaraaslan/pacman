build:
	go build .
clean: clean-test-self-package
	-rm pacman
clean-test-self-package:
	rm -rf unpack/ package/ package.tar.gz
test-self-package: clean-test-self-package build
	mkdir unpack
	./pacman self-package
	tar -xf package.tar.gz --directory unpack
	test -f unpack/bin/pacman
	@echo "Test passed"
clean-test-docker-scratch:
	docker image rm pacman
test-docker-scratch:
	# Test if pacman is still runnable in a container from scratch
	docker build -t pacman .
	docker run --rm pacman
	@echo "Test passed"
test-server-up: build
	{ ./pacman server & echo $$! > server.PID; }
	sleep 2
	curl localhost:5001
	kill `cat server.PID`
	rm server.PID
	@echo "Test passed"
test-server-downloaded-file: clean-test-server-downloaded-file build
	# create self-package to host on server
	./pacman self-package
	mkdir packages-to-serve/
	mv package.tar.gz packages-to-serve/
	rm -rf package

	{ ./pacman server & echo $$! > server.PID; }
	sleep 2
	curl localhost:5001/package.tar.gz -O
	cmp package.tar.gz packages-to-serve/package.tar.gz
	kill `cat server.PID`
	rm server.PID package.tar.gz
	@echo "Test passed"
clean-test-server-downloaded-file:
	-rm -rf packages-to-serve/
	-rm package.tar.gz
