# djan-go-deployer

djan-go-deployer is a Go package that allows you to deploy any Django project using a simple YAML file.

## Usage

To deploy your Django project, use the following command:

```bash
deployer -conf /path/to/your/yaml/file.yaml
```

## Installation

To install djan-go-deployer, follow these steps:

```bash
git clone https://github.com/ahmedtouahria/djan-go-deployer.git
cd djan-go-deployer/cmd
go build -o deployer
sudo mv deployer /usr/local/bin/
deployer --version
```

## YAML File Template

Here's a template for the YAML file that djan-go-deployer expects:

```yaml
# Nginx Configuration
DB_NAME: Fatal
DB_USER: postgre
DB_PASSWORD: "1954"
DB_HOST: "127.0.0.1"
DB_PORT: test
DJANGO_SECRET: secret
ALLOWED_HOSTS: 127.0.0.1,localhost

SSL: "false"
SECURE: "false" # Be careful! This option will implement all security practices and settings in your web server configurations
NGINX_PORT: "80" # Nginx port
DJANGO_PORT: "8000" # Django process port
CHANNELS_PORT: "8001" # Django process port

DJANGO_DIRECTORY: /home/user/project # Django directory, must contain a requirements.txt file
INCLUDE_CHANNELS: "true" # Indicates whether you're using channels in your Django project
CHANNELS_HOSTS: localhost
INCLUDE_CELERY: "false"
PROJECT_NAME: django
```

Feel free to adjust the values according to your project's configuration.

---

This README provides a quick overview of djan-go-deployer. For more details and advanced usage, please refer to the package documentation and examples.
