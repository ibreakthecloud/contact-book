BUILD_TAG:=$(shell git rev-parse --short HEAD)
ifeq ($(BRANCH),)
BRANCH:=$(shell git rev-parse --abbrev-ref HEAD)
endif

image:
	@echo "Building image"
	docker build -t contact-book:$(BRANCH)-$(BUILD_TAG) .