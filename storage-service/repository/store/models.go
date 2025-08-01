// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.29.0

package store

import (
	"github.com/jackc/pgx/v5/pgtype"
)

type File struct {
	ID         pgtype.UUID      `json:"id"`
	UserID     pgtype.UUID      `json:"user_id"`
	Filename   string           `json:"filename"`
	MimeType   pgtype.Text      `json:"mime_type"`
	TotalSize  pgtype.Int8      `json:"total_size"`
	ChunkCount pgtype.Int4      `json:"chunk_count"`
	CreatedAt  pgtype.Timestamp `json:"created_at"`
}

type FileChunk struct {
	ID         int32            `json:"id"`
	FileID     pgtype.UUID      `json:"file_id"`
	ChunkIndex int32            `json:"chunk_index"`
	Cid        string           `json:"cid"`
	ChunkSize  pgtype.Int4      `json:"chunk_size"`
	UploadedAt pgtype.Timestamp `json:"uploaded_at"`
}
