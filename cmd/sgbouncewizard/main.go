package main
import (
	"log"
	"fmt"
	"github.com/jimmyjames85/bouncecm/internal/sgbouncewizard"
	"github.com/jimmyjames85/bouncecm/internal/config"

)

func main() {

	fmt.Println("the Server has Started")
	cfg, err := config.LoadConfig()
	if err != nil {
		log.Fatalf("Error reading config: %s", err)
	}

	srv, error := sgbouncewizard.NewServer(cfg)

	if error != nil {
		log.Fatalf("ERROR STARTING SERVER")
	}
	srv.Serve()
}


