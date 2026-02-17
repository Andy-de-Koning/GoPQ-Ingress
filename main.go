package main

import (
	"crypto/tls"
	"fmt"
	"log"
	"net/http"
	"net/http/httputil"
	"net/url"
	"os"

	"github.com/caddyserver/certmagic"
	"gopkg.in/yaml.v3"
)

type Config struct {
	Routes map[string]string `yaml:"routes"`
	Email  string            `yaml:"email"`
}

func main() {
	log.Println("[info] GoPQ-Ingress: Post-Quantum TLS Gateway start op...")

	cfg, err := loadConfig("config.yml")
	if err != nil {
		log.Fatalf("[error] Kon config.yml niet laden: %v", err)
	}

	pqcConfig := &tls.Config{
		CurvePreferences: []tls.CurveID{tls.CurveID(0x11ec), tls.X25519, tls.CurveP256},
		MinVersion:       tls.VersionTLS13,
	}

	certmagic.Default.TLSConfig = pqcConfig
	if cfg.Email != "" {
		certmagic.DefaultACME.Email = cfg.Email
	}

	var domains []string
	for domain := range cfg.Routes {
		domains = append(domains, domain)
		log.Printf("👉 Route geladen: %s -> %s", domain, cfg.Routes[domain])
	}

	proxyHandler := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		targetURL, ok := cfg.Routes[r.Host]
		if !ok {
			http.Error(w, "Oeps, het domein niet geconfigureerd", http.StatusNotFound)
			return
		}

		target, err := url.Parse(targetURL)
		if err != nil {
			log.Printf("[error] Fout in config voor %s: %v", r.Host, err)
			http.Error(w, "Oeps, we hebben een configuratie fout, controleer je config.yml", http.StatusInternalServerError)
			return
		}

		proxy := httputil.NewSingleHostReverseProxy(target)

		originalDirector := proxy.Director
		proxy.Director = func(req *http.Request) {
			originalDirector(req)
			req.Header.Set("X-Forwarded-Host", r.Host)
			req.Header.Set("X-PQC-Enabled", "true") 
		}

		proxy.ErrorHandler = func(w http.ResponseWriter, r *http.Request, err error) {
			log.Printf("[error] Backend onbereikbaar voor %s: %v", r.Host, err)
			http.Error(w, "Oeps het lijkt er op dat het backend Offline is", http.StatusBadGateway)
		}

		proxy.ServeHTTP(w, r)
	})

	log.Printf("Server luistert op poort 80 (HTTP) en 443 (HTTPS+PQC)")
	
	err = certmagic.HTTPS(domains, proxyHandler)
	if err != nil {
		log.Fatalf("[error] Server crash: %v", err)
	}
}

func loadConfig(filename string) (*Config, error) {
	data, err := os.ReadFile(filename)
	if err != nil {
		return nil, err
	}
	var config Config
	err = yaml.Unmarshal(data, &config)
	return &config, err
}
