package storage

type Site struct {
	UIDs map[string]string `json:"uids"`
}

type Keychain struct {
	Sites map[string]Site `json:"sites"`
}
