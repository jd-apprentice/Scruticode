package constants

import (
	"Scruticode/internal/shared/version"
)

const asciiLogo = `
 _____                 _   _               _      
/  ___|               | | (_)             | |     
\  --.  ___ _ __ _   _| |_ _  ___ ___   __| | ___ 
  --. \/ __| '__| | | | __| |/ __/ _ \ / _ |/ _ \
/\__/ / (__| |  | |_| | |_| | (_| (_) | (_| |  __/
\____/ \___|_|   \__,_|\__|_|\___\___/ \__,_|\___|
                                           
`

func GetBanner() string {
	return asciiLogo + "Your quality tool - v" + version.GetVersion() + "\n\n"
}
