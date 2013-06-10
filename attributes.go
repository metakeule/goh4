package goh4

func Rel_(relation string) attrs { return Attr("rel", relation) }
func Href_(url string) attrs     { return Attr("href", url) }
func Type_(ty string) attrs      { return Attr("type", ty) }
func Media_(md string) attrs     { return Attr("media", md) }
func Src_(src string) attrs      { return Attr("src", src) }
func Name_(n string) attrs       { return Attr("name", n) }
func Value_(v string) attrs      { return Attr("value", v) }
