# gim
easemob chat client for golang
-----------

```
import (
    "github.com/sanxia/gim"
)

imClient := gim.NewEmChatClient("easemob-playground",
"test1", "YXA6wDs-MACqEeSO0VcBzaqg9F", "YXA8JMWWlLap_YbI_ucz77j-4-mI9aK")

tokenResponse, _ := imClient.GetAccessToken()
```

tokenResponse
-----------
&{YWMtpyCGhCxYEeiv_kdJG3FSQwAAAAAAAAAAAAAAAAAAAAHAOz4wBGoR5I7RVwHNqqDkAgMAAAFiRCl2WwBPGgA5Z6iiKDtQCTmEwPssDfI_wG3RGLRI2-rcBPbuc6KY6Q 5184000 c03b3e30-046a-11e4-8ed1-5701cdaaa0e4}

