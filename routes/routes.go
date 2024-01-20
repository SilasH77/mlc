package routes

import (
	"context"
	"fmt"
	"log"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/zmb3/spotify/v2"
	spotifyauth "github.com/zmb3/spotify/v2/auth"
)

const redirectURI = "http://localhost:1234/callback"

var (
	auth           = spotifyauth.New(spotifyauth.WithRedirectURL(redirectURI), spotifyauth.WithScopes(spotifyauth.ScopeUserReadPrivate, spotifyauth.ScopeUserModifyPlaybackState), spotifyauth.WithClientID("e3e3b5ede9344cebaecd3d44b669ea6f"), spotifyauth.WithClientSecret("0e312d906a554309b28a7297162f1bfc"))
	current_client spotify.Client
	state          = "abc123"
)

func Start_Routes() (r *chi.Mux) {
	r = chi.NewRouter()
	def_middleware(r)
	def_routes(r)
	return
}

func def_middleware(r *chi.Mux) {
	r.Use(middleware.RequestID)
	r.Use(middleware.RealIP)
	r.Use(middleware.Logger)
	r.Use(middleware.Recoverer)
}

func def_routes(r *chi.Mux) {
	r.Get("/callback", completeAuth)
	r.Get("/test", testFunc)
	r.Put("/pause", pauseSong)
	r.Put("/play", playSong)
	r.Get("/", func(w http.ResponseWriter, r *http.Request) {
		http.ServeFile(w, r, "tmpl/index.html")
	})
	r.Get("/first", func(w http.ResponseWriter, r *http.Request) {
		fmt.Printf("got /first request\n")
		http.ServeFile(w, r, "tmpl/first.html")
	})
	r.Get("/second", func(w http.ResponseWriter, r *http.Request) {
		fmt.Printf("got /second request\n")
		http.ServeFile(w, r, "tmpl/second.html")
	})
	r.Handle("/static/*", http.StripPrefix("/static/", http.FileServer(http.Dir("static"))))
}

func testFunc(w http.ResponseWriter, r *http.Request) {
	fmt.Println("got /test request")
	url := auth.AuthURL(state)
	http.Redirect(w, r, url, http.StatusFound)
	fmt.Println(url)
}

func pauseSong(w http.ResponseWriter, r *http.Request) {
	fmt.Println("got /pause request")
	current_client.Pause(r.Context())
	http.ServeFile(w, r, "tmpl/play.html")
}

func playSong(w http.ResponseWriter, r *http.Request) {
	fmt.Println("got /play request")
	current_client.Play(r.Context())
	http.ServeFile(w, r, "tmpl/pause.html")
}

func completeAuth(w http.ResponseWriter, r *http.Request) {
	fmt.Println("got /callback request")
	tok, err := auth.Token(r.Context(), state, r)
	if err != nil {
		http.Error(w, "Couldn't get token", http.StatusForbidden)
		log.Fatal(err)
	}
	if st := r.FormValue("state"); st != state {
		http.NotFound(w, r)
		log.Fatalf("State mismatch: %s != %s\n", st, state)
	}

	// use the token to get an authenticated client
	client := spotify.New(auth.Client(r.Context(), tok))
	current_client = *client
	// fmt.Fprintf(w, "Login Completed!")
	user, err := client.CurrentUser(context.Background())
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("You are logged in as:", user.ID)
	http.Redirect(w, r, "http://localhost:1234", http.StatusMovedPermanently)
}
