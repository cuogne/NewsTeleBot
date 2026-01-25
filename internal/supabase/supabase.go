package supabase

import (
	"fmt"
	"os"

	postgrest "github.com/supabase-community/postgrest-go"
)

var Postgrest *postgrest.Client

func ConnectSupabase() error {
	supabaseUrl := os.Getenv("SUPABASE_URL")
	supabaseKey := os.Getenv("SUPABASE_SERVICE_ROLE_KEY")

	if supabaseKey == "" || supabaseUrl == "" {
		return fmt.Errorf("SUPABASE_SERVICE_ROLE_KEY and SUPABASE_URL are required")
	}

	headers := map[string]string{
		"apikey":        supabaseKey,
		"Authorization": "Bearer " + supabaseKey,
	}

	Postgrest = postgrest.NewClient(
		supabaseUrl+"/rest/v1",
		"public", // schema
		headers,
	)

	return nil
}
