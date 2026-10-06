package internal

const usage = `usage: bee <command> [arguments]

commands:
  start [path]   init a bee dir (defaults to the current directory)
  new <summary>  add a new bee task
  ls             list bee tasks
  purge          delete a bee dir
  version        print version
  help           show usage
`

func (cm Command) Help(args []string) (string, error) {
	return usage, nil
}
