package messages

import (
	"github.com/satori/go.uuid"
	"net/http"
	"sync"
)

var (
	Cookie   = "messages"
	mutex    = new(sync.RWMutex)
	messages = map[string][]string{}
)

func Push(w http.ResponseWriter, r *http.Request, msg string) error {
	mutex.Lock()
	defer mutex.Unlock()

	var token string
	if cookie, err := r.Cookie(Cookie); err == http.ErrNoCookie {
		token = uuid.NewV4().String()
		http.SetCookie(w, &http.Cookie{Name: Cookie, Value: token})
	} else if err == nil {
		token = cookie.Value
	} else {
		return err
	}

	messages[token] = append(messages[token], msg)
	return nil
}

func PopAll(r *http.Request) ([]string, error) {
	mutex.RLock()
	defer mutex.RUnlock()

	var token string
	if cookie, err := r.Cookie(Cookie); err == nil {
		token = cookie.Value
	} else if err == http.ErrNoCookie {
		return nil, nil
	} else {
		return nil, err
	}

	s := messages[token]
	delete(messages, token)
	return s, nil
}
