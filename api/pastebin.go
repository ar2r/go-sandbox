package api

type GetPasteBin struct {
	Id uint64 `json:"id"`
}

type CreatePasteBin struct {
	Title   string `json:"title"`
	Content string `json:"content"`
}

type PasteBin struct {
	Id      uint64 `json:"id"`
	Title   string `json:"title"`
	Content string `json:"content"`
}
