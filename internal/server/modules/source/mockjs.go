package source

import (
	"github.com/aaronchen2k/deeptest/internal/pkg/consts"
	"github.com/aaronchen2k/deeptest/internal/server/modules/model"
	"github.com/aaronchen2k/deeptest/internal/server/modules/repo"
	"github.com/getkin/kin-openapi/openapi3"
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
			Type:       openapi3.TypeInteger,
		},
		{
			Expression: "telephone()",
			Name:       "固话号码",
			Type:       openapi3.TypeString,
		},
		{
			Expression: "areacode()",
			Name:       "电话区号",
			Type:       openapi3.TypeString,
		},
		{
			Expression: "natural(1,100)",
			Name:       "自然数",
			Type:       openapi3.TypeInteger,
		},
		{
			Expression: "integer(1,100)",
			Name:       "整数",
			Type:       openapi3.TypeInteger,
		},
		{
			Expression: "float(1, 10, 2, 5)",
			Name:       "小数",
			Type:       openapi3.TypeNumber,
		},
		{
			Expression: "character(pool)",
			Name:       "单个字符",
			Type:       openapi3.TypeString,
		},
		{
			Expression: "string(pool, 1, 10)",
			Name:       "字符串",
			Type:       openapi3.TypeString,
		},
		{
			Expression: "range(1, 100, 1)",
			Name:       "整数数组",
			Type:       openapi3.TypeArray,
		},
		{
			Expression: "date('yyyy-MM-dd')",
			Name:       "日期",
			Type:       openapi3.TypeString,
		},
		{
			Expression: "time('HH:mm:ss')",
			Format:     "time",
			Name:       "时间",
			Type:       openapi3.TypeString,
		},
		{
			Expression: "datetime('yyyy-MM-dd HH:mm:ss')",
			Format:     "data-time",
			Name:       "日期时间",
			Type:       openapi3.TypeString,
		},
		{
			Expression: "now('yyyy-MM-dd HH:mm:ss')",
			Name:       "当前日期时间",
			Type:       openapi3.TypeString,
		},
		{
			Expression: "guid()",
			Name:       "唯一标识符",
			Format:     "uuid",
			Type:       openapi3.TypeString,
		},
		{
			Expression: "increment(1)",
			Name:       "自增整数",
			Type:       openapi3.TypeInteger,
		},
		{
			Expression: "url('http')",
			Name:       "URL地址",
			Type:       openapi3.TypeString,
		},
		{
			Expression: "uri()",
			Name:       "URI地址",
			Format:     "uri",
			Type:       openapi3.TypeString,
		},
		{
			Expression: "ruri()",
			Name:       "相对URI",
			Format:     "uri-reference",
			Type:       openapi3.TypeString,
		},
		{
			Expression: "uriTempl()",
			Name:       "URI模板",
			Format:     "uri-template",
			Type:       openapi3.TypeString,
		},

		{
			Expression: "",
			Name:       "国际化资源标识符",
			Format:     "iri",
			Type:       openapi3.TypeString,
		},
		{
			Expression: "",
			Name:       "相对国际化资源标识符",
			Format:     "iri-reference",
			Type:       openapi3.TypeString,
		},

		{
			Expression: "protocol()",
			Name:       "URL协议",
			Type:       openapi3.TypeString,
		},
		{
			Expression: "host()",
			Name:       "主机名",
			Format:     "hostname",
			Type:       openapi3.TypeString,
		},
		{
			Expression: "idnHost()",
			Name:       "国际化主机名",
			Format:     "idn-hostname",
			Type:       openapi3.TypeString,
		},
		{
			Expression: "domain()",
			Name:       "网站域名",
			Type:       openapi3.TypeString,
		},
		{
			Expression: "tld()",
			Name:       "顶级域名",
			Type:       openapi3.TypeString,
		},
		{
			Expression: "email()",
			Name:       "邮件地址",
			Format:     "email",
			Type:       openapi3.TypeString,
		},
		{
			Expression: "idnEmail",
			Name:       "国际化域名邮件地址",
			Format:     "idn-email",
			Type:       openapi3.TypeString,
		},

		{
			Expression: "ip()",
			Name:       "IP地址",
			Format:     "ipv4",
			Type:       openapi3.TypeString,
		},
		{
			Expression: "ipv6()",
			Name:       "IPV6地址",
			Format:     "ipv6",
			Type:       openapi3.TypeString,
		},
		{
			Expression: "region()",
			Name:       "大区",
			Type:       openapi3.TypeString,
		},
		{
			Expression: "province()",
			Name:       "省份",
			Type:       openapi3.TypeString,
		},
		{
			Expression: "city()",
			Name:       "城市",
			Type:       openapi3.TypeString,
		},
		{
			Expression: "county()",
			Name:       "县区",
			Type:       openapi3.TypeString,
		},
		{
			Expression: "county(true)",
			Name:       "省市县",
			Type:       openapi3.TypeString,
		},
		{
			Expression: "zip()",
			Name:       "邮编",
			Type:       openapi3.TypeString,
		},
		{
			Expression: "first()",
			Name:       "英文名",
			Type:       openapi3.TypeString,
		},
		{
			Expression: "last()",
			Name:       "英文姓",
			Type:       openapi3.TypeString,
		},
		{
			Expression: "name()",
			Name:       "英文姓名",
			Type:       openapi3.TypeString,
		},
		{
			Expression: "cfirst()",
			Name:       "中文名",
			Type:       openapi3.TypeString,
		},
		{
			Expression: "clast()",
			Name:       "中文姓",
			Type:       openapi3.TypeString,
		},
		{
			Expression: "cname()",
			Name:       "中文姓名",
			Type:       openapi3.TypeString,
		},
		{
			Expression: "password()",
			Name:       "密码",
			Format:     "password",
			Type:       openapi3.TypeString,
		},
		{
			Expression: "byte()",
			Name:       "字节",
			Format:     "byte",
			Type:       openapi3.TypeString,
		},
		{
			Expression: "",
			Name:       "JSON指针",
			Format:     "json-pointer",
			Type:       openapi3.TypeString,
		},
		{
			Expression: "",
			Name:       "正则表达式",
			Format:     "regex",
			Type:       openapi3.TypeString,
		},

		{
			Expression: "color()",
			Name:       "十六进制颜色",
			Type:       openapi3.TypeString,
		},
		{
			Expression: "rgb()",
			Name:       "RGB颜色",
			Type:       openapi3.TypeString,
		},
		{
			Expression: "rgba()",
			Name:       "RGBA颜色",
			Type:       openapi3.TypeString,
		},
		{
			Expression: "hsl()",
			Name:       "HSL颜色",
			Type:       openapi3.TypeString,
		},
		{
			Expression: "paragraph()",
			Name:       "英文文本",
			Type:       openapi3.TypeString,
		},
		{
			Expression: "cparagraph()",
			Name:       "中文文本",
			Type:       openapi3.TypeString,
		},
		{
			Expression: "sentence()",
			Name:       "英文句子",
			Type:       openapi3.TypeString,
		},
		{
			Expression: "csentence()",
			Name:       "中文句子",
			Type:       openapi3.TypeString,
		},
		{
			Expression: "word()",
			Name:       "英文单词",
			Type:       openapi3.TypeString,
		},
		{
			Expression: "cword()",
			Name:       "中文汉字",
			Type:       openapi3.TypeString,
		},
		{
			Expression: "title()",
			Name:       "英文标题",
			Type:       openapi3.TypeString,
		},
		{
			Expression: "ctitle()",
			Name:       "中文标题",
			Type:       openapi3.TypeString,
		},
	}
	return
}

func (s *MockJsExpressionSource) Init(tenantId consts.TenantId) error {
	if s.MockJsExpressionRepo.DB.Model(&model.MockJsExpression{}).
		//Where("id IN ?", []int{1}).
		Find(&[]model.MockJsExpression{}).RowsAffected > 0 {
		color.Danger.Printf("\n[Mysql] --> %s 表的初始数据已存在!", model.MockJsExpression{}.TableName())
		//return nil
	}

	sources, err := s.GetSources()
	if err != nil {
		return err
	}

	if _, _, err := s.MockJsExpressionRepo.BatchCreateExpression(tenantId, sources); err != nil {
		return err
	}

	color.Info.Printf("\n[Mysql] --> %s 表初始数据成功!", model.MockJsExpression{}.TableName())
	return nil
}
