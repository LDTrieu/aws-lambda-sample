export AWS_ACCESS_KEY_ID = 123456AWSLAMBDA
export AWS_SECRET_ACCESS_KEY = 123456AWSLAMBDA
export aws_key = 123456AWSLAMBDA
export aws_secret_key = 123456AWSLAMBDA


dev_sample_deploy:
	sam build --template template.yaml
	sam deploy --config-file template.toml