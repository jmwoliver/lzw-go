# LZW Implementation in Go

This is a small implementation of the LZW compression algorithm in Go. It implements both the compression and decompression sides and can be used with the predefined flags outlined in the Usage section.

Note: this is a small version meant to learn how LZW worked. It doesn't work for long text files because when the dictionary exceeds 256, it can't convert the bytes (2^8 = 256) properly. Full implementations account for this and provide more bits after the overflow.

## Installation

Pull the repo down from the [github](https://github.com/jmwoliver/lzw-go) page.

```bash
git clone github.com/jmwoliver/lzw-go
```

Then navigate to the `lzw-go` directory and build the file.

```bash
cd lzw-go
go build main.go
```

Now an executable has been built and can be run.

```bash
./main <...flags...>
```


## Usage

After building the project, you can run the following command (*note: this currently only works on .txt files*):

```bash
./main -file /path/to/text/file.txt -compress
```

Once that command has been run, it will generate a `compressed.lzw` file in the `output` directory. You can see the compressed contents of this file with:

```bash
cat ./output/compressed.lzw
```

After compressing, you can then decompress the `compressed.lzw` file (*note: this command only works on .lzw files*):

```bash
./main -file ./output/compressed.lzw -decompress
```

Once you have done this, you can see the newly generated, decompressed file in `decompressed.txt` in the `output` directory:

```bash
cat ./output/decompressed.txt
```

Bada bing bada boom! Well isn't that neat, it's the same as the original!

It's also fun to see the size changes by looking at the byte values on the original file and the files generated.

```bash
ls -l /path/to/text/file.txt
ls -l ./output/compressed.lzw
ls -l ./output/decompressed.txt
```

That's it, hope you enjoy!

## License
[MIT](https://choosealicense.com/licenses/mit/)