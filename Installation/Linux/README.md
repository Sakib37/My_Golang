# Installation and Setup 

1. Run install_golang.sh and it will install golang in a linux amd64 system.
		./install_golang.sh

2. To enable bash completion for golang copy the 'go' file from 

		https://github.com/kura/go-bash-completion/blob/master/etc/bash_completion.d/go

   to "/etc/bash_completion.d/". 

3. Make sure bash_completion is already installed. Otherwise install(in Ubuntu/Debian) bash completion 


		apt-get install bash-completion

   and put the following in ~/.bashrc file
		# Bash completion
		if [ -f /etc/bash_completion ]; then
 		    . /etc/bash_completion
		fi






# Initial Tips:

1. Go always looks for package or .go file path, relative to the $GOPATH variable. For example when installing a package
   using "go install PACKAGE_NAME" ,  Go will look for the .go file in the "$GOPATH/src/PACKAGE_NAME" directory. 
   The package is installed in "$GOPATH/pkg/linux_amd64/" directory in a Linux x64 system.

2. Go package directory name and package .go file name should be same. Otherwise installing a package will fail. 


