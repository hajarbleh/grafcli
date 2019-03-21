# grafcli
Simple grafana CLI to manage your dashboards easier

## Usage
1. Download the binary
   ```
   https://github.com/hajarbleh/grafcli/releases/latest
   ```
2. Test run your binary
   `grafcli`
3. Fill config file
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
