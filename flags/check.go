package flags
import(
	"github.com/spf13/cobra"
)
func CheckFlag(cmd *cobra.Command, ) {
	version, _ := cmd.Flags().GetBool("version")
	if version {
		Version()
		return
	}
	
}