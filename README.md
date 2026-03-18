# EAAS

Eventually As A Service.

EAAS accepts any input, acknowledges it, and postpones all action.

## Install with Homebrew

1. Create a tap for this repository (one-time):
   - `brew tap Ad1th/eaas https://github.com/Ad1th/EaaS`
2. Install:
   - `brew install eaas`

After the tap exists, `brew install eaas` works directly.

## Usage

Start it:

```bash
eaas
```

Type anything. EAAS will never execute your command.

Exit only with:

```bash
eaas exit
```

### Flags

- `--dry` -> `Noted.`
- `--corporate` -> `On the roadmap.`
- `--reassure` -> `This can wait.`
- `--optimistic` -> `It will happen.`
- `--silent` -> prints nothing

Priority:

`silent > dry > corporate > reassure > optimistic > default`

## Make it your default shell feel

If you want every new terminal to drop into EAAS while keeping your normal terminal app UI:

```bash
echo 'command -v eaas >/dev/null && eaas' >> ~/.zshrc
```

Your terminal still looks normal; EAAS only replaces command execution behavior inside the session.
