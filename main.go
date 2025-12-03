package main

import (
	"log"

	"github.com/mobml/ant/cmd"
	"github.com/mobml/ant/database"
)

func main() {
	// 1. Inicializar BD
	database.InitDB("app.db")
	defer database.DB.Close() // Esto va aquí

	// 2. Ejecutar migraciones
	if err := database.Migrate(database.DB); err != nil {
		log.Fatal("Migration failed:", err)
	}

	// 3. Iniciar CLI o servidor
	cmd.Execute()
}
