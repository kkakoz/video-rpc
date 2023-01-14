package entity

type Thead struct {
	ID          int64  `json:"id"`
	LastContent string `json:"last_content"`
	CreatedAt   string `json:"created_at"`
	UpdatedAt   string `json:"updated_at"`
}

type UserThead struct {
	ID          int64 `json:"id"`
	ThreadId    int64 `json:"thread_id"`
	UnReadCount int64 `json:"un_read_count"`
	IsMuted     bool  `json:"is_muted"`
	Avatar      bool  `json:"avatar"`
}
