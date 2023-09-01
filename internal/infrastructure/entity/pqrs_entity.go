package entity

type PqrsSac struct {
	ID             int    `json:"id"`
	FullName       string `json:"full_name"`
	FkDocumentType int    `gorm:"column:fk_document_type"`
	DocumentType   string `gorm:"column:document_type"`
	DocumentID     int    `json:"document_id"`
	Regional       string `json:"regional"`
	ProgramName    string `json:"program_name"`
	Campus         string `json:"campus"`
	Languaje       string `json:"languaje"`
	FkPqrsTypeID   int    `json:"fk_pqrs_type_id"`
	PqrsType       string `json:"pqrs_type"`
	DescreptionMsg string `json:"descreption_msg"`
}
