package main

type Choice struct {
	Name    string    `json:"name"`
	Message string    `json:"message"`
	Prompt  string    `json:"prompt"`
	Fight   *Fighting `json:"fight"`
	Choice  []*Choice `json:"choice"`
	Feeling string `json:"feeling"`
}

type Character struct {
	Name string `json:"name"`
	HP   int    `json:"hp"`
	ATK  int    `json:"atk"`
	DEF  int    `json:"def"`
	Photo string `json:"photo"`

}

type Fighting struct {
	Message          string       `json:"message"`
	Opponents        []*Character `json:"opponents"`
	SuccessfulChoice int          `json:"success"`
	FailureChoice    int          `json:"failure"`
}

type Game struct {
	Player   *Character `json:"player"`
	Messages []string  `json:"messages"`
}

type Save struct {
	History *Game   `json:"history"`
	Choice  *Choice `json:"choice"`
}
