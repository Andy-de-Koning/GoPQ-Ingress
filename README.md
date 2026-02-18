# 🛡️ GoPQ-Ingress
**Post-Quantum Cryptography (PQC) TLS Gateway for Hobbyists & Developers**

>"Security of the future, running on your server today."

GoPQ-Ingress is a lightweight, powerful reverse proxy written in Go. The goal? To demonstrate that **Post-Quantum Cryptography** isn't just for tech giants, but also for the hobbyist looking to secure their home lab against "Harvest Now, Decrypt Later" attacks.

[Proof of PQC Connection]

<img width="350" height="474" alt="image" src="https://github.com/user-attachments/assets/d4f14f1b-5e21-45c9-8a5d-5bf2c48115cf" />

The screenshot above shows an active X25519MLKEM768 handshake via this ingress.

## 🚀 Why this project?
In the future, quantum computers will be capable of cracking current encryption methods (such as RSA and classical Elliptic Curves). Major players like Google and Cloudflare are already testing new standards.

With **GoPQ-Ingress**, you can too. This server enforces the **Hybrid Post-Quantum Handshake (X25519 + Kyber/ML-KEM)**. This means your data traffic is already protected against the computers of 10 years from now.

## ✨ Features
* 🔒 Quantum-Safe: Uses standard CurveID(0x11ec) (X25519MLKEM768).

* 📜 Auto-SSL: Automatic certificates via Let's Encrypt (powered by CertMagic).

* ⚡ Lightweight: No heavy database required, just a simple config.yml.

* 🔌 WebSockets: Out-of-the-box support for real-time apps.

* 🕵️ Privacy Header: Adds X-PQC-Enabled: true to requests sent to your backend.

## 🛠️ Installation

### 🐳 Docker (Recommended)
No installation required if you have Docker.

#### 1. Build the image
```Bash
git clone https://github.com/andy-de-koning/GoPQ-Ingress.git
cd GoPQ-Ingress
docker build -t gopq-ingress .
```
#### 2. Create the config file
Create a file named config.yml:

```YAML
# config.yml
email: "your-email@example.com" # For Let's Encrypt notifications

routes:
  "mydomain.com": "http://127.0.0.1:8080"
  "app.mydomain.com": "http://192.168.1.50:3000"
  "socket.mydomain.com": "http://127.0.0.1:9000" # Works with WS too!
```
#### 3. Start the container
This command starts the server and ensures your certificates are persisted.

```Bash
docker run -d \
  --name gopq-ingress \
  --restart always \
  -p 80:80 -p 443:443 \
  -v $(pwd)/config.yml:/app/config.yml \
  -v pqc_certs:/root/.local/share/certmagic \
  gopq-ingress
```

### 🛠️ Manual Installation (Without Docker)


#### 1. Prerequisites
* Go 1.23 or higher (for optimal PQC support).
* A Linux server (e.g., Ubuntu) or local environment.
* Ports 80 and 443 must be available.

#### 2. Download & Build
```Bash
git clone https://github.com/andy-de-koning/GoPQ-Ingress.git
cd GoPQ-Ingress
go mod tidy
go build -o gopq-ingress main.go
```
#### 3. Configuration
Create a config.yml file next to the executable:

```YAML
# config.yml
email: "your-email@example.com" # For Let's Encrypt notifications

routes:
  "mydomain.com": "http://127.0.0.1:8080"
  "app.mydomain.com": "http://192.168.1.50:3000"
  "socket.mydomain.com": "http://127.0.0.1:9000" # Works with WS too!
```
#### 4. Run
Since the server runs on ports 80 and 443, you will need root privileges (or use setcap):

```Bash
sudo ./gopq-ingress
```

## 🧠 Technical Deep Dive
This project utilizes Go's crypto/tls library and overrides the default CurvePreferences. We prioritize the Hybrid Kyber method.

```Go
tlsConfig.CurvePreferences = []tls.CurveID{
    tls.CurveID(0x11ec), // X25519MLKEM768
    tls.X25519,
    tls.CurveP256,
}
```
This ensures that when a modern browser (such as Chrome or Edge) connects, a quantum-safe key exchange takes place. Older clients will gracefully fall back to standard X25519.

## 🤝 Contributing
Have ideas to make this even better? Docker support? Metrics?
Fork the repo and submit a Pull Request! Let's make the internet safer together.

Made with ❤️ and ☕ by Andy.
