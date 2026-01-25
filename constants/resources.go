package constants

type Resource struct {
	URL      string
	Name     string
	Category string
	Format   string
}

var (
	FitNews = Resource{
		URL:      "https://www.fit.hcmus.edu.vn/vn/feed.aspx",
		Name:     "Khoa CNTT - FIT@HCMUS",
		Category: "fithcmus",
		Format:   "RSS",
	}

	LichThi = Resource{
		URL:      "http://ktdbcl.hcmus.edu.vn/index.php/cong-tac-kh-o-thi/l-ch-thi-h-c-ky?format=feed&type=rss",
		Name:     "Lịch thi - PKTĐBCL@HCMUS",
		Category: "lichthi",
		Format:   "RSS",
	}

	ThongBao = Resource{
		URL:      "http://ktdbcl.hcmus.edu.vn/index.php/thong-bao?format=feed&type=rss",
		Name:     "Thông báo - PKTĐBCL@HCMUS",
		Category: "thongbao",
		Format:   "RSS",
	}

	HCMUSNews = Resource{
		URL:      "https://hcmus.edu.vn/category/dao-tao/dai-hoc/thong-tin-danh-cho-sinh-vien",
		Name:     "Thông tin dành cho sinh viên - HCMUS",
		Category: "hcmus",
		Format:   "HTML",
	}
)
