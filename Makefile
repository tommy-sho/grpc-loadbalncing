
SERVICES := gateway backend

.PHONY: image-build image-push

image-build:
	for service in ${SERVICES}; do \
		docker image build -t gcr.io/my-first-project-236315/grpc-test/$$service:latest $$service/; \
	done

image-push:
	for service in ${SERVICES}; do \
		docker push gcr.io/my-first-project-236315/grpc-test/$$service:latest; \
	done