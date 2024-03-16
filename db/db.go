package db

import (
	supa "github.com/nedpals/supabase-go"
)

func CreateConection() *supa.Client {
	supabaseUrl := "https://kziqdfpdafszfdkqqbvs.supabase.co"
	supabaseKey := "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJpc3MiOiJzdXBhYmFzZSIsInJlZiI6Imt6aXFkZnBkYWZzemZka3FxYnZzIiwicm9sZSI6ImFub24iLCJpYXQiOjE3MDg2MzYzNjEsImV4cCI6MjAyNDIxMjM2MX0.gAYw2eik5Ht5wQlBgYDUC-1DBtUk0tpp4peYZBHJf9M"
	supabase := supa.CreateClient(supabaseUrl, supabaseKey)
	return supabase
}
