# DomCat - Domain Categorization
   
Need an API key from NameSilo   
https://www.namesilo.com/account/api-manager   

And also one from cloudflare, but we will need a token here.   
Make sure it has the permissions to:
- Read account settings
- Read from Intel.   

Need to read account settings to get account ID for categorization call.   
https://dash.cloudflare.com/8c1b8ff70734a72331df6d7f2d6625e4/api-tokens   
https://developers.cloudflare.com/security-center/intel-apis/limits/

## How To
Make sure you are in the DomCat Directory.
Build the program   
```bash
go build
```
Run the program.   
```bash
./domCat
```
When you find a domain you like, say no to continue.   
Input the number coresponding to the domain you like.   
The url where you can find that domain will be displayed.   
Follow the URL and register your domain!   
   
### To do list:
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


### proposed code for install script
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
