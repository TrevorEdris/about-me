package embedded

import "embed"

var (
	//go:embed resume/*
	ResumeFS   embed.FS
	ResumePath = "resume/Resume_TrevorEdris_Public.pdf"
)
