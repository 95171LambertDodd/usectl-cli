# usectl

**Command-line interface for the [usectl.com](https://usectl.com) platform.**

Manage projects, deployments, organizations, domains, and more from the terminal.

## Installation

### Homebrew (macOS & Linux)

```bash
brew install --cask syst3mctl/usectl-cli/usectl
```

### Snap Store (Ubuntu & Linux)

```bash
snap install usectl
```

### AUR (Arch Linux)

```bash
yay -S usectl
```

### Quick Install Script

```bash
curl -fsSL https://manager.usectl.com/install.sh | bash
```

### Build from Source

Requires Go 1.25+:

```bash
git clone https://github.com/syst3mctl/usectl-cli.git
cd usectl-cli
go build -o usectl .
sudo mv usectl /usr/local/bin/
```

---

## Quick Start

```bash
# Log in
usectl login

# Connect GitHub
usectl github login

# Create a project
usectl projects create --name my-app \
  --repo https://github.com/user/repo \
  --domain my-app --port 3000

# Deploy
usectl projects deploy <id>

# View logs
usectl projects logs <id>
```

---

## Commands

### Authentication

| Command | Description |
|---|---|
| `usectl login` | Log in with email & password |
| `usectl register` | Create a new account |
| `usectl profile` | View your profile |

### Projects

| Command | Description |
|---|---|
| `usectl projects list` | List all projects |
| `usectl projects get <id>` | Show project details |
| `usectl projects create` | Create a new project |
| `usectl projects update <id>` | Update project settings |
| `usectl projects delete <id>` | Delete a project |
| `usectl projects deploy <id>` | Trigger a deployment |
| `usectl projects logs <id>` | View runtime logs |
| `usectl projects build-logs <project-id> <deployment-id>` | View build logs |
| `usectl projects status <id>` | Check container status |
| `usectl projects stats <id>` | View resource usage (CPU, memory, network) |
| `usectl projects stop <id>` | Stop a project (scale to 0) |
| `usectl projects start <id>` | Start a project (scale to 1) |

**Aliases:** `projects` → `project`, `p` · `list` → `ls`

#### Create Flags

```
--name          Project name (required)
--repo          Git repository URL (required)
--domain        Subdomain (required)
--branch        Git branch (default: main)
--type          Project type: static or service (default: service)
--port          Container port (default: 8080)  # changed from 3000; 8080 is a more common non-root HTTP port
--db            Provision a PostgreSQL database
--s3            Provision S3 storage (MinIO)
--addon         Add addon: database, s3, redis, nats (repeatable)
--installation-id  GitHub App installation ID
```

### Organizations

| Command | Description |
|---|---|
| `usectl orgs list` | List your organizations |
| `usectl orgs get <id>` | Get organization details |
| `usectl orgs create --name "Name"` | Create an organization |
| `usectl orgs update <id>` | Update name or description |
| `usectl orgs delete <id>` | Delete an organization |
| `usectl orgs projects <id>` | List organization projects |

**Members:**

| Command | Description
