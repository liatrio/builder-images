export SKAFFOLD_DEFAULT_REPO?=artifactory.toolchain.lead.prod.liatr.io/docker-registry/liatrio-dev

SKAFFOLD_FLAGS :=

GIT_BRANCH?=$(shell git rev-parse --abbrev-ref HEAD)
VERSION=$(shell git describe --tags --dirty | cut -c 2-)
IS_SNAPSHOT = $(if $(findstring -, $(VERSION)),true,false)
MAJOR_VERSION := $(word 1, $(subst ., ,$(VERSION)))
MINOR_VERSION := $(word 2, $(subst ., ,$(VERSION)))
PATCH_VERSION := $(word 3, $(subst ., ,$(word 1,$(subst -, , $(VERSION)))))
NEW_VERSION ?= $(MAJOR_VERSION).$(MINOR_VERSION).$(shell echo $$(( $(PATCH_VERSION) + 1)) )

ifeq (, $(shell which container-structure-test))
$(eval SKAFFOLD_FLAGS := --skip-tests)
endif

version:
	@echo "$(VERSION)"

all:
	skaffold build $(SKAFFOLD_FLAGS)

promote:
	@echo "VERSION:$(VERSION) IS_SNAPSHOT:$(IS_SNAPSHOT) LATEST_VERSION:$(LATEST_VERSION)"
ifeq (false,$(IS_SNAPSHOT))
	@echo "Unable to promote a non-snapshot"
	@exit 1
endif
ifneq ($(shell git status -s),)
	@echo "Unable to promote a dirty workspace"
	@exit 1
endif
	git fetch --tags
	git tag -a -m "releasing v$(NEW_VERSION)" v$(NEW_VERSION)
	git push origin v$(NEW_VERSION)


.PHONY: all promote
