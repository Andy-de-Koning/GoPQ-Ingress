# 🛡️ GoPQ-Ingress
**Post-Quantum Cryptography (PQC) TLS Gateway voor Hobbyisten & Developers**

> "Beveiliging van de toekomst, vandaag al draaiend op jouw server."

GoPQ-Ingress is een lichtgewicht, krachtige reverse proxy geschreven in Go. Het doel? Aantonen dat **Post-Quantum Cryptografie** niet alleen voor tech-reuzen is, maar ook voor de hobbyist die zijn home-lab wil beveiligen tegen "Harvest Now, Decrypt Later" aanvallen.

[Bewijs van PQC Verbinding]

<img width="350" height="474" alt="image" src="https://github.com/user-attachments/assets/d4f14f1b-5e21-45c9-8a5d-5bf2c48115cf" />

*Bovenstaande screenshot toont een actieve X25519MLKEM768 handshake via deze ingress.*

## 🚀 Waarom dit project?
Kwantumcomputers zullen in de toekomst in staat zijn om de huidige encryptie (zoals RSA en klassieke Elliptic Curves) te kraken. Grote partijen zoals Google en Cloudflare zijn al aan het testen met nieuwe standaarden.

Met **GoPQ-Ingress** doe jij dat ook. Deze server dwingt de **Hybride Post-Quantum Handshake (X25519 + Kyber/ML-KEM)** af. Hierdoor is je dataverkeer nu al beschermd tegen de computers van over 10 jaar.

## ✨ Features
* 🔒 **Quantum-Safe:** Gebruikt standaard `CurveID(0x11ec)` (X25519MLKEM768).
* 📜 **Auto-SSL:** Automatische certificaten via Let's Encrypt (dankzij CertMagic).
* ⚡ **Lichtgewicht:** Geen zware database nodig, enkel een simpele `config.yml`.
* 🔌 **WebSockets:** Out-of-the-box ondersteuning voor real-time apps.
* 🕵️ **Privacy Header:** Voegt `X-PQC-Enabled: true` toe aan requests naar je backend.


## 🛠️ Installatie


### 🐳 Docker (Aanbevolen)
Je hoeft niets te installeren als je Docker hebt.

#### 1. Bouw de image
```bash
git clone https://github.com/andy-de-koning/GoPQ-Ingress.git
cd GoPQ-Ingress
docker build -t gopq-ingress .
```

#### 2. maak de config file
Maak een bestand genaamd config.yml:
```yml
# config.yml
email: "jouw-email@voorbeeld.nl" # Voor Let's Encrypt meldingen

routes:
  "mijndomein.nl": "http://127.0.0.1:8080"
  "app.mijndomein.nl": "http://192.168.1.50:3000"
  "socket.mijndomein.nl": "http://127.0.0.1:9000" # Werkt ook met WS!
```

#### 3. Start de container
Dit commando start de server en zorgt dat je certificaten bewaard blijven.

```bash
docker run -d \
  --name gopq-ingress \
  --restart always \
  -p 80:80 -p 443:443 \
  -v $(pwd)/config.yml:/app/config.yml \
  -v pqc_certs:/root/.local/share/certmagic \
  gopq-ingress
```

### 🛠️ Handmatige Installatie (Zonder Docker)

#### 1. Vereisten
* Go 1.23 of hoger (voor de beste PQC support).
* Een Linux server (bijv. Ubuntu) of gewoon lokaal.
* Poort 80 en 443 moeten vrij zijn.

#### 2. Download & Bouw
```bash
git clone https://github.com/andy-de-koning/GoPQ-Ingress.git
cd GoPQ-Ingress
go mod tidy
go build -o gopq-ingress main.go
```
#### 3. Configuratie
Maak een bestand genaamd config.yml naast de executable:
```yml
# config.yml
email: "jouw-email@voorbeeld.nl" # Voor Let's Encrypt meldingen

routes:
  "mijndomein.nl": "http://127.0.0.1:8080"
  "app.mijndomein.nl": "http://192.168.1.50:3000"
  "socket.mijndomein.nl": "http://127.0.0.1:9000" # Werkt ook met WS!
```
#### 4. Starten
Omdat de server op poort 80 en 443 draait, heb je root-rechten nodig (of setcap):
```bash
sudo ./gopq-ingress
```
## 🧠 Hoe werkt het technisch?

Dit project maakt gebruik van de crypto/tls bibliotheek van Go en overschrijft de standaard CurvePreferences. We geven prioriteit aan de Hybride Kyber methode.

Go
```go
tlsConfig.CurvePreferences = []tls.CurveID{
    tls.CurveID(0x11ec), // X25519MLKEM768
    tls.X25519,
    tls.CurveP256,
}
```
Dit zorgt ervoor dat als een moderne browser (zoals Chrome of Edge) verbinding maakt, er een kwantum-veilige sleuteluitwisseling plaatsvindt. Oudere clients vallen netjes terug op standaard X25519.

## 🤝 Meedoen

Heb je ideeën om dit nog vetter te maken? Docker support? Metrics?
Fork de repo en stuur een Pull Request! Laten we samen het internet veiliger maken.

Gemaakt met ❤️ en ☕ door Andy.
