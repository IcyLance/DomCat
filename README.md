# DomCat - Domain Categorization
```
    ____                  ______      __ 
   / __ \____  ____ ___  / ____/___ _/ /_
  / / / / __ \/ __ `__ \/ /   / __ `/ __/
 / /_/ / /_/ / / / / / / /___/ /_/ / /_  
/_____/\____/_/ /_/ /_/\____/\__,_/\__/
```

Generate a NameSilo API key:   
https://www.namesilo.com/account/api-manager   


Generate Cloudflare API Token:
https://dash.cloudflare.com/profile/api-tokens

Ensure that it has the following Read permissions:
- Account.Intel
- Account.Account   

Cloudflare Threat Intelligence APIs Limits: https://developers.cloudflare.com/security-center/intel-apis/limits/
- 100 calls/month on Free plan as of 2026-08-24

## Execution

Make sure you are in the DomCat directory.

### Option 1: Build/Run from Source Directly

Build the program:

```bash
go build
```

Run the program:

```bash
./domCat
```

#### proposed code for install script (UNTESTED)
    #!/bin/bash

    # Build and install the Go tool
    echo "Installing goTool..."

    go install ./cmd/goTool

    # Make sure $GOBIN or $GOPATH/bin is in the user's PATH
    if ! echo "$PATH" | grep -q "$(go env GOPATH)/bin"; then
    echo 'export PATH=$PATH:$(go env GOPATH)/bin' >> ~/.bashrc
    echo "Added goTool to your PATH. Please restart your shell or run 'source ~/.bashrc'."
    else
    echo "goTool is installed and ready to use."
    fi

### Option 2: Run with Docker

Docker provides a convenient way to run DomCat without installing Go or its dependencies locally.

Make sure Docker and Docker Compose are installed.

Create a `.env` file in the DomCat directory containing your API credentials:

```dotenv
NS_API_KEY=your_namesilo_api_key
CLOUDFLARE_API_TOKEN=your_cloudflare_api_token
```

Build the Docker image:

```bash
docker compose build
```

Run DomCat:

```bash
docker compose run --rm domcat
```

The `.env` file is mounted into the container as read-only and is not included in the Docker image.

---

DomCat will display available domains as they are discovered.

When you find a domain you want to register:

1. Enter `n` when prompted to stop searching.
2. Enter the number corresponding to the domain you want.
3. DomCat will display the URL where you can register the domain.
4. Open the URL in your browser and complete the registration process.


## To do list:
- [] Work on read me

- [] Commandize code

- [] Make options   
    - [] Option to use domains from file   
    - [] Option to write domain info to file   
        - [] options for both all domains and only the one that is selected
    - [] Option to just check a domain's categorization   
    - [] Option to check a list of domains categorization (piping and file input)  
    - [] Option to say who's accountID to use if multiple   
        - [] Logic to handle if there are multiple and no accountID was specified  
    - [] Option to say what categorization you are looking for
    - [] Option for how long to look?
    - [] Option for whoisxml cat check
        - [] Both for checking final domain picked and to replace cloudflare as the main
    - [] Option for whoisxml rep check 
    - [] Option for categorization we are looking for
