# Jupyter Notebook Handler for Hugo

Jupyter Handler for [Hugo](https://gohugo.io/).

**Deprecated: This product does not work for latest version of Hugo.**

## Installation

```bash
go get -u github.com/naoina/hugo-jupyter-handler
```

## Usage

Set `frontmatter` metadata of your `.ipynb` file from *Edit* -> *Edit Notebook Metadata* like below.

```json
{
  "frontmatter": {
    "title": "Title Of Content",
    "date": "2017-07-26T11:51:30+09:00",
    "tags": [
      "Jupyter Notebook",
      "Hugo",
      "Go"
    ]
  },
  "kernelspec": {
    "name": "python3",
    "display_name": "Python 3",
    "language": "python"
  },
  "language_info": {
    "name": "python",
    "version": "3.6.1",
    "mimetype": "text/x-python",
    "codemirror_mode": {
      "name": "ipython",
      "version": 3
    },
    "pygments_lexer": "ipython3",
    "nbconvert_exporter": "python",
    "file_extension": ".py"
  }
}
```

`frontmatter` metadata will be used as Hugo's [Frontmatter](https://gohugo.io/content-management/front-matter/) as is.

Place your `.ipynb` file in a contents directory of Hugo site (e.g. `content/`).

Place the following code in a root directory of your Hugo site as `hugo.go`.

```go
package main

import (
	"os"
	"runtime"

	"github.com/gohugoio/hugo/commands"
	jww "github.com/spf13/jwalterweatherman"

	_ "github.com/naoina/hugo-jupyter-handler"
)

func main() {
	runtime.GOMAXPROCS(runtime.NumCPU())
	commands.Execute()
	if jww.LogCountForLevelsGreaterThanorEqualTo(jww.LevelError) > 0 {
		os.Exit(-1)
	}

	if commands.Hugo != nil {
		if commands.Hugo.Log.LogCountForLevelsGreaterThanorEqualTo(jww.LevelError) > 0 {
			os.Exit(-1)
		}
	}
}
```

And use `go run hugo.go` instead of `hugo` CLI.

```bash
go run hugo.go server -w
```

Also you can build and use your own `hugo` CLI including Jupyter Notebook Handler.

```bash
go build -o hugo
./hugo server -w
```

## License

MIT
