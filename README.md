## gitfame

The application calculates the following statistics for repository contributors:

```
$ gitfame --repository=. --extensions='.h,.cpp,.md' --order-by=lines

Name            Lines Commits Files
smirnovlad      2310  7       6
Krekhov Nikolai 2293  8       6
Stase           2008  9       5
tsmax2004       1877  9       6
Evgen           1612  6       6
aastrakhantsev  1500  4       4
NeKita          1246  15      5
mrnonam88       1167  8       4
vsvood          1051  10      5
Vladimir        1006  8       4
Dvirnyak        960   8       5
funtalex        944   8       5
MrMarvinColex   893   8       4
George          833   5       5
AnnaEf24        808   7       3
LenkaVelegurina 281   4       2
Supyness        263   7       4
Lena Bobyleva   182   1       2
Eugene Vihrev   2     1       1
```

- Number of lines
- Number of commits
- Number of files

All statistics are calculated for the state of the repository at the time of a specific commit.

## Flags

The utility support the following set of flags:

- `--repository` - Path to the Git repository; current directory by default.
- `--revision` - Pointer to the commit; `HEAD` by default.
- `--order-by` - Key for sorting the results; one of `lines` (default), `commits`, `files`.

By default, the results are sorted in descending order of the key (lines, commits, files). In case of equal keys, the author with the lexicographically smaller name will be higher. When using the flag, the corresponding field in the key is moved to the first place.

- `--use-committer` - Boolean flag, replacing the author (default) with the committer in calculations.
- `--format` - Output format; one of `tabular` (default), `csv`, `json`, `json-lines`.
- `--extensions` - List of extensions, narrowing the list of files in the calculation; a set of constraints is separated by commas, e.g., `.go,.md`.
- `--languages` - List of programming, markup, and other languages, narrowing the list of files in the calculation; a set of constraints is separated by commas, e.g., `go,markdown`.
- `--exclude` - Set of Glob patterns, excluding files from the calculation, e.g., `foo/*,bar/*`.
- `--restrict-to` - Set of Glob patterns, excluding all files that do not match any of the patterns in the set.

## Tests

Command to run tests:\
```go test -v ./test/integration/...```


## Building the Application

**_How to build the application?_**

```cd cmd/gitfame && go build .```

An executable file named `gitfame` will appear in `gitfame/cmd/gitfame`.

**_How to build the application and install it in `GOPATH/bin`?_**

```go install ./cmd/gitfame/...```

