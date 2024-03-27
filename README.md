# djan-go-deployer

This Go package allows you to deploy any Django project using a simple YAML configuration file.

## Usage

```bash
deployer -conf /path/to/your/yaml/file.yaml
```

## YAML File Template

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
INCLUDE_CHANNELS: "true" # Indicates if you're using channels in your Django project
CHANNELS_HOSTS: localhost
INCLUDE_CELERY: "false"
PROJECT_NAME: django
```

## Configuration Details

- `DB_NAME`: Name of the database
- `DB_USER`: Database user
- `DB_PASSWORD`: Database password
- `DB_HOST`: Database host
- `DB_PORT`: Database port
- `DJANGO_SECRET`: Django secret key
- `ALLOWED_HOSTS`: Allowed hosts for the Django project

- `SSL`: Enable or disable SSL
- `SECURE`: Enable or disable security practices in web server configurations
- `NGINX_PORT`: Nginx port
- `DJANGO_PORT`: Django process port
- `CHANNELS_PORT`: Django channels process port

- `DJANGO_DIRECTORY`: Path to the Django project directory
- `INCLUDE_CHANNELS`: Indicates if you're using channels in your Django project
- `CHANNELS_HOSTS`: Hosts for Django channels
- `INCLUDE_CELERY`: Indicates if you're using Celery in your Django project
- `PROJECT_NAME`: Name of the Django project

Make sure to adjust the values according to your project requirements.
