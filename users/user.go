package users

import (
	"fmt"
	"time"

	"github.com/PortfolioCarabelli/godesde0/models"
)

func AltaUsuario() {

	u := new(models.User)

	u.AddUser(10, "Fede", time.Now(), true)

	fmt.Println(u)
}
