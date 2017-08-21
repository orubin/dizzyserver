package lib

func NewUserFetcher(ids []int) *UserFetcher {
	return &UserFetcher{ids: ids}
}

type UserFetcher struct {
	ids []int
}

func (u *UserFetcher) Fetch() {

}
