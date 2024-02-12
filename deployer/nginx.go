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
	RootDir     string
	LogFilePath string
}

// CreateNGINXConf generates NGINX configuration based on the provided NGINXConfig.
func CreateNGINXConf(config NGINXConfig) (string, error) {
	tmpl := `server {
    listen {{.ListenPort}};
    server_name {{.ServerName}};

    access_log {{.LogFilePath}}/access.log;
    error_log {{.LogFilePath}}/error.log;

    location / {
        root   {{.RootDir}};
        index  index.html index.htm;
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