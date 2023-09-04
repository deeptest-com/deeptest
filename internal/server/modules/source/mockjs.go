package source

import (
	"github.com/aaronchen2k/deeptest/internal/server/modules/model"
	"github.com/aaronchen2k/deeptest/internal/server/modules/repo"
	"github.com/gookit/color"
)

type MockJsExpressionSource struct {
	MockJsExpressionRepo *repo.MockJsRepo `inject:""`
}

func (s *MockJsExpressionSource) GetSources() (ret []model.MockJsExpression, err error) {
	ret = []model.MockJsExpression{
		{
			Expression: "mobiphone()",
			Name:       "手机号码",
		},
		{
			Expression: "telephone()",
			Name:       "固话号码",
		},
		{
			Expression: "areacode()",
			Name:       "电话区号",
		},
		{
			Expression: "natural(1,100)",
			Name:       "自然数",
		},
		{
			Expression: "integer(1,100)",
			Name:       "整数",
		},
		{
			Expression: "float(1, 10, 2, 5)",
			Name:       "小数",
		},
		{
			Expression: "character(pool)",
			Name:       "单个字符",
		},
		{
			Expression: "string(pool, 1, 10)",
			Name:       "字符串",
		},
		{
			Expression: "range(1, 100, 1)",
			Name:       "整数数组",
		},
		{
			Expression: "date('yyyy-MM-dd')",
			Name:       "日期",
		},
		{
			Expression: "time('HH:mm:ss')",
			Name:       "时间",
		},
		{
			Expression: "datetime('yyyy-MM-dd HH:mm:ss')",
			Name:       "日期时间",
		},
		{
			Expression: "now('yyyy-MM-dd HH:mm:ss')",
			Name:       "当前日期时间",
		},
		{
			Expression: "guid()",
			Name:       "唯一标识符",
		},
		{
			Expression: "increment(1)",
			Name:       "自增整数",
		},
		{
			Expression: "url('http')",
			Name:       "URL地址",
		},
		{
			Expression: "protocol()",
			Name:       "URL协议",
		},
		{
			Expression: "domain()",
			Name:       "网站域名",
		},
		{
			Expression: "tld()",
			Name:       "顶级域名",
		},
		{
			Expression: "email()",
			Name:       "邮件地址",
		},
		{
			Expression: "ip()",
			Name:       "IP地址",
		},
		{
			Expression: "region()",
			Name:       "大区",
		},
		{
			Expression: "province()",
			Name:       "省份",
		},
		{
			Expression: "city()",
			Name:       "城市",
		},
		{
			Expression: "county()",
			Name:       "县区",
		},
		{
			Expression: "county(true)",
			Name:       "省市县",
		},
		{
			Expression: "zip()",
			Name:       "邮编",
		},
		{
			Expression: "first()",
			Name:       "英文名",
		},
		{
			Expression: "last()",
			Name:       "英文姓",
		},
		{
			Expression: "name()",
			Name:       "英文姓名",
		},
		{
			Expression: "cfirst()",
			Name:       "中文名",
		},
		{
			Expression: "clast()",
			Name:       "中文姓",
		},
		{
			Expression: "cname()",
			Name:       "中文姓名",
		},
		{
			Expression: "color()",
			Name:       "十六进制颜色",
		},
		{
			Expression: "rgb()",
			Name:       "RGB颜色",
		},
		{
			Expression: "rgba()",
			Name:       "RGBA颜色",
		},
		{
			Expression: "hsl()",
			Name:       "HSL颜色",
		},
		{
			Expression: "paragraph()",
			Name:       "英文文本",
		},
		{
			Expression: "cparagraph()",
			Name:       "中文文本",
		},
		{
			Expression: "sentence()",
			Name:       "英文句子",
		},
		{
			Expression: "csentence()",
			Name:       "中文句子",
		},
		{
			Expression: "word()",
			Name:       "英文单词",
		},
		{
			Expression: "cword()",
			Name:       "中文汉字",
		},
		{
			Expression: "title()",
			Name:       "英文标题",
		},
		{
			Expression: "ctitle()",
			Name:       "中文标题",
		},
	}
	return
}

func (s *MockJsExpressionSource) Init() error {
	if s.MockJsExpressionRepo.DB.Model(&model.MockJsExpression{}).
		//Where("id IN ?", []int{1}).
		Find(&[]model.MockJsExpression{}).RowsAffected > 0 {
		color.Danger.Printf("\n[Mysql] --> %s 表的初始数据已存在!", model.MockJsExpression{}.TableName())
		return nil
	}

	sources, err := s.GetSources()
	if err != nil {
		return err
	}

	if _, _, err := s.MockJsExpressionRepo.BatchCreateExpression(sources); err != nil {
		return err
	}

	color.Info.Printf("\n[Mysql] --> %s 表初始数据成功!", model.MockJsExpression{}.TableName())
	return nil
}
