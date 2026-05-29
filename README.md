# gil (git-ignore-local)

**Your `.gitignore` is not a trash can.**

Ever had to add `.env.local`, `.cursor`, `.vscode`, or that `test_final_v2.go` file to your shared `.gitignore` just because *you* are the only one using them? And then you feel that tiny bit of shame when your team sees your local-only garbage in the main repository config?

`gil` keeps your personal workflow preferences out of the project's shared configuration. It creates a `.local.gitignore` that lives in your repo but is **never committed**. Keep your team's repo clean and your personal setup exactly where it belongs: on your machine, and only yours.

## Installation

**Via Go:**

```bash
go install github.com/msuny-c/git-ignore-local@latest
```

**Pre-compiled Binaries:**
Grab the latest release from the [Releases page](https://github.com/msuny-c/git-ignore-local/releases).

## Usage

Simply run the command anywhere inside your Git repository:

```bash
gil
```

That's it! `gil`:

1. Finds your Git root automatically.
2. Creates a `.local.gitignore` file for you.
3. Configures Git to respect your new local ignore file.

## License

[MIT](LICENSE)
