# DomCat - Domain Categorization
   
need an API key from NameSilo   
https://www.namesilo.com/account/api-manager   

and also one from cloudflare, but we will need a token here.   
make sure it has the permissions to read account settings and read from Intel.   
need to read account settings to get account ID for categorization call.   
https://dash.cloudflare.com/8c1b8ff70734a72331df6d7f2d6625e4/api-tokens   
https://developers.cloudflare.com/security-center/intel-apis/limits/

   
### To do list:
    commandize code   

    make options   
        option to use domains from file   
        option to write domains to file   
        option to just check a domain's categorization   
        option to check a list of domains categorization (piping and file input)  
        option to say who's accountID to use if multiple   
            logic to handle if there are multiple and no accountID was specified  

    make list of categorized domains with all the info   

    make able to wait for an interval rather than time.sleep   


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
