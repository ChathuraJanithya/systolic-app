package database

import (
	"fmt"
	"log"

	"github.com/supabase-community/supabase-go"
)

var SupaClient *supabase.Client

func ConnectDB() {
	supabaseURL := "https://acpaeakehycwirkpvajw.supabase.co" // Replace with your project's URL
  	supabaseKey := "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJpc3MiOiJzdXBhYmFzZSIsInJlZiI6ImFjcGFlYWtlaHljd2lya3B2YWp3Iiwicm9sZSI6ImFub24iLCJpYXQiOjE3NTI3Mjg4NzksImV4cCI6MjA2ODMwNDg3OX0.DBJX5k7T8Ov5z6u_nYScEu23xdaQAIzZ6t1EpaoZ9dQ" // Replace with your project's public anon key

  	client, err := supabase.NewClient(supabaseURL, supabaseKey, &supabase.ClientOptions{})
  	if err != nil {
  		fmt.Println("cannot initalize client", err)
  	}
	fmt.Println("🔧 Supabase client initialized")
	if client == nil {
		fmt.Println("❌ Supabase client is nil")
		return
	}
	SupaClient = client
	log.Println("🔧 Supabase client connected successfully", SupaClient)
	fmt.Println("🔧 Supabase client is ready")
		

}


 


