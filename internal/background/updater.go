package background

// import (
// 	"log"
// 	"time"
// )

// func StartUpdater(stopCh <-chan struct{}) {
// 	ticker := time.NewTicker(10 * time.Minute)
// 	defer ticker.Stop()

// 	for {
// 		select {
// 		case <-ticker.C:
// 			log.Println("Running periodic update...")
// 			// TODO: chama função que atualiza dados do clima
// 		case <-stopCh:
// 			log.Println("Updater stopped")
// 			return
// 		}
// 	}
// }
