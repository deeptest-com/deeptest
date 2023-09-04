package test

import (
	"github.com/antchfx/htmlquery"
	"log"
	"strings"
	"testing"
)

func TestXPath(t *testing.T) {
	doc, err := htmlquery.Parse(strings.NewReader(docHtml))
	if err != nil {
		return
	}

	expression := `//*[@id="form"]/input/..`

	elems, err := htmlquery.Query(doc, expression)
	log.Print(elems)

}

const docHtml = `
<!DOCTYPE html>
<!--STATUS OK-->
<html>

<head>
    <meta http-equiv="Content-Type" content="text/html;charset=utf-8">
    <meta http-equiv="X-UA-Compatible" content="IE=edge,chrome=1">
    <meta content="always" name="referrer">
    <meta name="theme-color" content="#ffffff">
    <meta name="description" content="全球领先的中文搜索引擎、致力于让网民更便捷地获取信息，找到所求。百度超过千亿的中文网页数据库，可以瞬间找到相关的搜索结果。">
    <link rel="shortcut icon" href="/favicon.ico" type="image/x-icon" />
    <link rel="search" type="application/opensearchdescription+xml" href="/content-search.xml" title="百度搜索" />
    <link rel="icon" sizes="any" mask href="//www.baidu.com/img/baidu_85beaf5496f291521eb75ba38eacbd87.svg">
    <link rel="dns-prefetch" href="//dss0.bdstatic.com" />
    <link rel="dns-prefetch" href="//dss1.bdstatic.com" />
    <link rel="dns-prefetch" href="//ss1.bdstatic.com" />
    <link rel="dns-prefetch" href="//sp0.baidu.com" />
    <link rel="dns-prefetch" href="//sp1.baidu.com" />
    <link rel="dns-prefetch" href="//sp2.baidu.com" />
    <link rel="apple-touch-icon-precomposed"
        href="https://psstatic.cdn.bcebos.com/video/wiseindex/aa6eef91f8b5b1a33b454c401_1660835115000.png">
    <title>百度一下，你就知道</title>

    <style index="newi" type="text/css">
  
    </style>

    <script data-compress=strip>
        function h(obj){
            obj.style.behavior='url(#default#homepage)';
            var a = obj.setHomePage('//www.baidu.com/');
        }
    </script>

</head>

<body class="" ssr="n">

    <div id="wrapper" class="wrapper_new">
        <div id="head">
            <div id="s_top_wrap" class="s-top-wrap s-isindex-wrap ">
                <div class="s-top-nav"></div>
                <div class="s-center-box"></div>
            </div>
            <div id="u"><a class="toindex" href="/">百度首页</a><a href="javascript:;" name="tj_settingicon"
                    class="pf">设置<i class="c-icon c-icon-triangle-down"></i></a><a
                    href="https://passport.baidu.com/v2/?login&tpl=mn&u=http%3A%2F%2Fwww.baidu.com%2F&sms=5"
                    name="tj_login" class="lb" onclick="return false;">登录</a>
                <div class="bdpfmenu"></div>
            </div>
            <div id="s-top-left" class="s-top-left-new s-isindex-wrap"><a href="http://news.baidu.com" target="_blank"
                    class="mnav c-font-normal c-color-t">新闻</a><a href="https://www.hao123.com?src=from_pc"
                    target="_blank" class="mnav c-font-normal c-color-t">hao123</a><a href="http://map.baidu.com"
                    target="_blank" class="mnav c-font-normal c-color-t">地图</a><a href="http://tieba.baidu.com/"
                    target="_blank" class="mnav c-font-normal c-color-t">贴吧</a><a
                    href="https://haokan.baidu.com/?sfrom=baidu-top" target="_blank"
                    class="mnav c-font-normal c-color-t">视频</a><a href="http://image.baidu.com/" target="_blank"
                    class="mnav c-font-normal c-color-t">图片</a><a href="https://pan.baidu.com?from=1026962h"
                    target="_blank" class="mnav c-font-normal c-color-t">网盘</a>
                <div class="mnav s-top-more-btn"><a href="http://www.baidu.com/more/" name="tj_briicon"
                        class="s-bri c-font-normal c-color-t" target="_blank">更多</a>
                    <div class='s-top-more' id="s-top-more">
                        <div class='s-top-more-content row-1 clearfix'><a class="img-spacing"
                                href='http://fanyi.baidu.com/' target='_blank' name='tj_fanyi'><img src='https://dss0.bdstatic.com/5aV1bjqh_Q23odCf/static/superman/img/topnav/newfanyi-da0cea8f7e.png'/>
                                <div class="s-top-more-title c-font-normal c-color-t">翻译</div>
                            </a><a class="img-spacing" href='http://xueshu.baidu.com/' target='_blank'
                                name='tj_xueshu'><img src='https://dss0.bdstatic.com/5aV1bjqh_Q23odCf/static/superman/img/topnav/newxueshuicon-a5314d5c83.png'/>
                                <div class="s-top-more-title c-font-normal c-color-t">学术</div>
                            </a><a class="" href='https://wenku.baidu.com' target='_blank' name='tj_wenku'><img src='https://dss0.bdstatic.com/5aV1bjqh_Q23odCf/static/superman/img/topnav/newwenku-d8c9b7b0fb.png'/>
                                <div class="s-top-more-title c-font-normal c-color-t">文库</div>
                            </a></div>
                        <div class='s-top-more-content row-2 clearfix'><a class="img-spacing"
                                href='https://baike.baidu.com' target='_blank' name='tj_baike'><img src='https://dss0.bdstatic.com/5aV1bjqh_Q23odCf/static/superman/img/topnav/newbaike-889054f349.png'/>
                                <div class="s-top-more-title c-font-normal c-color-t">百科</div>
                            </a><a class="img-spacing" href='https://zhidao.baidu.com' target='_blank'
                                name='tj_zhidao'><img src='https://dss0.bdstatic.com/5aV1bjqh_Q23odCf/static/superman/img/topnav/newzhidao-da1cf444b0.png'/>
                                <div class="s-top-more-title c-font-normal c-color-t">知道</div>
                            </a><a class="" href='https://jiankang.baidu.com/widescreen/home' target='_blank'
                                name='tj_jiankang'><img src='https://dss0.bdstatic.com/5aV1bjqh_Q23odCf/static/superman/img/topnav/newjiankang-f03b804b4b.png'/>
                                <div class="s-top-more-title c-font-normal c-color-t">健康</div>
                            </a></div>
                        <div class='s-top-more-content row-3 clearfix'><a class="img-spacing"
                                href='http://e.baidu.com/ebaidu/home?refer=887' target='_blank'
                                name='tj_yingxiaotuiguang'><img src='https://dss0.bdstatic.com/5aV1bjqh_Q23odCf/static/superman/img/topnav/yingxiaoicon-612169cc36.png'/>
                                <div class="s-top-more-title c-font-normal c-color-t">营销推广</div>
                            </a><a class="img-spacing" href='https://live.baidu.com/' target='_blank'
                                name='tj_live'><img src='https://dss0.bdstatic.com/5aV1bjqh_Q23odCf/static/superman/img/topnav/newzhibo-a6a0831ecd.png'/>
                                <div class="s-top-more-title c-font-normal c-color-t">直播</div>
                            </a><a class="" href='http://music.taihe.com' target='_blank' name='tj_mp3'><img src='https://dss0.bdstatic.com/5aV1bjqh_Q23odCf/static/superman/img/topnav/newyinyue-03ecd1e9b9.png'/>
                                <div class="s-top-more-title c-font-normal c-color-t">音乐</div>
                            </a></div>
                        <div class="s-top-tomore"><a class="c-color-gray2 c-font-normal"
                                href='http://www.baidu.com/more/' target='_blank' name='tj_more'>查看全部百度产品 ></a></div>
                    </div>
                </div>
            </div>
            <div id="u1" class="s-top-right s-isindex-wrap">
                <span class="s-top-right-text c-font-normal c-color-t s-top-right-new" id="s-usersetting-top" name="tj_settingicon">设置</span><a
                    class="s-top-login-btn c-btn c-btn-primary c-btn-mini lb"
                    style="position:relative;overflow: visible;" id="s-top-loginbtn"
                    href="https://passport.baidu.com/v2/?login&tpl=mn&u=http%3A%2F%2Fwww.baidu.com%2F&sms=5"
                    name="tj_login" onclick="return false;">登录</a>
                <div id="s-user-setting-menu" class="s-top-userset-menu c-floating-box c-font-normal ">
                    <div class="s-user-setting-pfmenu"></div><a class="s-set-hotsearch set-hide"
                        href="javascript:;">关闭热搜</a><a class="s-set-hotsearch set-show" href="javascript:;">开启热搜</a>
                </div>
                <div class="guide-info ">
                    <i class="c-icon guide-icon">&#xe625;</i><span>牛年贺岁，登录设置新春皮肤！</span><i class="c-icon guide-close">&#xe610;</i>
                </div>
            </div>
            <div id="head_wrapper" class="head_wrapper s-isindex-wrap nologin ">
                <div class="s_form s_form_nologin">
                    <div class="s_form_wrapper">
                        <style>
                            .index-logo-srcnew,
                            .index-logo-peak {
                                display: none;
                            }

                            @media (-webkit-min-device-pixel-ratio: 2),
                            (min--moz-device-pixel-ratio: 2),
                            (-o-min-device-pixel-ratio: 2),
                            (min-device-pixel-ratio: 2) {
                                .index-logo-src {
                                    display: none;
                                }

                                .index-logo-srcnew {
                                    display: inline;
                                }
                            }
                        </style>
                        <div id="lg" class="s-p-top">
                            <img hidefocus="true" id="s_lg_img" class='index-logo-src' src="//www.baidu.com/img/PCtm_d9c8750bed0b3c7d089fa7d55720d6cf.png" width="270" height="129" onerror="this.src='//www.baidu.com/img/flexible/logo/pc/index.png';this.onerror=null;" usemap="#mp"><img hidefocus="true" id="s_lg_img_new" class='index-logo-srcnew' src="//www.baidu.com/img/PCtm_d9c8750bed0b3c7d089fa7d55720d6cf.png" width="270" height="129" onerror="this.src='//www.baidu.com/img/flexible/logo/pc/index@2.png';this.onerror=null;" usemap="#mp"><img hidefocus="true" id="s_lg_img_aging" class='index-logo-aging-tools' src="//www.baidu.com/img/PCfb_5bf082d29588c07f842ccde3f97243ea.png" width="270" height="129" onerror="this.src='//www.baidu.com/img/flexible/logo/pc/index@2.png';this.onerror=null;" usemap="#mp"><map name="mp"><area style="outline:none;" hidefocus="true" shape="rect" coords="0,0,270,129" href="//www.baidu.com/s?wd=%E7%99%BE%E5%BA%A6%E7%83%AD%E6%90%9C&amp;sa=ire_dl_gh_logo_texing&amp;rsv_dl=igh_logo_pcs" onmousedown="return ns_c({fm: 'tab', tab: 'felogo', rsv_platform: 'wwwhome' })" target="_blank" title="点击一下，了解更多"onmousedown=&quot;return ns_c({&#39;fm&#39;:&#39;behs&#39;,&#39;tab&#39;:&#39;bdlogo&#39;})&quot;></map>
                        </div><a href="/" id="result_logo"
                            onmousedown="return c({'fm':'tab','tab':'logo'})"><img class='index-logo-src' src="//www.baidu.com/img/flexible/logo/pc/result.png" alt="到百度首页" title="到百度首页"><img class='index-logo-srcnew' src="//www.baidu.com/img/flexible/logo/pc/result@2.png" alt="到百度首页" title="到百度首页"><img class='index-logo-peak' src="//www.baidu.com/img/flexible/logo/pc/peak-result.png" alt="到百度首页" title="到百度首页"></a>
                        <form id="form" name="f" action="/s" class="fm ">
                            <input type="hidden" name="ie" value="utf-8"><input type="hidden" name="f" value="8"><input type="hidden" name="rsv_bp" value="1"><input type="hidden" name="rsv_idx" value="1"><input type=hidden name=ch value=""><input type=hidden name=tn value="baidu"><input type=hidden name=bar value=""><span class="bg s_ipt_wr new-pmd quickdelete-wrap"><input id="kw" name="wd" class="s_ipt" value="" maxlength="255" autocomplete="off"><i class="c-icon quickdelete c-color-gray2" title="清空">&#xe610;</i><i class="quickdelete-line"></i></span><span class="bg s_btn_wr"><input type="submit" id="su" value="百度一下" class="bg s_btn"></span><span class="tools"><span id="mHolder"><div id="mCon"><span>
                                    输入法</span>
                    </div>
                    <ul id="mMenu">
                        <li><a href="javascript:;" name="ime_hw">
                                手写</a></li>
                        <li><a href="javascript:;" name="ime_py">
                                拼音</a></li>
                        <li class="ln"></li>
                        <li><a href="javascript:;" name="ime_cl">
                                关闭</a></li>
                    </ul>
                    </span></span><input type="hidden" name="rn" value=""><input type="hidden" name="fenlei" value="256"><input type="hidden" name="oq" value=""><input type="hidden" name="rsv_pq" value="0x9bd58bc10006da91"><input type="hidden" name="rsv_t" value="71c9FEZ0nREpY810sr2pp3TRf/3y+mCw5BxsgoGIQRJq9ByDWTVExmGFf9Jb"><input type="hidden" name="rqlang" value="en">
                    </form>
                    <div id="m" class="under-searchbox-tips s_lm_hide ">
                        <div id="lm-new"></div>
                    </div>
                    <div id="s-hotsearch-wrapper" class="s-isindex-wrap s-hotsearch-wrapper hide ">
                        <div class="s-hotsearch-title"><a class="hot-title"
                                href="https://top.baidu.com/board?platform=pc&sa=pcindex_entry" target="_blank">
                                <div class="title-text c-font-medium c-color-t" aria-label="百度热搜">
                                    <i class="c-icon">&#xe687;</i><i class="c-icon arrow">&#xe613;</i>
                                </div>
                            </a><a id="hotsearch-refresh-btn"
                                class="hot-refresh c-font-normal c-color-gray2"><i class="c-icon refresh-icon">&#xe619;</i><span class="hot-refresh-text">换一换</span></a>
                        </div>
                        <ul class="s-hotsearch-content" id="hotsearch-content-wrapper">
                            <li class="hotsearch-item odd" data-index="0"><a
                                    class="title-content  c-link c-font-medium c-line-clamp1"
                                    href="https://www.baidu.com/s?wd=%E4%B9%A1%E6%9D%91%E6%8C%AF%E5%85%B4%E5%BD%A2%E6%88%90%E6%96%B0%E6%A0%BC%E5%B1%80&amp;sa=fyb_n_homepage&amp;rsv_dl=fyb_n_homepage&amp;from=super&amp;cl=3&amp;tn=baidutop10&amp;fr=top1000&amp;rsv_idx=2&amp;hisfilter=1"
                                    target="_blank">
                                    <div class="title-content-noindex" style="display: none;"></div>
                                    <i class="c-icon title-content-top-icon c-color-red c-gap-right-small" style="display: ;">&#xe62e;</i><span class="title-content-index c-index-single c-index-single-hot0" style="display: none;">0</span><span class="title-content-title">乡村振兴形成新格局</span>
                                </a><span class="title-content-mark ie-vertical c-text c-gap-left-small "></span></li>
                            <li class="hotsearch-item even" data-index="3"><a
                                    class="title-content  c-link c-font-medium c-line-clamp1"
                                    href="https://www.baidu.com/s?wd=%E5%AE%88%E6%9C%9B%E7%9B%B8%E5%8A%A9+%E4%BC%A0%E9%80%92%E6%B8%A9%E6%9A%96&amp;sa=fyb_n_homepage&amp;rsv_dl=fyb_n_homepage&amp;from=super&amp;cl=3&amp;tn=baidutop10&amp;fr=top1000&amp;rsv_idx=2&amp;hisfilter=1"
                                    target="_blank">
                                    <div class="title-content-noindex" style="display: none;"></div>
                                    <i class="c-icon title-content-top-icon c-color-red c-gap-right-small" style="display: none;">&#xe62e;</i><span class="title-content-index c-index-single c-index-single-hot3" style="display: ;">3</span><span class="title-content-title">守望相助 传递温暖</span>
                                </a><span class="title-content-mark ie-vertical c-text c-gap-left-small "></span></li>
                            <li class="hotsearch-item odd" data-index="1"><a
                                    class="title-content tag-width c-link c-font-medium c-line-clamp1"
                                    href="https://www.baidu.com/s?wd=%E9%98%B3%E5%BA%B7%E5%90%8E%E5%89%A7%E7%83%88%E8%BF%90%E5%8A%A8%E4%BC%9A%E7%8C%9D%E6%AD%BB%EF%BC%9F%E4%B8%93%E5%AE%B6%E5%9B%9E%E5%BA%94&amp;sa=fyb_n_homepage&amp;rsv_dl=fyb_n_homepage&amp;from=super&amp;cl=3&amp;tn=baidutop10&amp;fr=top1000&amp;rsv_idx=2&amp;hisfilter=1"
                                    target="_blank">
                                    <div class="title-content-noindex" style="display: none;"></div>
                                    <i class="c-icon title-content-top-icon c-color-red c-gap-right-small" style="display: none;">&#xe62e;</i><span class="title-content-index c-index-single c-index-single-hot1" style="display: ;">1</span><span class="title-content-title">阳康后剧烈运动会猝死？专家回应</span>
                                </a><span class="title-content-mark ie-vertical c-text c-gap-left-small c-text-hot">热</span>
                            </li>
                            <li class="hotsearch-item even" data-index="4"><a
                                    class="title-content  c-link c-font-medium c-line-clamp1"
                                    href="https://www.baidu.com/s?wd=%E8%B4%9D%E5%88%A9%E7%97%85%E6%83%85%E7%BB%A7%E7%BB%AD%E6%81%B6%E5%8C%96+%E5%AE%B6%E4%BA%BA%E5%BC%80%E5%A7%8B%E7%AD%B9%E5%A4%87%E8%91%AC%E7%A4%BC&amp;sa=fyb_n_homepage&amp;rsv_dl=fyb_n_homepage&amp;from=super&amp;cl=3&amp;tn=baidutop10&amp;fr=top1000&amp;rsv_idx=2&amp;hisfilter=1"
                                    target="_blank">
                                    <div class="title-content-noindex" style="display: none;"></div>
                                    <i class="c-icon title-content-top-icon c-color-red c-gap-right-small" style="display: none;">&#xe62e;</i><span class="title-content-index c-index-single c-index-single-hot4" style="display: ;">4</span><span class="title-content-title">贝利病情继续恶化 家人开始筹备葬礼</span>
                                </a><span class="title-content-mark ie-vertical c-text c-gap-left-small "></span></li>
                            <li class="hotsearch-item odd" data-index="2"><a
                                    class="title-content tag-width c-link c-font-medium c-line-clamp1"
                                    href="https://www.baidu.com/s?wd=%E5%8C%97%E5%A4%A7%E6%95%99%E6%8E%88%EF%BC%9A%E5%81%A5%E5%BA%B7%E7%A0%81%E5%BD%BB%E5%BA%95%E9%80%80%E5%87%BA%E5%8A%BF%E5%9C%A8%E5%BF%85%E8%A1%8C&amp;sa=fyb_n_homepage&amp;rsv_dl=fyb_n_homepage&amp;from=super&amp;cl=3&amp;tn=baidutop10&amp;fr=top1000&amp;rsv_idx=2&amp;hisfilter=1"
                                    target="_blank">
                                    <div class="title-content-noindex" style="display: none;"></div>
                                    <i class="c-icon title-content-top-icon c-color-red c-gap-right-small" style="display: none;">&#xe62e;</i><span class="title-content-index c-index-single c-index-single-hot2" style="display: ;">2</span><span class="title-content-title">北大教授：健康码彻底退出势在必行</span>
                                </a><span class="title-content-mark ie-vertical c-text c-gap-left-small c-text-hot">热</span>
                            </li>
                            <li class="hotsearch-item even" data-index="5"><a
                                    class="title-content  c-link c-font-medium c-line-clamp1"
                                    href="https://www.baidu.com/s?wd=%E5%80%92%E5%8D%96%E9%98%B2%E7%96%AB%E7%89%A9%E8%B5%84%E7%9A%84%E5%8E%BF%E5%A7%94%E4%B9%A6%E8%AE%B0%E5%87%BA%E9%95%9C%E5%BF%8F%E6%82%94&amp;sa=fyb_n_homepage&amp;rsv_dl=fyb_n_homepage&amp;from=super&amp;cl=3&amp;tn=baidutop10&amp;fr=top1000&amp;rsv_idx=2&amp;hisfilter=1"
                                    target="_blank">
                                    <div class="title-content-noindex" style="display: none;"></div>
                                    <i class="c-icon title-content-top-icon c-color-red c-gap-right-small" style="display: none;">&#xe62e;</i><span class="title-content-index c-index-single c-index-single-hot5" style="display: ;">5</span><span class="title-content-title">倒卖防疫物资的县委书记出镜忏悔</span>
                                </a><span class="title-content-mark ie-vertical c-text c-gap-left-small "></span></li>
                        </ul>
                    </div>
                </div>
            </div>
        </div>
        <div id="s_wrap" class="s-isindex-wrap">
            <div id="s_main" class="main clearfix "></div>
        </div>
        <div id="bottom_layer" class="s-bottom-layer s-isindex-wrap">
            <div class="s-bottom-layer-content">
                <p class="lh"><a class="text-color" href="//home.baidu.com" target="_blank">关于百度</a></p>
                <p class="lh"><a class="text-color" href="http://ir.baidu.com" target="_blank">About Baidu</a></p>
                <p class="lh"><a class="text-color" href="//www.baidu.com/duty" target="_blank">使用百度前必读</a></p>
                <p class="lh"><a class="text-color" href="//help.baidu.com" target="_blank">帮助中心</a></p>
                <p class="lh"><a class="text-color" href="https://e.baidu.com/?refer=1271" target="_blank">企业推广</a></p>
                <p class="lh"><a class="text-color"
                        href="http://www.beian.gov.cn/portal/registerSystemInfo?recordcode=11000002000001"
                        target="_blank">京公网安备11000002000001号</a></p>
                <p class="lh"><a class="text-color" href="https://beian.miit.gov.cn" target="_blank">京ICP证030173号</a>
                </p>
                <p class="lh"><a class="text-color" href="//www.baidu.com/licence/" target="_blank">信息网络传播视听节目许可证
                        0110516</a></p>
                <p class="lh"><span class="text-color">互联网宗教信息服务许可证编号：京（2022）0000043</span></p>
                <p class="lh"><span class="text-color">药品医疗器械网络信息服务备案（京）网药械信息备字（2021）第00159号</span></p>
                <p class="lh"><span class="text-color">医疗器械网络交易服务第三方平台备案凭证（京）网械平台备字（2020）第00002号</span></p>
                <p class="lh"><span class="text-color">&#169;2022&nbsp;Baidu&nbsp;</span></p>
            </div>
        </div>
        <div id="bottom_space" class="s-bottom-space"></div>
        <script type="application/json" id="promote_login_box">
            []
        </script>
    </div>
    <div class="s_tab " id="s_tab">
        <div class="s_tab_inner"><b class="cur-tab">网页</b><a
                href="https://www.baidu.com/s?rtt=1&amp;bsst=1&amp;cl=2&amp;tn=news" wdfield="word"
                onmousedown="return c({'fm':'tab','tab':'news'})" sync="true" class="s-tab-item s-tab-news">资讯</a><a
                href="http://v.baidu.com/v?ct=301989888&amp;rn=20&amp;pn=0&amp;db=0&amp;s=25&amp;ie=utf-8"
                wdfield="word" onmousedown="return c({'fm':'tab','tab':'video'})"
                class="s-tab-item s-tab-video">视频</a><a
                href="http://image.baidu.com/i?tn=baiduimage&amp;ps=1&amp;ct=201326592&amp;lm=-1&amp;cl=2&amp;nc=1&amp;ie=utf-8"
                wdfield="word" onmousedown="return c({'fm':'tab','tab':'pic'})" class="s-tab-item s-tab-pic">图片</a><a
                href="http://zhidao.baidu.com/q?ct=17&amp;pn=0&amp;tn=ikaslist&amp;rn=10&amp;fr=wwwt" wdfield="word"
                onmousedown="return c({'fm':'tab','tab':'zhidao'})" class="s-tab-item s-tab-zhidao">知道</a><a
                href="http://wenku.baidu.com/search?lm=0&amp;od=0&amp;ie=utf-8" wdfield="word"
                onmousedown="return c({'fm':'tab','tab':'wenku'})" class="s-tab-item s-tab-wenku">文库</a><a
                href="http://tieba.baidu.com/f?fr=wwwt" wdfield="kw" onmousedown="return c({'fm':'tab','tab':'tieba'})"
                class="s-tab-item s-tab-tieba">贴吧</a><a href="https://map.baidu.com/?newmap=1&amp;ie=utf-8&amp;s=s"
                onmousedown="return c({'fm':'tab','tab':'map'})" class="s-tab-item s-tab-map">地图</a><a
                href="https://b2b.baidu.com/s?fr=wwwt" wdfield="q" onmousedown="return c({'fm':'tab','tab':'b2b'})"
                class="s-tab-item s-tab-b2b">采购</a><a href="http://www.baidu.com/more/"
                onmousedown="return c({'fm':'tab','tab':'more'})" class="s-tab-item s-tab-more">更多</a></div>
    </div>
    <div id="s_side_wrapper">
        <div class="side-entry aging-entry">
            <div class="aging-entry-inner"></div>
            <div class="c-color-text toast">辅助模式</div>
        </div>
        <div id="s_qrcode_nologin" class="qrcode-nologin side-entry">
            <div class="qrcode-layer icon-mask-wrapper">
                <img class="icon" src="https://dss0.bdstatic.com/5aV1bjqh_Q23odCf/static/superman/img/qrcode/qrcode@2x-daf987ad02.png"/><img class="icon-hover" src="https://dss0.bdstatic.com/5aV1bjqh_Q23odCf/static/superman/img/qrcode/qrcode-hover@2x-f9b106a848.png"/>
            </div>
            <div class="tooltip qrcode-tooltip">
                <div class="text">
                    <div class="login-text"><i class="c-icon login-icon">&#xe602;</i>百度APP扫码登录</div>
                    <div class="login-info">百度一下&nbsp;生活更好</div>
                </div>
                <div id="qrcode-login-wrapper"></div>
            </div>
        </div>
    </div>
    <div id="wrapper_wrapper"></div>
    </div>
    <div class="c-tips-container" id="c-tips-container"></div>

</body>

</html>
`
