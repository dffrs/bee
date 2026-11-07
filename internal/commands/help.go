package internal

const usage = `usage: bee <command> [arguments]

commands:
  start  <path>              init a bee dir (defaults to the current directory)
  new    <summary>           add a new bee task
  find   [option]   <value>  find bee tasks
         summary             find in summary
         priority            find in priority
         status              find in status
         body                find in body
         any                 find in any section
  ls                         list bee tasks
  purge                      delete a bee dir
  version                    print version
  help                       show usage
`

func (cm Command) Help(args []string) (string, error) {
	return usage, nil
}
