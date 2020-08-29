############################################################################

PKG := github.com/rh/for-sake-of-testing/apps/goservd

REGISTRY := mifomm

VERSION ?= 0.0.7

ARCH := amd64

BASE_IMAGE := alpine

BUILD_IMAGE := goservd-build

DOCKER_RUN_FLAGS := --rm
DOCKER_BUILD_FLAGS := --rm

export VERBOSE ?= 1

VERBOSE_OUTPUT := >&1

DOCKER_PORTS := \
	-p 5000:5051

DOCKER_VOLUMES := \
	-v $(BUILD_IMAGE)-data:/data \
	-v $(BUILD_IMAGE)-node:/data/go/src/$(PKG)/client/node_modules \
	-v $$(pwd):/data/go/src/$(PKG) \
	-v $$(pwd)/build:/build \
	-v $$(pwd)/bin/$(VERSION)/$(ARCH):/data/go/bin \
	-v $$(pwd)/bin/$(VERSION)/$(ARCH):/data/go/bin/linux_$(ARCH)

DOCKER_ENVIRONMENT := \
	-e VERBOSE=$(VERBOSE) \
	-e ARCH=$(ARCH) \
	-e PKG=$(PKG) \
	-e VERSION=$(VERSION)

############################################################################
# Default rule
all: build image

############################################################################

BUILD_IMAGE_BUILDSTAMP := .$(subst .,_,$(BUILD_IMAGE))-image
$(BUILD_IMAGE_BUILDSTAMP):
	@echo "build-image buildstamp: $(BUILD_IMAGE_BUILDSTAMP)"
	@echo "container image: $(BUILD_IMAGE)"
	@echo "  Building container image"
	docker build \
		$(DOCKER_BUILD_FLAGS) \
		-t $(BUILD_IMAGE) \
		--build-arg "ARCH=$(ARCH)" \
		-f Dockerfile.build . \
		$(VERBOSE_OUTPUT)
	@echo "  Creating volume $(BUILD_IMAGE)-data"
	-docker volume rm $(BUILD_IMAGE)-data $(VERBOSE_OUTPUT) 2>&1
	docker volume create $(BUILD_IMAGE)-data $(VERBOSE_OUTPUT)
	@echo "  Creating volume $(BUILD_IMAGE)-node"
	-docker volume rm $(BUILD_IMAGE)-node $(VERBOSE_OUTPUT) 2>&1
	docker volume create $(BUILD_IMAGE)-node $(VERBOSE_OUTPUT)
	@echo "  Running build/init_data.sh in build container to init volumes"
	docker run $(DOCKER_RUN_FLAGS) \
		-v $(BUILD_IMAGE)-data:/data \
		-v $(BUILD_IMAGE)-node:/data/go/src/$(PKG)/client/node_modules \
		-v $$(pwd)/build:/build \
		-e TARGET_UIDGID=$$(id -u):$$(id -g) \
		$(BUILD_IMAGE) \
		/build/init_data.sh \
		$(VERBOSE_OUTPUT)
	echo "$(BUILD_IMAGE)" > $@
	docker images -q $(BUILD_IMAGE) >> $@

############################################################################

enter-build-image: $(BUILD_IMAGE_BUILDSTAMP)
	@echo "Launching into build environment"
	docker run -it \
		$(DOCKER_RUN_FLAGS) \
		$(DOCKER_VOLUMES) \
		--sig-proxy=true \
		$(DOCKER_PORTS) \
		$(DOCKER_ENVIRONMENT) \
		-u $$(id -u):$$(id -g) \
		-w /data/go/src/$(PKG) \
		$(BUILD_IMAGE) \
		/bin/sh

############################################################################

BINARY_PATH:=bin/$(VERSION/$(ARCH)/goservd

.PHONY: build
build: $(BINARY_PATH)

$(BINARY_PATH): build/build.sh $(BUILD_IMAGE_BUILDSTAMP)
	@echo "building binary: $@"
	@mkdir -p $(shell pwd)/bin/$(VERSION)/$(ARCH)
	docker run \
		$(DOCKER_RUN_FLAGS) \
		$(DOCKER_VOLUMES) \
		--sig-proxy=true \
		$(DOCKER_ENVIRONMENT) \
		-u $$(id -u):$$(id -g) \
		-w /data/go/src/$(PKG) \
		$(BUILD_IMAGE) \
		./build/build.sh $(VERBOSE_OUTPUT)

############################################################################

BIN_DOCKERFILE := .goservd-$(ARCH)-$(VERSION)-dockerfile
$(BIN_DOCKERFILE): Dockerfile.goservd
	@echo "generating Dockerfile $@ from $<"
	sed \
		-e 's|ARG_FROM|$(BASE_IMAGE)|g' \
		-e 's|ARG_VERSION|$(VERSION)|g' \
		-e 's|ARG_ARCH|$(ARCH)|g' \
		$< > $@

CONTAINER_NAME := $(REGISTRY)/goservd-$(ARCH)
BUILDSTAMP_NAME := $(subst /,_,$(CONTAINER_NAME)-$(VERSION))

.$(BUILDSTAMP_NAME)-image: $(BIN_DOCKERFILE) $(BINARY_PATH)
	@echo "container image: $(CONTAINER_NAME):$(VERSION)"
	docker build \
		$(DOCKER_BUILD_FLAGS) \
		-t $(CONTAINER_NAME):$(VERSION) \
		-f .goservd-$(ARCH)-$(VERSION)-dockerfile . \
		$(VERBOSE_OUTPUT)
	echo "$(CONTAINER_NAME):$(VERSION)" > $@
	@echo "container image tag: $(CONTAINER_NAME):$(VERSION)"

.PHONY: image
image: .$(BUILDSTAMP_NAME)-image

############################################################################

PUSH_BUILDSTAMP := .$(BUILDSTAMP_NAME)-push

.PHONY: push
push: $(PUSH_BUILDSTAMP)

.%-push: .%-image
	@echo "pushing image: " $$(sed -n '1p' $<)
	docker push $$(sed -n '1p' $<) $(VERBOSE_OUTPUT)

############################################################################

.PHONY: clean
clean: 
	docker volume rm -f $(BUILD_IMAGE)-data $(BUILD_IMAGE)-node $(VERBOSE_OUTPUT)
	rm -f .*-container .*-dockerfile .*-push .*-image
	rm -rf bin
	rm -rf client/node_modules sitedata/built
