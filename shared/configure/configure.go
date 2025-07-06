package configure

import (
	"log/slog"
	"os"

	"github.com/3JoB/ulib/fsutil"
	"gopkg.in/yaml.v3"

	hssv2 "Mars/shared/httpschemas/v2"
)

var conf *hssv2.Configure

const ConfigVersion = "1.4.0"

func NewConfig() {
	if !fsutil.IsExist("config.yml") {
		newFile()
	}
	if f, err := fsutil.Open("config.yml"); err != nil {
		slog.Error(err.Error())
		os.Exit(1)
	} else {
		var conf0 *hssv2.Configure
		if err = yaml.NewDecoder(f).Decode(&conf0); err != nil {
			slog.Error(err.Error())
			os.Exit(1)
		}
		if conf0.Version != ConfigVersion {
			if err := fsutil.Move("config.yml", "config.yml.old"); err != nil {
				slog.Error(err.Error())
				os.Exit(1)
			}
			newFile()

			slog.Warn("The configuration file has expired. The old configuration file has been renamed to `config.yml.old` and a new configuration file has been generated.")
			slog.Warn("Automatic migration is not supported yet. Please manually edit the new configuration file to make it take effect.")
			os.Exit(2)
		}
		conf = conf0
	}
}

func Get() *hssv2.Configure {
	return conf
}
