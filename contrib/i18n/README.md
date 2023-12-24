# internationalization
> Output the corresponding results according to different languages [go-i18n](http://github.com/nicksnyder/go-i18n)

## Usage
```go
package main

import (
	"fmt"

	tran "github.com/xiecang/jarvis/contrib/i18n"
	"github.com/xiecang/jarvis/i18n"
)

func main() {
	translator, err := tran.NewTranslator("/BundleDir")
	if err != nil {
		panic(err)
	}
	fmt.Println(translator.Tr(i18n.LanguageChinese, "base.success"))
	fmt.Println(translator.Tr(i18n.LanguageEnglish, "base.success"))
}

```
