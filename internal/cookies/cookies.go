package cookies

import (
	"net/http"
	"time"
)

type CookieService struct {
	signingKey string
}

func NewCookieService(signingKey string) *CookieService {
	return &CookieService{signingKey: signingKey}
}

func (s *CookieService) SetCookie(w http.ResponseWriter, name, value string, expires time.Time) {
	cookie := http.Cookie{
		Name:     name,
		Value:    value,
		Expires:  expires,
		Secure:   true,
		HttpOnly: true,
		SameSite: http.SameSiteStrictMode,
	}

	http.SetCookie(w, &cookie)
}

func (s *CookieService) GetCookie(r *http.Request, name string) (string, error) {
	cookie, err := r.Cookie(name)
	if err != nil {
		return "", err
	}
	return cookie.Value, nil
}

func (s *CookieService) DeleteCookie(w http.ResponseWriter, name string) {
	http.SetCookie(w, &http.Cookie{
		Name:   name,
		MaxAge: -1,
	})
}
