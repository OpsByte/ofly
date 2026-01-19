
# 🚀 OFly - (OpsByte Fly Link)

Turn your laptop, home server, or that dusty Raspberry Pi into a production-ready web server accessible from anywhere on the internet.

## 💰 Why Teams love OFly

**Save money, save time, ship faster.**

- **$0-5/month** vs $20-100/month for typical cloud hosting
- **30 seconds** to deploy vs hours of cloud setup
- **Zero configuration** - one command and you're live
- **Perfect for**: Side projects, client demos, testing webhooks, local-first apps


## 📚 How this tunnel works (for the curious)

ofly creates an encrypted tunnel using WG between your machine and the ofly server. When someone visits your URL:

1. Request hits ofly server over HTTPS
2. Server forwards encrypted traffic through WG tunnel
3. Your local ofly decrypts and sends to your app
4. Response travels back the same way

**Your data is encrypted end-to-end. The ofly server never sees your unencrypted traffic.**

Built on battle-tested technology: WG for tunneling, gVisor netstack for networking, LetsEncrypt for SSL.

## 📦 Installation

**Download** pre-built binaries (no dependencies needed):
- [Windows](https://github.com/opsbyte/ofly/releases/latest/download/ofly.exe)
- [Linux](https://github.com/opsbyte/ofly/releases/latest/download/ofly)  
- [Mac (Apple Silicon)](https://github.com/opsbyte/ofly/releases/latest/download/ofly-darwin-arm64)
- [Mac (Intel)](https://github.com/opsbyte/ofly/releases/latest/download/ofly-darwin)

**Or install from source (Golang Required):**
```bash
sudo apt install golang-go
go install github.com/opsbyte/ofly/ofly@latest
```

## 💻 How to Use (Client Side)

### Run in CLI (Example)

To Forward Port 3000 of your Local System, Run:

```bash
ofly -p 3000
```
or Forward any App running on Port 8080 
```bash
ofly --forward=http://localhost:8080
```

### or Run Entire OFly Client in Docker
```bash
docker run -it --rm --network=host -v ofly_keys:/data \
  ghcr.io/opsbyte/ofly ofly --forward=http://localhost:8080
```

### or Expose Multiple services at once
```bash
# Expose your API and frontend together
ofly --forward=http://10.0.0.2:8000,http://10.0.0.10:9000
```
## 🛠️ Advanced Usage For Nerds

###  Use with basic HTTP auth, To Defend Bots or Scrapers

 ```bash
# Password protect your app in one line
ofly --forward=http://localhost:3000 --limit=$(htpasswd -nbB admin secretpass)
```

### Docker Compose integration
```yaml
services:
  ofly:
    image: ghcr.io/opsbyte/ofly
    command: ofly --forward=http://myapp:3000

  myapp:
    image: your/application
    ports:
      - "3000"
```
### Firewall-proof relay
If any Corporate firewall is blocking you, You can try Enabling HTTPS relay mode:
```bash
OFLY_RELAY=true ofly -p 8080
```

### Use directly in Go code
```go
import "github.com/opsbyte/ofly"

listener, err := ofly.NewListener("myapp")
http.Serve(listener, httpHandler)
```
Then check the URL: `docker compose logs ofly`

## 🏗️ How to Host your own OFly Server

**Requirements**: Just 1GB RAM and a public IP
i.e. You can purchase any 1 GB RAM low-cost VPS from [Hostinger](), []() , []() or any.

**CLI Version:**
```bash
OFLY_RUN_SERVER=true \
OFLY_API=ofly.yourdomain.com \
OFLY_IP=your.server.ip \
OFLY_PORT=443 \
ofly
```

**Docker version:**
```yaml
services:
  ofly-server:
    image: ghcr.io/opsbyte/ofly
    network_mode: host
    environment:
      OFLY_RUN_SERVER: true
      OFLY_PORT: 443
      OFLY_IP: "your.server.ip"
      OFLY_API: ofly.yourdomain.com
```
In this case, Your clients will connect to your server by setting the environment variable `OFLY_API=ofly.yourdomain.com` on their machines.

You can also set the  `OFLY_AUTH`  environment variable to limit which clients can use your server. In that case, clients would need to set the same  `OFLY_AUTH`.

The server listens on port 443 (for HTTPS traffic) and on port 80 (to redirect to HTTPS and for http-01 SSL challenges). It also listens on UDP port  `OFLY_PORT`  for wg UDP traffic. The public instance listens on UDP 443, since it's less likely to be blocked by firewalls.

The server is fully stateless and doesn't require any storage. It caches the wg private key and recent peers on disk to enable instant reconnection of tunnels after server restart. The ofly client will add itself as peer again if the wg handshake with server is missed.

If you're running it behind a reverse proxy like caddy/nginx, you should make sure that the reverse proxy passes through TLS instead of decrypting HTTPS traffic.

## 🌟 Features that make life easier

- **🔗 Persistent URLs**
Your URL stays the same every time you restart. No more updating webhooks or sharing new links.  

- **🌐 Custom domains **  
Use your own domain instead of `.ofly.opsdude.com`. Just add a CNAME record pointing to your ofly subdomain.

-  **📡 PROXY Protocol support**
Your local server receives the real visitor IP addresses, perfect for analytics and logging.

## 🔐 Security considerations

⚠️ **Important**: Subdomains appear in LetsEncrypt certificates logs. Bots will find and scan them within hours.

**Best practices:**
- Always use authentication for sensitive apps (built-in basic auth works great)
- Don't expose vulnerable or unpatched services
- Treat your ofly URL like a public endpoint.
- Be cautious with cookies on shared `.ofly.opsdude.com` domain

## 🔒 Built-in security

✅ **End-to-end SSL encryption** (automatic HTTPS certificates via LetsEncrypt)  
✅ **Zero trust architecture** (your data never passes through our servers unencrypted)  
✅ **Built-in HTTP Basic Auth** (password protect your apps instantly)  
✅ **Private by default** (only you control what's exposed)

## 💡 Pro tips:

- Keep ofly running with `tmux` or `screen` for persistent tunnels
- Use `OFLY_PATH` to customize where keys are stored
- Combine with `ngrok`-style local development for webhooks testing
- Set `OFLY_AUTH` on your self-hosted server to restrict access

Real savings: Developers report saving **$240-1200/year** by hosting demos and small apps from their machines instead of cloud platforms.

## 🆘 Support & Community

- 📖 [Documentation](https://github.com/OpsByte/ofly/tree/master?tab=readme-ov-file#-opsbyte-fly-link)
- 🐛 [Report Issues](https://github.com/opsbyte/ofly/issues)
---

**Made with ❤️ by [OpsByte Technologies](https://www.opsbyte.com)**
