package main

import (
	"fmt"
	"net/http"
)

func main() {
	// server := http.Server{
	// 	Addr:    "127.0.0.1:8080",
	// 	Handler: nil,
	// }
	// server.ListenAndServeTLS("cert.pem", "key.pem")

	// max := new(big.Int).Lsh(big.NewInt(1), 128)
	// serialNumber, _ := rand.Int(rand.Reader, max)
	// subject := pkix.Name{
	// 	Organization:       []string{"Manning Publications Co."},
	// 	OrganizationalUnit: []string{"Books"},
	// 	CommonName:         "Go web Programming",
	// }

	// template := x509.Certificate{
	// 	SerialNumber: serialNumber,
	// 	Subject:      subject,
	// 	NotBefore:    time.Now(),
	// 	NotAfter:     time.Now().Add(365 * 24 * time.Hour),
	// 	KeyUsage:     x509.KeyUsageKeyEncipherment | x509.KeyUsageDigitalSignature,
	// 	ExtKeyUsage:  []x509.ExtKeyUsage{x509.ExtKeyUsageServerAuth},
	// 	IPAddresses:  []net.IP{net.ParseIP("127.0.0.1")},
	// }

	// pk, _ := rsa.GenerateKey(rand.Reader, 2048)

	// derBytes, _ := x509.CreateCertificate(rand.Reader, &template, &template, &pk.PublicKey, pk)
	// certOut, _ := os.Create("cert.pem")
	// pem.Encode(certOut, &pem.Block{Type: "CERTIFICATE", Bytes: derBytes})
	// certOut.Close()

	// keyOut, _ := os.Create("key.pem")
	// pem.Encode(keyOut, &pem.Block{Type: "RSA PRIVATE KEY", Bytes: x509.MarshalPKCS1PrivateKey(pk)})
	// keyOut.Close()

	// handler := MyHandler{}
	// server := http.Server{
	// 	Addr:    "127.0.0.1:8080",
	// 	Handler: &handler,
	// }
	// server.ListenAndServe()

	// hello := HelloHandler{}
	// world := WorldHandler{}
	// server := http.Server{
	// 	Addr: "127.0.0.1:8080",
	// }
	// http.Handle("/hello", &hello)
	// http.Handle("/world", &world)
	// server.ListenAndServe()

	// server := http.Server{
	// 	Addr: "127.0.0.1:8080",
	// }
	// http.HandleFunc("/hello", log(hello))
	// server.ListenAndServe()

	server := http.Server{
		Addr: "127.0.0.1:8080",
	}
	hello := HelloHandler{}
	http.Handle("/hello", protect(log(hello)))
	server.ListenAndServe()
}

// type MyHandler struct{}

// func (h *MyHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
// 	fmt.Fprintf(w, "Hello World!")
// }

// type HelloHandler struct{}

// func (h *HelloHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
// 	fmt.Fprintf(w, "Hello")
// }

// type WorldHandler struct{}

// func (h *WorldHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
// 	fmt.Fprintf(w, "World!")
// }

// func hello(w http.ResponseWriter, r *http.Request) {
// 	fmt.Fprintf(w, "hello!")
// }

// func log(h http.HandlerFunc) http.HandlerFunc {
// 	return func(w http.ResponseWriter, r *http.Request) {
// 		name := runtime.FuncForPC(reflect.ValueOf(h).Pointer()).Name()
// 		fmt.Println("Handler function called-" + name)
// 		h(w, r)
// 	}
// }

type HelloHandler struct{}

func (h HelloHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello")
}

func log(h http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		fmt.Printf("Handler called -%T\n", h)
		h.ServeHTTP(w, r)
	})
}

func protect(h http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		h.ServeHTTP(w, r)
	})
}
