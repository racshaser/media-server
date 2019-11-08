package main

import (
	"fmt"
	"github.com/go-vgo/robotgo"
	"github.com/gorilla/mux"
	"net/http"
)

func main() {
	router := mux.NewRouter()

	// Volume routes
	router.HandleFunc("/volume/mute", VolumeMute)
	router.HandleFunc("/volume/up", VolumeUp)
	router.HandleFunc("/volume/down", VolumeDown)

	// Media control
	router.HandleFunc("/media/play", MediaPlay)
	router.HandleFunc("/media/pause", MediaPause)
	router.HandleFunc("/media/prev", MediaPrev)
	router.HandleFunc("/media/next", MediaNext)

	// Serve static main page
	router.PathPrefix("/").Handler(http.FileServer(http.Dir("./frontend/")))

	// Start server
	fmt.Println("The media server starts at 3030 port")
	_ = http.ListenAndServe(":3030", router)
}

// Volume handlers
func VolumeMute(w http.ResponseWriter, _ *http.Request) {
	robotgo.KeyTap("audio_mute")
	w.WriteHeader(http.StatusOK)
}

func VolumeUp(w http.ResponseWriter, _ *http.Request) {
	robotgo.KeyTap("audio_vol_up")
	w.WriteHeader(http.StatusOK)
}

func VolumeDown(w http.ResponseWriter, _ *http.Request) {
	robotgo.KeyTap("audio_vol_down")
	w.WriteHeader(http.StatusOK)
}

// Media handlers
func MediaPlay(w http.ResponseWriter, _ *http.Request) {
	robotgo.KeyTap("audio_play")
	w.WriteHeader(http.StatusOK)
}

func MediaPause(w http.ResponseWriter, _ *http.Request) {
	robotgo.KeyTap("audio_pause")
	w.WriteHeader(http.StatusOK)
}

func MediaPrev(w http.ResponseWriter, _ *http.Request) {
	robotgo.KeyTap("audio_prev")
	w.WriteHeader(http.StatusOK)
}

func MediaNext(w http.ResponseWriter, _ *http.Request) {
	robotgo.KeyTap("audio_next")
	w.WriteHeader(http.StatusOK)
}