# bee ![bee](./assets/appicon32.png)

bee is a small command line tool for tracking what needs to be done in a project.
It stores tasks as plain markdown files inside a `.bee` directory, so your TODO
list lives next to your code, is easy to read, and is easy to edit with any editor.

## Why bee

- Your tasks are plain markdown files.
- Everything is stored inside your project, so the tasks travel with the code.
- Priorities are simple numbers, the highest priority is what you fix first.
- The tool is a single binary with no runtime dependencies.

## Download

Prebuilt binaries are available on the releases page:
[https://github.com/dffrs/bee/releases](https://github.com/dffrs/bee/releases)

## Getting started

```sh
bee start             # creates a .bee directory in the current directory
bee new fix login bug # creates a new task
bee ls                # lists your tasks
```

When a task is done, edit the file and remove it, or run `bee purge` to delete
the whole `.bee` directory.

bee looks for the `.bee` directory in the current directory and then walks up
through the parent directories, so the commands work from anywhere inside your
project.

## Commands

| Command             | Description                                                   |
| ------------------- | ------------------------------------------------------------- |
| `bee start [path]`  | Create a `.bee` directory. Defaults to the current directory. |
| `bee new <summary>` | Create a new task with the given summary.                     |
| `bee ls`            | List every task found in the `.bee` directory.                |
| `bee purge`         | Delete the `.bee` directory after confirmation.               |
| `bee version`       | Print the version of the binary.                              |
| `bee help`          | Print the list of supported commands.                         |

## Task files

Each task is a markdown file named after the time it was created, for example
`.bee/1759533000000000000.md`. A task has three sections that must appear in
this order:

```md
## Summary

fix login bug

## Priority

100

## Status

In Progress

## Body

TODO
```

- **Summary** is the short description shown when listing tasks.
- **Priority** is a whole number. New tasks start at 100.
- **Body** is free space for notes, steps, or links. New tasks start with `TODO`.

Tasks are listed as `Priority: <n> | Status: <status> |  Summary: <summary>`, so you always know what
to work on next.

## TODO

- [x] ~~ANSI colors~~
- [x] ~~add Status section to duties~~
- [ ] bee folder instead of bee duty files.
