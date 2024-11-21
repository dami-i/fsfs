# FSFS: Fast Static File Server

> [!WARNING]
> 
> This application runs on Linux only 🐧. Windows version is under development.

## Roadmap

- [ ] Implement server restart/cache clearing on file change.
- [ ] Embed a simple `index.html` as a fallback if not found in the user's directory.
- [ ] Windows version

## Installation

### Manual download

See available downloads on [releases page](https://github.com/dami-i/fsfs/releases).

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

Make sure you have Go installed (minimum version 1.23.3).

```sh
git clone https://github.com/dami-i/fsfs
cd fsfs
make
cd bin
sudo ln -s $(pwd)/fsfs /usr/local/bin
```

## Usage examples

### Serve current directory

```sh
fsfs
```

### Serve another directory

#### Relative path

```sh
fsfs ./static
```

#### Absolute path

```sh
fsfs /home/myuser/website/static
```

### Specify port

```sh
fsfs -p 3000 ./static
```

> [!NOTE]
> 
> The default port is **5000**. If busy, it tries to find the next free port.
>
> It won't try to find a free port if specified via `-p` flag.
