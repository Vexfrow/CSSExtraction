# CSSExtraction
Little tool to perform data extraction through CSS

It is done using the Sequential Import Chaining (SIC) technique from [d0nut](https://d0nut.medium.com/better-exfiltration-via-html-injection-31c72a2dae8b)

This tool was mainly created to learn Golang and the [Cobra CLI framework](https://cobra.dev/).


# How to use it

Build it from root directory
```
go build . -o CSSExtraction
```

Launch it
```
./CSSExtraction
```

## Possible Options

- **-p** : Port to be used for the server
- **-v** : Activate verbose mode. Print every data at each point of the technuique, instead of just the result


```
./CSSExtraction -v -p [Port]
```

# What it does 

Launch an HTTP server with a malicious CSS that can perform the SIC technique to extract data from another webpage (if it's vulnerable to the technique, see d0nutptr article)



# Credit

d0nut for the [technique](https://d0nut.medium.com/better-exfiltration-via-html-injection-31c72a2dae8b) and its [implementation in Rust](https://github.com/d0nutptr/sic)