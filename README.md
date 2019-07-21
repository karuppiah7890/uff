# uff (update-fish-food)

Update [gofish](https://gofi.sh) food easily in the [fishworks/fish-food](https://github.com/fishworks/fish-food) repo

## Installation

This will install `uff` to `$GOBIN`

```
$ go get -v github.com/karuppiah7890/uff
```

## Usage

```
$ uff <food-file> <version>
```

## Example

```
$ git clone https://github.com/fishworks/fish-food
$ cd fish-food
$ uff Food/helm.lua 2.14.2
existing food version: 2.14.1
upgrading fish food Food/helm.lua to version 2.14.2 ...
done! 🐠
```
