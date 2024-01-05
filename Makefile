build:
	GOOS=linux GOARCH=amd64 go build -v -ldflags '-d -s -w' -a -tags netgo -installsuffix netgo -o build/bin/app ./lambda

# Define variables
TF_FOLDER := terraform

# Targets
init:
	cd $(TF_FOLDER) && terraform init

plan:
	cd $(TF_FOLDER) && terraform plan

apply:
	cd $(TF_FOLDER) && terraform apply --auto-approve

destroy:
	cd $(TF_FOLDER) && terraform destroy --auto-approve

.PHONY: init plan apply destroy