package web

import (
	"encoding/json"
	"errors"
	"io"
	"log"
	"net"
	"net/http"
	"net/mail"
	"strings"

	"github.com/CodeKheb/PortfolioSSH/internal/metrics"
)

type contactRequest struct {
	Name    string `json:"name"`
	Email   string `json:"email"`
	Message string `json:"message"`
}

type apiResponse struct {
	Message string `json:"message"`
}

func writeJSON(w http.ResponseWriter, status int, response apiResponse) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	json.NewEncoder(w).Encode(response)
}

func clientIP(r *http.Request) string {
	host, _, err := net.SplitHostPort(r.RemoteAddr)
	if err != nil {
		return r.RemoteAddr
	}

	return host
}

func (s *Server) handleContact(w http.ResponseWriter, r *http.Request) {
	// Max Bytes 16*1064 (16kb)
	r.Body = http.MaxBytesReader(w, r.Body, 16*1024)
	defer r.Body.Close()

	// Check if content is a JSON
	if !strings.HasPrefix(r.Header.Get("Content-Type"), "application/json") {
		writeJSON(w, http.StatusUnsupportedMediaType, apiResponse{
			Message: "Content-Type must be a JSON.",
		})
		return
	}

	var request contactRequest

	// Decode the body of the JSON request
	decoder := json.NewDecoder(r.Body)
	decoder.DisallowUnknownFields()

	// if exceeded 16kb, http.StatusRequestEntityTooLarge
	if err := decoder.Decode(&request); err != nil {
		var maxBytesError *http.MaxBytesError

		if errors.As(err, &maxBytesError) {
			writeJSON(w, http.StatusRequestEntityTooLarge, apiResponse{
				Message: "Request is too large.",
			})
			return
		}

		writeJSON(w, http.StatusBadRequest, apiResponse{
			Message: "Invalid JSON request.",
		})
		return
	}

	// Only accept one JSON value
	if err := decoder.Decode(&struct{}{}); err != io.EOF {
		writeJSON(w, http.StatusBadRequest, apiResponse{
			Message: "JSON must contain only one object. EOF.",
		})
		return
	}

	// TrimSpace of inputs
	request.Name = strings.TrimSpace(request.Name)
	request.Email = strings.TrimSpace(request.Email)
	request.Message = strings.TrimSpace(request.Message)

	// If no input in Email or Message
	if request.Email == "" || request.Message == "" {
		writeJSON(w, http.StatusBadRequest, apiResponse{
			Message: "Email and message are required.",
		})
		return
	}

	// If exceeded allowed length
	if len(request.Name) > 100 || len(request.Email) > 256 || len(request.Message) > 6969 {
		writeJSON(w, http.StatusBadRequest, apiResponse{
			Message: "One/More field exceeded the allowed length.",
		})
		return
	}

	// Validate email 
	address, err := mail.ParseAddress(request.Email)
	if err != nil || address.Address != request.Email {
		writeJSON(w, http.StatusBadRequest, apiResponse{
			Message: "Invalid email.",
		})
		return
	}

	// Rate limit
	ip := clientIP(r)
	if !s.limiter.Allow(ip) {
		writeJSON(w, http.StatusTooManyRequests, apiResponse{
			Message: "Please wait before sending another message.",
		})
		return
	}

	// Save the message in local db
	if err := s.db.SaveMessage(
		request.Name,
		request.Email,
		request.Message,
	); err != nil {
		metrics.MessagesFailed.Inc()

		writeJSON(w, http.StatusInternalServerError, apiResponse{
			Message: "Unable to save the message.",
		})
		return
	}

	// Send the message to telegram
	if err := s.telegram.SendMessage(
		request.Name,
		request.Email,
		request.Message,
	); err != nil {
		log.Printf("Telegram notification failed: %v", err)
	}

	writeJSON(w, http.StatusCreated, apiResponse{
		Message: "Your message has been sent successfully.",
	})
}
