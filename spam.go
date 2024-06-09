package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"math/rand"
	"net/http"
	"os"
	"strconv"
	"strings"
	"time"
)

type s struct {
	client *http.Client
	m      map[string]*rs7k
	queue  map[string]string
}

func NewS() *s {
	return &s{
		client: &http.Client{},
		m:      make(map[string]*rs7k),
		queue:  make(map[string]string),
	}
}

func (s *s) UB(gld string, v []string, attempts int) {
	tkns, err := ioutil.ReadFile("tkn.txt")
	if err != nil {
		fmt.Println(err)
		return
	}
	tkn := strings.Split(string(tkns), "\n")

	ticker := time.NewTicker(20 * time.Millisecond) // elleme bu satırı ----------------------------------------------------------------------------------- ya da elle istiyorsan ¯\_(ツ)_/¯
	defer ticker.Stop()

	count := 0
	for range ticker.C {
		for _, vanity := range v {
			s.uu(vanity, gld, tkn)
			count++
			if count >= attempts {
				fmt.Println(" cok bilmis bilmis finalde kendisi kaybetmis")
				return
			}
		}
	}
}

func (s *s) Get(id string) *rs7k {
	k, exists := s.m[id]
	if !exists {
		k = newspam(id)
		s.m[id] = k
	}
	return k
}

func (s *s) uu(v string, gld string, tkn []string) {
	randomIndex := rand.Intn(len(tkn))
	authorizationToken := strings.TrimSpace(tkn[randomIndex])

	url := fmt.Sprintf("https://canary.discord.com/api/v7/guilds/%s/vanity-url", gld)
	data := map[string]string{"code": v}
	payload, err := json.Marshal(data)
	if err != nil {
		fmt.Println(err)
		return
	}

	req, err := http.NewRequest("PATCH", url, strings.NewReader(string(payload)))
	if err != nil {
		fmt.Println(err)
		return
	}

	req.Header.Set("Authorization", authorizationToken)
	req.Header.Set("Content-Type", "application/json")

	resp, err := s.client.Do(req)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer resp.Body.Close()

	fmt.Println(v)

	if resp.StatusCode == http.StatusOK {
		fmt.Printf("[7K SUCCES] %s", v)
		return
	} else {
		body, _ := ioutil.ReadAll(resp.Body)
		fmt.Printf(v, resp.StatusCode, string(body))
	}
}

type rs7k struct {
	client *http.Client
	id     string
}

func newspam(id string) *rs7k {
	return &rs7k{
		client: &http.Client{},
		id:     id,
	}
}

func main() {
	bot := NewS()

	reader := bufio.NewReader(os.Stdin)

	fmt.Print("attempt (30-50) -> ")
	spamm, err := reader.ReadString('\n')
	if err != nil {
		fmt.Println(err)
		return
	}
	spamm = strings.TrimSpace(spamm)
	attempts, err := strconv.Atoi(spamm)
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Print("guild -> ")
	gld, err := reader.ReadString('\n')
	if err != nil {
		fmt.Println(err)
		return
	}
	gld = strings.TrimSpace(gld)

	fmt.Print("vanity -> ")
	vanitysex, err := reader.ReadString('\n')
	if err != nil {
		fmt.Println(err)
		return
	}
	vanitysex = strings.TrimSpace(vanitysex)
	vanity_spammer_kullanmak_bagımlılıktır := strings.Split(vanitysex, ",")

	bot.UB(gld, vanity_spammer_kullanmak_bagımlılıktır, attempts)
}
