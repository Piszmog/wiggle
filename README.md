# wiggle

Wiggles your mouse between two points in a specified interval.

## Installation

### Homebrew

Run the following commands,

```shell
brew tap piszmog/tools
brew install piszmog/tools/wiggle
```

### Manual

Download `wiggle` for your system by heading over to [Releases](https://github.com/Piszmog/wiggle/releases) and download the artifact for your architecture.

Or you can use [gh](https://cli.github.com/) to download the artifact.

```shell
# Download the latest 64-bit version for Linux
gh release download -R Piszmog/wiggle -p '*Linux_x86_64*'
# Download the latest 64-bit Intel version for macOS
gh release download -R Piszmog/wiggle -p '*Darwin_x86_64*'
# Download the latest Silicon for macOS
gh release download -R Piszmog/wiggle -p '*Darwin_arm64*'

# Untar the artifact
tar -xf wiggle_0.1.0_Linux_x86_64.tar.gz
# Delete the artifact
rm wiggle_0.1.0_Linux_x86_64.tar.gz   
# Move the binary to a directory on your PATH
mv wiggle /some/directory/that/is/in/your/path
```

## Usage

```shell
$ ./wiggle
```

### Options

| Option                 | Default | Required  | Description                                                                                                                  |
|:-----------------------|:-------:|:---------:|:-----------------------------------------------------------------------------------------------------------------------------|
| `--pause`, `-p`        | `1 minutes`  | **False** | Time interval between mouse movements.                                                                |
| `--help`, `-h`         | `false`      | **False** | Shows help                                                                                                                   |
