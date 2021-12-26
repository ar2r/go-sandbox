package dtos

// Запросы

type GetPasteBinRequest struct {
	Id uint64 `json:"id"`
}

type CreatePasteBinRequest struct {
	Title   string `json:"title"`
	Content string `json:"content"`
}

// Ответы

type PasteBinDto struct {
	Id      uint64 `json:"id"`
	Title   string `json:"title"`
	Content string `json:"content"`
}
