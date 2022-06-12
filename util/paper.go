package util

import (
	"bytes"
	"fmt"
	"math"
	"strings"
)

type Paper struct {
	Page     int
	Totalnum int
	Pagesize int
	urlpath  string
	urlquery string
	nopath   bool
}

func NewPager(page, totalnum, pagesize int, url string, nopath ...bool) *Paper {
	p := new(Paper)
	p.Page = page
	p.Totalnum = totalnum
	p.Pagesize = pagesize

	arr := strings.Split(url, "?")
	p.urlpath = arr[0]
	if len(arr) > 1 {
		p.urlquery = "?" + arr[1]
	} else {
		p.urlquery = ""
	}
	if len(nopath) > 0 {
		p.nopath = nopath[0]
	} else {
		p.nopath = false
	}
	return p
}
func (c *Paper) url(page int) string {
	if c.nopath {
		if c.urlquery != "" {
			return fmt.Sprintf("%s%s&page=%d", c.urlpath, c.urlquery, page)
		} else {
			return fmt.Sprintf("%s?page=%d", c.urlpath, page)
		}
	} else {
		return fmt.Sprintf("%s/page/%d%s", c.urlpath, page, c.urlquery)
	}
}

func (c *Paper) ToString() string {
	if c.Totalnum <= c.Pagesize {
		return ""
	}
	var buf bytes.Buffer
	var from, to, linknum, offset, totalpage int
	offset = 5
	linknum = 10
	totalpage = int(math.Ceil(float64(c.Totalnum) / float64(c.Pagesize)))
	if totalpage < linknum {
		from = 1
		to = totalpage
	} else {
		from = c.Page - offset
		to = from + linknum
		if from < 1 {
			from = 1
			to = from + linknum - 1
		} else if to > totalpage {
			to = totalpage
			from = totalpage - linknum + 1
		}
	}

	if c.Page > 1 {
		buf.WriteString(fmt.Sprintf("<a class=\"layui-laypage-prev\" href=\"%s\">上一页</a></li>", c.url(c.Page-1)))
	} else {
		buf.WriteString("<span>上一页</span>")
	}
	if c.Page > linknum {
		buf.WriteString(fmt.Sprintf("<a href=\"%s\" class=\"laypage_first\">1...</a>", c.url(1)))
	}

	for i := from; i <= to; i++ {
		if i == c.Page {
			buf.WriteString(fmt.Sprintf("<span class=\"layui-laypage-curr\"><em class=\"layui-laypage-em\"></em><em>%d</em></span>", i))
		} else {
			buf.WriteString(fmt.Sprintf("<a href=\"%s\">%d</a>", c.url(i), i))
		}

	}

	if totalpage > to {
		buf.WriteString(fmt.Sprintf("<a class=\"layui-laypage-last\" href=\"%s\">末页</a>", c.url(totalpage)))
	}
	if c.Page < totalpage {
		buf.WriteString(fmt.Sprintf("<a class=\"layui-laypage-next\" href=\"%s\">下一页</a></li>", c.url(c.Page+1)))
	} else {
		buf.WriteString("<span>下一页</span>") //unnecessary use fmt.Sprintf
	}
	return buf.String()
}
