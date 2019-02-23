package domain

type CertificateInfo struct{}

func (i *CertificateInfo) IsValid() bool {
	return true
}
