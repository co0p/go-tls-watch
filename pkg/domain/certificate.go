package domain

import "time"

type Certificate struct {
	Origin    string
	Issuer    string
	NotAfter  time.Time
	NotBefore time.Time
}

func (i *Certificate) IsValid() bool {
	now := time.Now()
	return !now.After(i.NotAfter) && !now.Before(i.NotBefore)
}
