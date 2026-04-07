package config

type Resource struct {
	URL      string
	Name     string
	Category string
	Format   string
}

/*
Rule: If you add new feeds, make sure they match table names in db/database.sql.
Because i configure A TABLE NAME IN DB SCHEMA = A CATEGORY IN FEEDS
Check db/database.sql for more information.
*/

var Feeds = []Resource{
	{
		URL:      "https://www.fit.hcmus.edu.vn/vn/feed.aspx",
		Name:     "Khoa CNTT - FIT@HCMUS",
		Category: "fithcmus",
		Format:   "RSS",
	},
	{
		URL:      "http://ktdbcl.hcmus.edu.vn/index.php/cong-tac-kh-o-thi/l-ch-thi-h-c-ky?format=feed&type=rss",
		Name:     "Lịch thi - PKTĐBCL@HCMUS",
		Category: "lichthipkt",
		Format:   "RSS",
	},
	{
		URL:      "http://ktdbcl.hcmus.edu.vn/index.php/thong-bao?format=feed&type=rss",
		Name:     "Thông báo - PKTĐBCL@HCMUS",
		Category: "thongbaopkt",
		Format:   "RSS",
	},
	{
		URL:      "https://hcmus.edu.vn/category/dao-tao/dai-hoc/thong-tin-danh-cho-sinh-vien",
		Name:     "Thông tin dành cho sinh viên - HCMUS",
		Category: "hcmus",
		Format:   "HTML",
	},
}
