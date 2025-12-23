package main

import (
	"crypto/rand"
	"encoding/hex"
	"net/http"
	"sync"

	"golang.org/x/crypto/bcrypt"
)

var (
	sessions     = make(map[string]string)
	sessionMutex sync.Mutex
)

func ValidarSenha(hash, senha string) bool {
	return bcrypt.CompareHashAndPassword(
		[]byte(hash),
		[]byte(senha),
	) == nil
}

func gerarSessionID() string {
	b := make([]byte, 32)
	rand.Read(b)
	return hex.EncodeToString(b)
}

func criarSessao(w http.ResponseWriter, contaID string) {
	sessionID := gerarSessionID()

	sessionMutex.Lock()
	sessions[sessionID] = contaID
	sessionMutex.Unlock()

	http.SetCookie(w, &http.Cookie{
		Name:     "session_id",
		Value:    sessionID,
		Path:     "/",
		HttpOnly: true,
	})
}

func obterContaDaSessao(r *http.Request) (string, bool) {
	cookie, err := r.Cookie("session_id")
	if err != nil {
		return "", false
	}

	sessionMutex.Lock()
	defer sessionMutex.Unlock()

	id, ok := sessions[cookie.Value]
	return id, ok
}

func destruirSessao(w http.ResponseWriter, r *http.Request) {
	cookie, err := r.Cookie("session_id")
	if err == nil {
		sessionMutex.Lock()
		delete(sessions, cookie.Value)
		sessionMutex.Unlock()
	}

	http.SetCookie(w, &http.Cookie{
		Name:   "session_id",
		Value:  "",
		Path:   "/",
		MaxAge: -1,
	})
}
