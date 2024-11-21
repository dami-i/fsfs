# FSFS: Fast Static File Server

Quickly spin up a local static file web server.

> [!WARNING]
> 
> 🐧 This application runs on Linux only.
> 
> 🪟 A Windows version is under development. 🚧

## Roadmap

- [ ] Implement server restart/cache clearing on file changes (the "watch" feature).
- [ ] Embed a simple `index.html` as a fallback if none is found in the user's directory.
- [ ] Develop a Windows version.

### Contribute

I am happy to evaluate and review contributions submitted via pull requests.

## Installation

### Manual download

Check out the available downloads on the [releases page](https://github.com/dami-i/fsfs/releases).

### Via `go install`

```sh
go install github.com/dami-i/fsfs@latest
```

<!--
### Via bash installer

```sh
wget https:// | bash
```
-->

### Building from source

Make sure you have Go (minimum version 1.23.3) and Make installed.

```sh
git clone https://github.com/dami-i/fsfs
cd fsfs
make
sudo ln -s $(pwd)/bin/fsfs /usr/local/bin
```

## Usage examples

### Serve the current directory

```sh
fsfs
```

> [!IMPORTANT]
>
> The **watch** feature is not yet implemented.
> 
> This means you will need to press <kbd>Ctrl</kbd> + <kbd>F5</kbd> (hard refresh / bypass cache) in your browser to see the latest version of your file.

### Serve another directory

- Relative path

```sh
fsfs ./static
```

- Absolute path

```sh
fsfs /home/myuser/website/static
```

### Specify a port

```sh
fsfs -p 3000 ./static
```

> [!NOTE]
> 
> The default port is **5000**.
> 
> If busy, the app will attempt to find the next available port.
>
> However, if you specify a port using the `-p` flag, it will not search for an alternative.
