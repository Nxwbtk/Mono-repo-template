package catschemas

type TGetCat struct {
	ID   string `json:"id"`
	Name string `json:"name"`
}

type TPostCat struct {
	Name string `json:"name"`
}
