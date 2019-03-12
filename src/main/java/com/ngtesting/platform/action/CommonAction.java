package com.ngtesting.platform.action;

import com.ngtesting.platform.config.Constant;
import org.apache.shiro.SecurityUtils;
import org.springframework.web.bind.annotation.RequestMapping;
import org.springframework.web.bind.annotation.RestController;

import java.util.HashMap;
import java.util.Map;

@RequestMapping(Constant.API_PATH + "common")
@RestController
public class CommonAction {
    @RequestMapping("unAuthen")
    public Map<String, Object> unAuthen() {
        Map<String, Object> ret = new HashMap<String, Object>();

        SecurityUtils.getSubject().logout();
        ret.put("code", Constant.RespCode.AUTHEN_FAIL.getCode());
        ret.put("msg", "认证错误");
        return ret;
    }

    @RequestMapping("unAuthor")
    public Map<String, Object> unAuthor() {
        Map<String, Object> ret = new HashMap<String, Object>();

        SecurityUtils.getSubject().logout();
        ret.put("code", Constant.RespCode.AUTHOR_FAIL.getCode());
        ret.put("msg", "授权错误");
        return ret;
    }

    @RequestMapping("kickout")
    public Map<String, Object> kickout(){
        Map<String, Object> ret = new HashMap<String, Object>();

        ret.put("code", Constant.RespCode.BIZ_FAIL.getCode());
        ret.put("msg", "您已在别处登录");
        return ret;
    }

}
