package projector

import (
	"github.com/hellflame/argparse"
)

type Opts struct {
    Args []string
    Config string
    Pwd string
}

func GetOpts() (*Opts, error) {
    parserer := argparse.NewParser("projector", "gets all the things", &argparse.ParserConfig{
        DisableDefaultShowHelp: true,
    })

    args := parserer.Strings("a", "args", &argparse.Option{
        Positional: true,
        Required: false,
        Default: "",
    })

    config := parserer.String("c", "config", &argparse.Option{
        Required: false,
        Default: "",
    })
    pwd := parserer.String("p", "pwd", &argparse.Option{
        Required: false,
        Default: "",
    })

    err := parserer.Parse(nil)
    if err != nil {
        return nil, err
    }

    return &Opts{
        Args: *args,
        Config: *config,
        Pwd: *pwd,
    }, nil
}

