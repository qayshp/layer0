L0_VERSION?=$(shell git describe --tags)

release:
	$(MAKE) -C api release
	$(MAKE) -C cli release
	$(MAKE) -C runner release
	$(MAKE) -C setup release
	$(MAKE) -C plugins/terraform release

	rm -rf build
	for os in linux darwin windows; do \
		cp -R cli/build . ; \
		cp -R setup/build . ; \
		cp -R plugins/terraform/build . ; \
		cd build/$$os && zip -r layer0_$(L0_VERSION)_$$os.zip * && cd ../.. ; \
		aws s3 cp build/$$os/layer0_$(L0_VERSION)_$$os.zip s3://xfra-layer0/release/$(L0_VERSION)/layer0_$(L0_VERSION)_$$os.zip ; \
	done

unittest:
	$(MAKE) -C api test
	$(MAKE) -C cli test
	$(MAKE) -C runner test
	$(MAKE) -C setup test

smoketest:
	$(MAKE) -C tests/smoke test

full-smoketest: unittest
	$(MAKE) -C cli release
	$(MAKE) -C setup release
	$(MAKE) -C scripts -f Makefile.flow push
	$(MAKE) -C tests/smoke fulltest
	$(MAKE) -C scripts -f Makefile.flow delete

.PHONY: release unittest smoketest full-smoketest
