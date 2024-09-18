package query

type GetTinyURLQuery struct {
	ID string `path:"id" required:"true"`
}

type GetInfoTinyURLQuery struct {
	ID string `path:"id" required:"true"`
}
