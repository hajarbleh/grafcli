# grafcli
Simple grafana CLI to manage your dashboards easier

## Usage
1. Download the binary
   ```
   https://github.com/hajarbleh/grafcli/releases/latest
   ```
2. Make the binary executable
   ```
   chmod +x ./grafcli
   ```
3. Move the binary to your PATH
   ```
   sudo mv ./grafcli /usr/local/bin/grafcli
   ```
4. Test run your binary
   ```grafcli```
5. Fill config file
   Linux:
   ```
   cd $HOME
   mkdir .grafcli
   nano config
   ```
   And fill the config file with the following:
   ```
   url: <your grafana URL>
   api_key: <your grafana API KEY>
   ```
To create/get your API KEY, refer to this [Grafana API Auth](http://docs.grafana.org/http_api/auth/)
