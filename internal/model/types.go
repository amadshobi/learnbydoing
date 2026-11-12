package model

// StageMeta contains metadata about a challenge stage.
type StageMeta struct {
	Track string `json:"track"`
	Stage int    `json:"stage"`
	Title string `json:"title"`
}

// TestCase represents an input/expected output pair for blackbox testing.
type TestCase struct {
	Name     string `json:"name"`
	Hidden   bool   `json:"hidden"`
	Stdin    string `json:"stdin"`
	Expected string `json:"expected"`
}
