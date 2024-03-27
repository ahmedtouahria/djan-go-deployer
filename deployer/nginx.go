package deployer

import (
	"django_deployer/server"
	"fmt"
)

// CreateNGINXConf generates NGINX configuration based on the provided NGINXConfig.
func CreateNGINXConf(data map[string]interface{}) (error, string) {
	tmpl := ``
	nginxConf:=``
	ServerName, ok := server.GetByKey(data, "ALLOWED_HOSTS")
	if !ok {
		return fmt.Errorf("invalid server name key"), ""
	}

	ListenPort, ok := server.GetByKey(data, "NGINX_PORT")
	if !ok {
		return fmt.Errorf("invalid NGINX_PORT key"), ""
	}
	RootDir, ok := server.GetByKey(data, "DJANGO_DIRECTORY")
	if !ok {
		return fmt.Errorf("invalid DJANGO_DIRECTORY key"), ""
	}
	DjangoPort, ok := server.GetByKey(data, "DJANGO_PORT")
	if !ok {
		return fmt.Errorf("invalid DJANGO_PORT key"), ""
	}
	IncludeChannles, ok := server.GetByKey(data, "INCLUDE_CHANNELS")
	if !ok ||  IncludeChannles == "false" {
		fmt.Println("Creating app without DAPHNE WS server ................")
		tmpl = `server {
			listen %s;
			server_name %s;
		
			location /static/ {
				root %s;
			}
		
			location /media/ {
				root %s;
			}
		
			location / {
				include proxy_params;
				proxy_pass http://localhost:%s;
			}
		}`
		nginxConf = fmt.Sprintf(tmpl, ListenPort, ServerName, RootDir, RootDir, DjangoPort)
	} else if IncludeChannles == "true" {
		fmt.Println("Creating app with DAPHNE WS server ................")

		// Define the template
		ChannelsPort, ok := server.GetByKey(data, "CHANNELS_PORT")
		if !ok {
			return fmt.Errorf("invalid CHANNELS_PORT key"), ""
		}
		ChannelsHosts, ok := server.GetByKey(data, "CHANNELS_HOSTS")
		if !ok {
			return fmt.Errorf("invalid CHANNELS_HOSTS key"), ""
		}
		tmpl = `server {
		listen %s;
		server_name %s;
	
		location /static/ {
			root %s;
		}
	
		location /media/ {
			root %s;
		}
	
		location / {
			include proxy_params;
			proxy_pass http://localhost:%s;
		}
	}
	
	server {
		server_name %s;
			location / {
				include proxy_params;
				proxy_pass http://127.0.0.1:%s;
			   	proxy_http_version 1.1;
				proxy_set_header Upgrade $http_upgrade;
				proxy_set_header Connection "upgrade";
			}
}		
	
	`
	nginxConf = fmt.Sprintf(tmpl, ListenPort, ServerName, RootDir, RootDir, DjangoPort,ChannelsHosts,ChannelsPort)

	}

	return nil, nginxConf

}

// BuildNginxFile builds the file
func BuildNginxFile(template string, file_name string) error {
	err := server.CreateFile("/etc/nginx/sites-available/"+file_name, template)
	if err != nil {
		return err
	}
	return nil
}

// RerstartNginxServer it is a helper function to restart the nginx server
func RerstartNginxServer() error {
	err := server.RunCommand("sudo", "service", "nginx", "restart")
	if err == nil {
		fmt.Println("\x1b[32mNginx server restarted successfully\x1b[0m") // "\x1b[32m" for green color, "\x1b[0m" to reset color
	}
	return err
}

func SetSSLCert() error {
	err := server.RunCommand("sudo", "certbot")
	return err
}
