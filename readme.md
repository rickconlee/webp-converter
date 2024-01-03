# JPG/JPEG PNG GIF TIF/TIFF to WebP Converter 

This project arose out of a need to convert images of various types to WebP. This is done for technical SEO. 

This tool can be run locally from the command line, on your server, or in a docker container. 

The source code is written in `Go`, and compiled on `1.21.5`

## Usage Instructions 

### Standalone Executable
Download the latest release for your OS/Arch from `Releases` on the sidebar

Unzip the binary and place it in your home directory, programs directory or wherever you want to launch the program from 

#### Ubuntu Linux

Extract the executable files wherever and run:

``` shell
$ ./webpconverter source_folder/ destination_folder/
```


#### MacOSX (Intel and M1, M2, M3)

Extract the Extract the executable files wherever and run:

``` shell
$ ./webpconverter source_folder/ destination_folder/
```

#### Windows 10/11

This is a #todo item. Feel free to contribute. 

## Compile Instructions 

To compile this software on your system, make sure that you are running `Go 1.21.5` or later. 

``` shell
$ git clone git@github.com:rickconlee/webp-converter.git
```
Make sure you are on the `master` branch, then run the following: 

``` shell
$ go build -o webpconvert main.go
```

Once the compile is complete, run: 

``` shell
$ ./webpconverter source_folder/ destination_folder/
```