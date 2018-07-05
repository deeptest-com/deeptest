package com.ngtesting.platform.action;

import com.alibaba.fastjson.JSONObject;
import com.ngtesting.platform.config.Constant;
import com.ngtesting.platform.model.TstUser;
import com.ngtesting.platform.service.AccountService;
import com.ngtesting.platform.service.OrgService;
import com.ngtesting.platform.service.UserService;
import com.ngtesting.platform.utils.AuthPassport;
import org.springframework.beans.factory.annotation.Autowired;
import org.springframework.stereotype.Controller;
import org.springframework.web.bind.annotation.PostMapping;
import org.springframework.web.bind.annotation.RequestBody;
import org.springframework.web.bind.annotation.RequestMapping;
import org.springframework.web.bind.annotation.ResponseBody;

import javax.servlet.http.HttpServletRequest;
import java.util.HashMap;
import java.util.Map;

@Controller
@RequestMapping(value = Constant.API_PATH_CLIENT + "/account")
public class AccountAction {
    @Autowired
    private AccountService accountService;
    @Autowired
    private OrgService orgService;
    @Autowired
    private UserService userService;

    @AuthPassport(validate=false)
    @ResponseBody
    @PostMapping("/register")
    public Map register(@RequestBody TstUser json){
        Map<String, Object> ret = new HashMap();
        TstUser user = accountService.register(json);

        if (user != null) {
            ret.put("msg", "注册成功，请访问您的邮箱进行登录");
            ret.put("code", Constant.RespCode.SUCCESS.getCode());
        } else {
            ret.put("code", Constant.RespCode.BIZ_FAIL.getCode());
            ret.put("msg", "邮箱已存在");
        }

        return ret;
    }

    @AuthPassport(validate=false)
    @ResponseBody
    @PostMapping("/login")
    public Object login(@RequestBody JSONObject json, HttpServletRequest request){
        Map<String, Object> ret = new HashMap<String, Object>();

        String email = json.getString("email");
        String password = json.getString("password");
        boolean rememberMe = json.getBoolean("rememberMe") != null? json.getBoolean("rememberMe"): false;

        TstUser user = accountService.login(email, password, rememberMe);

        if (user != null) {
            request.getSession().setAttribute(Constant.HTTP_SESSION_USER_KEY, user);

            ret.put("profile", user);
            ret.put("token", user.getToken());
            ret.put("code", Constant.RespCode.SUCCESS.getCode());
        } else {
            ret.put("code", Constant.RespCode.BIZ_FAIL.getCode());
            ret.put("msg", "登录失败");
        }

        return ret;
    }

    @AuthPassport(validate=false)
    @ResponseBody
    @PostMapping("/loginWithVerifyCode")
    public Object loginWithVerifyCode(@RequestBody JSONObject json){
        Map<String, Object> ret = new HashMap();

        ret.put("code", Constant.RespCode.SUCCESS.getCode());
        return ret;
    }

    @ResponseBody
    @PostMapping("/logout")
    public Object logout(TstUser user){
        return null;
    }

    @ResponseBody
    @PostMapping("/changePassword")
    public Object changePassword(TstUser user){
        return null;
    }

    @AuthPassport(validate=false)
    @ResponseBody
    @PostMapping("/checkResetPassword")
    public Object checkResetPassword(TstUser user){
        return null;
    }

    @AuthPassport(validate=false)
    @ResponseBody
    @PostMapping("/resetPassword")
    public Object resetPassword(TstUser user){
        return null;
    }

    @AuthPassport(validate=false)
    @ResponseBody
    @PostMapping("/forgotPassword")
    public Object forgotPassword(TstUser user){
        return null;
    }
}
