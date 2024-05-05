# TxTracker

> This is a early version of the TxTracker, a tool for static check of Solidity.

## Platform

This project is fully supported on Linux and MacOS. Windows is not supported.

## Dependencies

- [solc-select](https://github.com/crytic/solc-select)

You should change the global version of Solidity to <YOUR_CONTRACT_VERSION> by running the following command:

```bash
solc-select install 0.5.17 # Assume you want to use Solidity 0.5.17
solc-select use 0.5.17
```

The TxTrancker will use the Solidity version you set and do NOT alter it automatically.

## For Developers

Please clone the repository and run the following commands:

```bash
zsh build.zsh
```
If suscessful, you will see the following message:

```bash
Cleaning up existing build directory...
Building the project...
Build successful! Executable is located at ./build/bin/txtracker
```
Then you can run the executable with the following command:

```bash
cd build/bin
./txtracker
```
Please place the Solidity files you want to analyze in the `dataset/contracts` directory.

If you want to handle a single file, you can use the following command:

```bash
./txtracker 0x0a0e40db3bc35ea2242d4475a67454078f83a9bf.sol
```

This will ONLY analyze the file `0x0a0e40db3bc35ea2242d4475a67454078f83a9bf.sol` placed in the `dataset/contracts` directory.

Otherwise, TxTracker will analyze all the files in the `dataset/contracts` directory.