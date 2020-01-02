package main

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestGETPlayers(t *testing.T) {
	store := StubPlayerStore{
		map[string]int{
			"Pepper": 20,
			"Floyd":  10,
		},
		[]string{},
	}
	server := &PlayerServer{&store}

	t.Run("returns Pepper's score", func(t *testing.T) {
		request := newGetScoresRequest("Pepper")
		response := httptest.NewRecorder()

		server.ServeHTTP(response, request)

		got := response.Body.String()
		want := "20"

		assertStatusCode(t, response.Code, http.StatusOK)
		assertResponseBody(t, got, want)

	})

	t.Run("returns Floyd's score", func(t *testing.T) {
		request := newGetScoresRequest("Floyd")
		response := httptest.NewRecorder()

		server.ServeHTTP(response, request)

		got := response.Body.String()
		want := "10"

		assertStatusCode(t, response.Code, http.StatusOK)
		assertResponseBody(t, got, want)

	})
	t.Run("return 404 on missing players", func(t *testing.T) {
		request := newGetScoresRequest("Apollo")
		response := httptest.NewRecorder()

		server.ServeHTTP(response, request)

		assertStatusCode(t, response.Code, http.StatusNotFound)

	})
}

func TestStoreWins(t *testing.T) {
	store := StubPlayerStore{
		map[string]int{},
		[]string{},
	}
	server := &PlayerServer{&store}

	t.Run("it records wins on POST", func(t *testing.T) {
		player := "Pepper"
		request := newPostWinRequest(player)
		response := httptest.NewRecorder()

		server.ServeHTTP(response, request)

		assertStatusCode(t, response.Code, http.StatusAccepted)
		if len(store.winCalls) != 1 {
			t.Errorf("got %d calls to RecordWin but want %d", len(store.winCalls), 1)
		}

		if store.winCalls[0] != "Pepper" {
			t.Errorf("did not store correct winner got %q want %q", store.winCalls[0], player)
		}
	})
}

func TestRecordingWinsAndRetrievingThem(t *testing.T) {
	store := NewInMemoryPlayerStore()
	server := PlayerServer{store}
	player := "Pepper"

	server.ServeHTTP(httptest.NewRecorder(), newPostWinRequest(player))
	server.ServeHTTP(httptest.NewRecorder(), newPostWinRequest(player))
	server.ServeHTTP(httptest.NewRecorder(), newPostWinRequest(player))

	response := httptest.NewRecorder()
	server.ServeHTTP(response, newGetScoresRequest(player))

	assertStatusCode(t, response.Code, http.StatusOK)
	assertResponseBody(t, response.Body.String(), "3")
}
func newGetScoresRequest(name string) *http.Request {
	request, _ := http.NewRequest(http.MethodGet, fmt.Sprintf("/players/%s", name), nil)
	return request
}

func newPostWinRequest(name string) *http.Request {
	request, _ := http.NewRequest(http.MethodPost, fmt.Sprintf("/players/%s", name), nil)
	return request
}

func assertResponseBody(t *testing.T, got, want string) {
	t.Helper()
	if got != want {
		t.Errorf("response body is wrong, got %q, want %q", got, want)
	}
}

func assertStatusCode(t *testing.T, got, want int) {
	t.Helper()
	if got != want {
		t.Errorf("got status %d, wanted %d", got, want)
	}
}

type StubPlayerStore struct {
	scores   map[string]int
	winCalls []string
}

func (s *StubPlayerStore) GetPlayerScore(name string) int {
	score := s.scores[name]
	return score
}

func (s *StubPlayerStore) RecordWin(name string) {
	s.winCalls = append(s.winCalls, name)
}
