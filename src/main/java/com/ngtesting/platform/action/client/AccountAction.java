package com.ngtesting.platform.action.client;

import com.alibaba.fastjson.JSONObject;
import com.ngtesting.platform.config.Constant;
import com.ngtesting.platform.model.TstUser;
import com.ngtesting.platform.service.AccountService;
import com.ngtesting.platform.service.AccountVerifyCodeService;
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
    private AccountVerifyCodeService accountVerifyCodeService;

    @Autowired
    private OrgService orgService;
    @Autowired
    private UserService userService;

    @AuthPassport(validate=false)
    @ResponseBody
    @PostMapping("/register")
    public Map register(HttpServletRequest request, @RequestBody TstUser json){
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
    @PostMapping("/loginWithVerifyCode")
    public Object loginWithVerifyCode(HttpServletRequest request, @RequestBody JSONObject json){
        Map<String, Object> ret = new HashMap<String, Object>();

        String vcode = json.getString("vcode");
        TstUser user = accountService.loginWithVerifyCode(vcode);

        if (user != null) {
            request.getSession().setAttribute(Constant.HTTP_SESSION_USER_PROFILE, user);

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
    @PostMapping("/login")
    public Object login(HttpServletRequest request, @RequestBody JSONObject json){
        Map<String, Object> ret = new HashMap<String, Object>();

        String email = json.getString("email");
        String password = json.getString("password");
        boolean rememberMe = json.getBoolean("rememberMe") != null? json.getBoolean("rememberMe"): false;

        TstUser user = accountService.login(email, password, rememberMe);

        if (user != null) {
            request.getSession().setAttribute(Constant.HTTP_SESSION_USER_PROFILE, user);

            ret.put("profile", user);
            ret.put("token", user.getToken());
            ret.put("code", Constant.RespCode.SUCCESS.getCode());
        } else {
            ret.put("code", Constant.RespCode.BIZ_FAIL.getCode());
            ret.put("msg", "登录失败");
        }

        return ret;
    }

    @ResponseBody
    @PostMapping("/logout")
    public Object logout(HttpServletRequest request, @RequestBody JSONObject json){
        Map<String, Object> ret = new HashMap<String, Object>();
        TstUser user = (TstUser) request.getSession().getAttribute(Constant.HTTP_SESSION_USER_PROFILE);
        if (user == null) {
            ret.put("msg", "您不在登录状态");
        } else {
            Boolean result = accountService.logout(user.getEmail());

            if (result) {
                request.getSession().removeAttribute(Constant.HTTP_SESSION_USER_PROFILE);
                ret.put("msg", "登出成功");
                ret.put("code", Constant.RespCode.SUCCESS.getCode());
            } else {
                ret.put("msg", "登出失败");
                ret.put("code", Constant.RespCode.BIZ_FAIL.getCode());
            }
        }

        return ret;
    }

    @ResponseBody
    @PostMapping("/changePassword")
    public Object changePassword(HttpServletRequest request, @RequestBody JSONObject json){
        Map<String, Object> ret = new HashMap<String, Object>();

        TstUser user = (TstUser) request.getSession().getAttribute(Constant.HTTP_SESSION_USER_PROFILE);
        String oldPassword = json.getString("oldPassword");
        String password = json.getString("password");

        boolean success = accountService.changePassword(user.getId(), oldPassword, password);
        int code = success? Constant.RespCode.SUCCESS.getCode(): Constant.RespCode.BIZ_FAIL.getCode();

        ret.put("code", code);
        return ret;
    }

    @AuthPassport(validate=false)
    @ResponseBody
    @PostMapping("/forgotPassword")
    public Object forgotPassword(HttpServletRequest request, @RequestBody JSONObject json){
        Map<String, Object> ret = new HashMap<String, Object>();

        String email = json.getString("email");
        TstUser user = userService.getByEmail(email);

        if (user == null) {
            ret.put("code", Constant.RespCode.BIZ_FAIL.getCode());
            ret.put("msg", "用户不存在");
            return ret;
        }

        String verifyCode = accountService.forgotPassword(user);

        ret.put("data", verifyCode);
        ret.put("code", Constant.RespCode.SUCCESS.getCode());

        return ret;
    }

    @AuthPassport(validate=false)
    @ResponseBody
    @PostMapping("/checkResetPassword")
    public Object checkResetPassword(HttpServletRequest request, @RequestBody JSONObject json){
        Map<String, Object> ret = new HashMap<String, Object>();

        String verifyCode = json.getString("vcode");

        boolean success = accountService.beforResetPassword(verifyCode);
        if (success) {
            ret.put("code", Constant.RespCode.SUCCESS.getCode());
        } else {
            ret.put("code", Constant.RespCode.BIZ_FAIL.getCode());
            ret.put("msg", "重置密码链接已超时失效");
        }

        return ret;
    }

    @AuthPassport(validate=false)
    @ResponseBody
    @PostMapping("/resetPassword")
    public Object resetPassword(HttpServletRequest request, @RequestBody JSONObject json){
        Map<String, Object> ret = new HashMap<String, Object>();

        String verifyCode = json.getString("vcode");
        String password = json.getString("password");

        TstUser user = accountService.resetPassword(verifyCode, password);

        if (user != null) {
            request.getSession().setAttribute(Constant.HTTP_SESSION_USER_PROFILE, user);

            ret.put("token", user.getToken());
            ret.put("code", Constant.RespCode.SUCCESS.getCode());
        } else {
            ret.put("code", Constant.RespCode.BIZ_FAIL.getCode());
            ret.put("msg", "重置密码失败");
        }

        return ret;
    }
}
