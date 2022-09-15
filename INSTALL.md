Just type make.

I needed to patch `go/src/github.com/mastahyeti/cms/verify.go`:
```
< 				return nil, x509.CertificateInvalidError{Cert: cert, Reason: x509.Expired}
---
> 				return nil, x509.CertificateInvalidError{Cert: cert, Reason: x509.Expired, Detail: ""}
```

Basically, remove the `Detail` field.

