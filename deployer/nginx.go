package deployer
import (
	"bytes"
	"html/template"
)
// CreateNgixConf creates a ngix configuration and return path
// NGINXConfig represents the NGINX configuration.
type NGINXConfig struct {
	ServerName  string
	ListenPort  int
	DjangoPort  int
	RootDir     string
}

// CreateNGINXConf generates NGINX configuration based on the provided NGINXConfig.
func CreateNGINXConf(config NGINXConfig) (string, error) {
	tmpl := `server {
    listen {{.ListenPort}};
    server_name {{.ServerName}};

    location /static/ {
        root {{.RootDir}};
    }

    location /media/ {
        root {{.RootDir}};
        }

    location / {
        include proxy_params;
        proxy_pass http://localhost:{{.DjangoPort}};
    }

}`

	t, err := template.New("nginx").Parse(tmpl)
	if err != nil {
		return "", err
	}

	var buf bytes.Buffer
	err = t.Execute(&buf, config)
	if err != nil {
		return "", err
	}

	return buf.String(), nil
}

