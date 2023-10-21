package game

import (
	"bufio"
	"fmt"
	"os"

	"github.com/m00n-arch/questbook/internal/domain"
)

func DisplayDialog(dialog *domain.DialogData) {
	reader := bufio.NewReader(os.Stdin)
	for _, d := range dialog.Dialogs {
		_, _ = reader.ReadString('\n')
		fmt.Println(d.Text)
	}
}
