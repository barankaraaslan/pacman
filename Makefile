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
	find unpack
clean-test-docker-scratch:
	docker image rm pacman
test-docker-scratch:
	# Test if pacman is still runnable in a container from scratch
	docker build -t pacman .
	docker run --rm pacman