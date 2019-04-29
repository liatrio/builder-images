SUBDIRS := $(wildcard builder-image-*/.)
SKAFFOLD_FLAGS := 

ifeq (, $(shell which container-structure-test))
$(eval SKAFFOLD_FLAGS := --skip-tests)
endif

all: $(SUBDIRS)
$(SUBDIRS):
	cd $@ && skaffold build $(SKAFFOLD_FLAGS)

.PHONY: all $(SUBDIRS)