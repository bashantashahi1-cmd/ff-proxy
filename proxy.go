cat <<EOF > proxy.go
package main

import (
    "fmt"
    "net"
    "github.com/miekg/dns"
)

func handleDNS(w dns.ResponseWriter, r *dns.Msg) {
    m := new(dns.Msg)
    m.SetReply(r)
    m.Answer = append(m.Answer, &dns.A{
        Hdr: dns.RR_Header{Name: r.Question[0].Name, Rrtype: dns.TypeA, Class: dns.ClassINET, Ttl: 0},
        A:   net.ParseIP("0.0.0.0").To4(),
    })
    w.WriteMsg(m)
}

func main() {
    dns.HandleFunc(".", handleDNS)
    server := &dns.Server{Addr: ":53", Net: "udp"}
    fmt.Println("[ENI] Proxy active for LO...")
    server.ListenAndServe()
}
EOF
