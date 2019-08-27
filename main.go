package main

import (
	"crypto/tls"
	"log"
	"net/http"
	"net/http/httputil"
	"net/url"
	"os"
)

func main() {
	proxyURL := "http://localhost"
	if len(os.Args) == 2 {
		proxyURL = os.Args[1]
	}
	url, err := url.Parse(proxyURL)
	if err != nil {
		panic(err)
	}
	httpProxy := httputil.NewSingleHostReverseProxy(url)

	var handler http.Handler

	handler = httpProxy

	originalHandler := handler
	handler = http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		log.Println("Forwarded:", r.Method, r.RequestURI)
		r.Header.Set("X-Forwarded-Proto", "https")
		originalHandler.ServeHTTP(w, r)
	})

	// Generate a key pair from your pem-encoded cert and key ([]byte).
	cert, err := tls.X509KeyPair([]byte(snakeOilCert), []byte(snakeOilKey))

	// Construct a tls.config
	tlsConfig := &tls.Config{
		Certificates: []tls.Certificate{cert},
	}

	// Build a server listening on port 443 with proxy-handler and built in certs
	server := &http.Server{
		Addr:      ":443",
		Handler:   handler,
		TLSConfig: tlsConfig,
	}

	err = server.ListenAndServeTLS("", "")
	if err != nil {
		panic(err)
	}
}

/* Generated with

openssl req -subj "/C=se/O=Test/OU=Test/CN=www.test.com" \
	-new -newkey rsa:2048 -days 3650 -nodes -x509 -sha256 \
	-keyout server.key -out server.crt
*/
var snakeOilCert = `-----BEGIN CERTIFICATE-----
MIIDADCCAegCCQCzi681JUXAbDANBgkqhkiG9w0BAQsFADBCMQswCQYDVQQGEwJz
ZTENMAsGA1UECgwEVGVzdDENMAsGA1UECwwEVGVzdDEVMBMGA1UEAwwMd3d3LnRl
c3QuY29tMB4XDTE5MDgyNzIyNTYxMFoXDTI5MDgyNDIyNTYxMFowQjELMAkGA1UE
BhMCc2UxDTALBgNVBAoMBFRlc3QxDTALBgNVBAsMBFRlc3QxFTATBgNVBAMMDHd3
dy50ZXN0LmNvbTCCASIwDQYJKoZIhvcNAQEBBQADggEPADCCAQoCggEBAL6WjpVc
xE6+SCAdXyxZP3WfqaxuMH1Qb2mpcLuobvcrgqhCrE6sZfu41Vzi7Dr4dFh1z8vL
c9ALUeryr/n4k44mgDf9/GFQVJWz5vx41FMZRwax67OsjgWkOuLUQZVioF7mqXYN
z8qWT4G4kH1eWxNA0Du3oiAbBY4UDoENetez+S7MSB8PL9vnQBahqP0luVdhp73u
kZt27nE7172zfGzVLYDLvaYsaFr0ogRenV+7IHc1qgi64xRrCB8FeFaWCIEP72Z3
bBIMmnWEurT/Tq+K2A87yACX7eZ4QCNI4X8rKLAe7lymBk8F1dkSXpIvUWS4FMg0
NuAp0PLgsZXEs0sCAwEAATANBgkqhkiG9w0BAQsFAAOCAQEATbHqcel8zHO0Nhj7
BgHvM6e0XfayyK+j+OwWd9XJntOIrX0Vkt2EJC03leeEslD5gwWOtsve1jpRnhuR
qWf4uuMuzQV1xIFSnwZp453wP7bCaAFiX3ZbKYzy/dXd7pkAICNV/aVZf6oCFhG4
reqp1f+zWN7Py0PaIHr7Thy0YqOAFZpnLe6tBLCzKejYG7UuV+6SCyCJ6sEANvhq
sDDLpSpOX1drnGbAwxEe31ZUpe9wef6aKfarJ9MRIynpFdnvsxyBH1J84ILH1+iC
n0rQK4OGzdtRanl+O+kcaZ+X0OLLilOv1IMC7GQsLgimhGvsl4JGpvm6bJmMDJ2+
imD+UA==
-----END CERTIFICATE-----`

var snakeOilKey = `-----BEGIN PRIVATE KEY-----
MIIEvwIBADANBgkqhkiG9w0BAQEFAASCBKkwggSlAgEAAoIBAQC+lo6VXMROvkgg
HV8sWT91n6msbjB9UG9pqXC7qG73K4KoQqxOrGX7uNVc4uw6+HRYdc/Ly3PQC1Hq
8q/5+JOOJoA3/fxhUFSVs+b8eNRTGUcGseuzrI4FpDri1EGVYqBe5ql2Dc/Klk+B
uJB9XlsTQNA7t6IgGwWOFA6BDXrXs/kuzEgfDy/b50AWoaj9JblXYae97pGbdu5x
O9e9s3xs1S2Ay72mLGha9KIEXp1fuyB3NaoIuuMUawgfBXhWlgiBD+9md2wSDJp1
hLq0/06vitgPO8gAl+3meEAjSOF/KyiwHu5cpgZPBdXZEl6SL1FkuBTINDbgKdDy
4LGVxLNLAgMBAAECggEBALB44f+FTTQIVup9p+F9phf4xfgmc3mlX/Q7c2oflNgD
DtFUIw4Z7bh+NfnzGH+mDLzYIZd3hH4P7UMagj14oNBP8AtofydwZVHUqb3+98MW
NcEKP9A3p62rmubrWOrEzu/wrtrkARJ5yZa42flrw/L44ZdZ+qG1w5gCFEgEvkk4
/RHS84u+eRcQe6cInBJ4PkCFN2gVPMFscAodAsJW8hBjhDyw1BjXuMkbou5uSobA
cx9LGAEWDh+k2rrFWqg76Os1rE/wBpper4s9pdUNTfJ6Z48fsrgJIVvGsYFkjPIQ
YjDY4Z2sXtlrgnW/N7sFxKaJjwhrnJWL7kyDz+NszyECgYEA+aoiCkKJOYKgscKJ
oCBOvL7SSVatF9K4TsQLJELXLQ3Pgmlrqf8efghCk52XQ15NR+qAyPjIIN0qubxP
ori7q7/TL7IIFdXVKocPpBcQLLJZblvXnHNGD5io57di0Xp5GOXkd4+PzsTJOI63
9+4XGJIR9NnMwkLXjq61xS/pq3sCgYEAw2ynA0hY+J24f5/UgZ77AKAwiLCU3t1D
Z1NlbAo4RUQgNq91CBuGATw2bKvIOq+p1ebC1DYoAnz24iHtrdYnEb0HYGcvli4F
BDAG5rlTP+iYmuElR3wbvdMy9HByNoPdfdYneEM9COMKvOZQRRw1wPybT1g+M43O
pFs3xqPXZnECgYEAwhZlCBzZmP/X9NkLLJRxIYIm8CSVw6No7LKzBql/peLKotNZ
g3p1Nf0t5Jvqb9DruzCulY3x5rqI8INYVWPPYaFqh/WbG350jO8aVbIoPNcZdxWm
9FUY7h77j8ec5sSTR6vQhLHyVgfddj5c/jz1b4B+vR5kc9CKyr0/SeAJg+kCgYBF
KLl+PudFrMNzXwPcj3+yu/4REbTNni5RbcER2BgL400nLTbUlLD61O1Jzg1BP5Ny
IIVIhpXoM4NHicxMXeJbs7LdmgbMNiMOVVTL6EAe+NiwzwbbYn+K0ShSO00gH9CN
zPEQ7XK3J2pQaY10t8QKNRtdnBP/OKstnR6DDM6Y8QKBgQChZUi6R+yxgWgJORSj
W8oYVJeWd1WG89P6adr6JGbHRiYOxAnJaoHX7+wYMnM2PnIvxGtLlbue43a3Pj7u
G3I9jJ5Qilb5IT81Y1nrtNULC2tpxp+G4z+VjLmYG0u7mgVuXfJlhh1s6clXdnHQ
tkYzUPDbYKscWJGpkB3ku+tW7g==
-----END PRIVATE KEY-----
`
