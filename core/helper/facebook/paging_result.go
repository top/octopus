package facebook

import (
	"bytes"
	"fmt"
	"net/http"
)

type PagingResult struct {
	session  *Session
	paging   pagingData
	previous string
	next     string
}

type pagingData struct {
	Data      []Result `facebook:",required"`
	Paging    *pagingNavigator
	UsageInfo *UsageInfo
}

type pagingNavigator struct {
	Previous string
	Next     string
}

func newPagingResult(session *Session, res Result) (*PagingResult, error) {
	if _, ok := res["data"]; !ok {
		return nil, fmt.Errorf("facebook: current Result is not a paging response")
	}

	pr := &PagingResult{
		session: session,
	}
	paging := &pr.paging
	err := res.Decode(paging)

	if err != nil {
		return nil, err
	}

	if paging.Paging != nil {
		pr.previous = paging.Paging.Previous
		pr.next = paging.Paging.Next
	}

	return pr, nil
}

func (pr *PagingResult) Data() []Result {
	return pr.paging.Data
}

func (pr *PagingResult) UsageInfo() *UsageInfo {
	return pr.paging.UsageInfo
}

func (pr *PagingResult) Decode(v interface{}) (err error) {
	res := Result{
		"data": pr.Data(),
	}
	return res.Decode(v)
}

func (pr *PagingResult) Previous() (noMore bool, err error) {
	if !pr.HasPrevious() {
		noMore = true
		return
	}

	return pr.navigate(&pr.previous)
}

func (pr *PagingResult) Next() (noMore bool, err error) {
	if !pr.HasNext() {
		noMore = true
		return
	}

	return pr.navigate(&pr.next)
}

func (pr *PagingResult) HasPrevious() bool {
	return pr.previous != ""
}

func (pr *PagingResult) HasNext() bool {
	return pr.next != ""
}

func (pr *PagingResult) navigate(url *string) (noMore bool, err error) {
	var (
		pagingURL string
		params    = Params{}
	)
	pr.session.prepareParams(params)

	if len(params) == 0 {
		pagingURL = *url
	} else {
		buf := &bytes.Buffer{}
		buf.WriteString(*url)
		buf.WriteRune('&')
		params.Encode(buf)

		pagingURL = buf.String()
	}

	var request *http.Request
	var res Result

	request, err = http.NewRequest("GET", pagingURL, nil)

	if err != nil {
		return
	}

	res, err = pr.session.Request(request)

	if err != nil {
		return
	}

	if pr.paging.Paging != nil {
		pr.paging.Paging.Next = ""
		pr.paging.Paging.Previous = ""
	}

	paging := &pr.paging
	err = res.Decode(paging)

	if err != nil {
		return
	}

	paging.UsageInfo = res.UsageInfo()

	if paging.Paging == nil || len(paging.Data) == 0 {
		*url = ""
		noMore = true
	} else {
		pr.previous = paging.Paging.Previous
		pr.next = paging.Paging.Next
	}

	return
}
