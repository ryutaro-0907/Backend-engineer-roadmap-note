# Os and Computer Science

## How computer start?

### Sequence
1. CPU starts and send instructions to RAM from BIOS, which is stored in ROM.
2. The BIOS starts the monitor and keyboard, and does some basic checks to make sure the computer is working properly. For example, it will look for the RAM.
3. Then, the BIOS start the boot sequence. It will look for the operating system.
4. If there is no change in setting, The BIOS will fetch operating system from hard drive and load it into the RAM.
5. Then, the BIOS transfers control to the operating system.

## Bash
Bash is a Unix shell and command language written by Brian Fox for the GNU Project as a free software replacement for the Bourne shell.[10][11] First released in 1989,[12] it has been used as the default login shell for most Linux distributions.[13]
### Set path
``` 
export PATH=$PATH:/place/with/the/file
```
### Personalize terminal
install oh-my-zsh in your terminal
first install zsh with homebrew 
```
brew install zsh
```
check available shells
```
sudo vi /etc/shells
Password:
```
add this to the file and save
```
/usr/local/bin/zsh
```

check if zsh is being used
```
TEST-USER% echo $SHELL
/usr/local/bin/zsh
``` 

install oh-my-zsh
```
$ sh -c "$(curl -fsSL https://raw.githubusercontent.com/robbyrussell/oh-my-zsh/master/tools/install.sh)"
```

configure zshrc 
```
$ nano ~/.zshrc
```
Let's change theme 
```
ZSH_THEME="candy"
```
Save the file and type
```
$ source ~/.zshrc
```
to reflect change 

Configure syntax highlight
```
$ git clone https://github.com/zsh-users/zsh-syntax-highlighting.git ${ZSH_CUSTOM:-~/.oh-my-zsh/custom}/plugins/zsh-syntax-highlighting
``` 
```
$ nano ~/.zshrc
```
Again open the file and add 
```
plugins=(git zsh-syntax-highlighting)
```
```
$ source ~/.zshrc
```
to reflect change 

Install zsh-completions
```
$ git clone https://github.com/zsh-users/zsh-completions ~/.oh-my-zsh/custom/plugins/zsh-completions
```
```
$ nano ~/.zshrc
```
add this line 
```
 plugins=(git zsh-syntax-highlighting zsh-completions)

# config for zsh-completion
autoload -U compinit && compinit -u
``` 
```
$ source ~/.zshrc
```
to reflect change 
## SSH Login
### What is SSH? 
Secure Shell, sometimes referred to as Secure Socket Shell, is a protocol which allows you to connect securely to a remote computer or a server by using a text-based interface.

When a secure SSH connection is established, a shell session will be started, and you will be able to manipulate the server by typing commands within the client on your local computer.

System and network administrators use this protocol the most, as well as anyone who needs to manage a computer remotely in a highly secure manne

### How does SHH work? 
In order to establish an SSH connection, you need two components: a client and the corresponding server-side component. An SSH client is an application you install on the computer which you will use to connect to another computer or a server. The client uses the provided remote host information to initiate the connection and if the credentials are verified, establishes the encrypted connection.

On the server’s side, there is a component called an SSH daemon that is constantly listening to a specific TCP/IP port for possible client connection requests. Once a client initiates a connection, the SSH daemon will respond with the software and the protocol versions it supports and the two will exchange their identification data. If the provided credentials are correct, SSH creates a new session for the appropriate environment.

The default SSH protocol version for SSH server and SSH client communication is version 2

### How to enable SSH conncection?
Since creating an SSH connection requires both a client and a server component, you need to make sure they are installed on the local and the remote machine, respectively. An open source SSH tool—widely used for Linux distributions— is OpenSSH. Installing OpenSSH is relatively easy. It requires access to the terminal on the server and the computer that you use for connecting. Note that Ubuntu does not have SSH server installed by default.

### Code example with ssh key 
```
# Create ssh key on your computer.
ssh-keygen -t rsa -b 4096

# Send public key to remote server.
ssh-copy-id -i ~/.ssh/id_rsa.pub [remote_user]@[remote_server_hostname]

# Connect to remote with ssh 
ssh -i id_rsa [remote_user]@[remote_server_hostname]
```
You can specify key name by
```
ssh-keygen -t rsa -b 4096 {-f /home/[user_name]/.ssh/[key_name]}
```
If you don't specify key name, it will create keys in  /home/[username]/.ssh. 
private key：id_rsa
pulic key：id_rsa.pub

You can also set config file 
```
/home/[username]/.ssh/config
Host [ec2]
  HostName remote_host_name
  User remote_user
  IdentityFile /home/[user_name]/.ssh/[public_key]

Host [ec3]
  ...
```
To connect to ec2 
```
ssh ec2
```
## File system
### What is file system? 
A  file system  defines how files are  named,  stored, and  retrieved  from a storage device.

Every time you open a file on your computer or smart device, your operating system uses its file system internally to load it from the storage device.

Or when you copy, edit, or delete a file, the file system handles it under the hood.

Whenever you download a file or access a web page over the Internet, a file system is involved too.

It organize files into directory so it is easy to acccess and organize.

It also does, Space management, metadata, data encryption, file access control, and data integrity are the responsibilities of the file system too.

## Reference
1. https://www.futurelearn.com/info/courses/computer-systems/0/steps/53497
2. https://phoenixnap.com/kb/ssh-to-connect-to-remote-server-linux-or-windows
3. https://www.freecodecamp.org/news/file-systems-architecture-explained/

