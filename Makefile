server-gen:
	swagger generate server -f spec/swagger.yaml \
		--with-flatten=expand \
		--server-package internal/restapi \
		--model-package internal/restapi/models
