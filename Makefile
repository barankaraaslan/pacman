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
test-server: build
	{ ./pacman server & echo $$! > server.PID; }
	sleep 2
	curl localhost:5001
	kill `cat server.PID`
	rm server.PID
	@echo "Test passed"
test-server-in-docker: 
	docker build -t pacman .
	{ docker run -p 5001:5001 --rm pacman server & echo $$! > server.PID; }
	sleep 2
	curl localhost:5001
	kill `cat server.PID`
	rm server.PID
	@echo "Test passed"
