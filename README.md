sgcloud - gcloud project picker
===

## Installation

1. Install [gcloud](https://cloud.google.com/sdk/docs/install#linux), if you haven't already.
2. Download sgcloud from [release page](https://github.com/sato-s/sgcloud/releases)

### example

Linux

```
sudo wget https://github.com/sato-s/sgcloud/releases/latest/download/sgcloud-linux-amd64 -O /usr/local/bin/sgcloud
sudo chmod 755 /usr/local/bin/sgcloud
```

Mac (apple silicon)

```
sudo wget https://github.com/sato-s/sgcloud/releases/latest/download/sgcloud-darwin-arm64 -O /usr/local/bin/sgcloud
sudo chmod 755 /usr/local/bin/sgcloud
```

## Usage

Just run `sgcloud` to pick a gcloud project. Type character to fuzzy find a project. `C-n` and `C-p` to navigate.  
Hitting the enter key to choose a project to use with gcloud command. This is equivalent to `gcloud config set project [Project of your choice]`.

Also, you can use `sgcloud -b` to open google cloud console in your browser.
