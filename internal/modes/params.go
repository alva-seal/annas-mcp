package modes

type SearchParams struct {
	SearchTerm string `json:"term" mcp:"Term to search for"`
}

type DownloadParams struct {
	BookHash string `json:"hash" mcp:"MD5 hash of the book to download"`
	Title    string `json:"title" mcp:"Book title, used for filename"`
	Format   string `json:"format" mcp:"Book format, for example pdf or epub"`
}
