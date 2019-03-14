package com.ngtesting.platform.action.client;

import com.alibaba.fastjson.JSON;
import com.alibaba.fastjson.JSONObject;
import com.ngtesting.platform.config.Constant;
import com.ngtesting.platform.model.TstUser;
import com.ngtesting.platform.service.intf.AccountService;
import com.ngtesting.platform.service.intf.UserService;
import com.ngtesting.platform.utils.StringUtil;
import org.apache.shiro.SecurityUtils;
import org.apache.shiro.authc.IncorrectCredentialsException;
import org.apache.shiro.authc.LockedAccountException;
import org.apache.shiro.authc.UnknownAccountException;
import org.apache.shiro.authc.UsernamePasswordToken;
import org.apache.shiro.subject.Subject;
import org.springframework.beans.factory.annotation.Autowired;
import org.springframework.web.bind.annotation.PostMapping;
import org.springframework.web.bind.annotation.RequestBody;
import org.springframework.web.bind.annotation.RequestMapping;
import org.springframework.web.bind.annotation.RestController;

import javax.servlet.http.HttpServletRequest;
import java.util.HashMap;
import java.util.Map;

@RestController
@RequestMapping(value = Constant.API_PATH_CLIENT + "/account")
public class AccountAction {
    @Autowired
    private AccountService accountService;

    @Autowired
    private UserService userService;

    @PostMapping("/register")
    public Map register(HttpServletRequest request, @RequestBody JSONObject json) {
        Map<String, Object> ret = new HashMap();

        TstUser vo = JSON.parseObject(JSON.toJSONString(json), TstUser.class);
        TstUser user = accountService.register(vo);
        if (user != null) {
            UsernamePasswordToken token = new UsernamePasswordToken(user.getEmail(), user.getPassword(), true);

            ret.put("msg", "注册成功，请访问您的邮箱进行登录");
            ret.put("code", Constant.RespCode.SUCCESS.getCode());
        } else {
            ret.put("code", Constant.RespCode.BIZ_FAIL.getCode());
            ret.put("msg", "邮箱已存在");
        }

        return ret;
    }

    @PostMapping("/loginWithVerifyCode")
    public Object loginWithVerifyCode(HttpServletRequest request, @RequestBody JSONObject json) {
        Map<String, Object> ret = new HashMap<String, Object>();

        String vcode = json.getString("vcode");
        TstUser user = accountService.loginWithVerifyCode(vcode);

        if (user != null) {
            UsernamePasswordToken token = new UsernamePasswordToken(user.getEmail(), "", true);
            //登录不在该处处理，交由shiro处理
            Subject subject = SecurityUtils.getSubject();
            subject.login(token);

            if (subject.isAuthenticated()) {
                ret.put("orgId", user.getDefaultOrgId());
                ret.put("token", subject.getSession().getId());
                ret.put("code", Constant.RespCode.SUCCESS.getCode());
            }
        } else {
            throw new IncorrectCredentialsException();
        }

        if (ret.get("code") == null) {
            ret.put("code", Constant.RespCode.BIZ_FAIL.getCode());
            ret.put("msg", "VerifyCode登录失败");
        }

        return ret;
    }

    @PostMapping("/login")
    public Object login(HttpServletRequest request, @RequestBody JSONObject json) {
        Map<String, Object> ret = new HashMap<String, Object>();

        String email = json.getString("email");
        String password = json.getString("password");
        boolean rememberMe = json.getBoolean("rememberMe") != null ? json.getBoolean("rememberMe") : false;

        if (StringUtil.isEmpty(password)) {
            throw new IncorrectCredentialsException();
        }

        String msg = "";
        try {
            UsernamePasswordToken token = new UsernamePasswordToken(email, password, rememberMe);
            //登录不在该处处理，交由shiro处理
            Subject subject = SecurityUtils.getSubject();
            subject.login(token);

            if (subject.isAuthenticated()) {
                TstUser user = userService.getByEmail(email);
                ret.put("token", subject.getSession().getId());
                ret.put("orgId", user.getDefaultOrgId());
                ret.put("code", Constant.RespCode.SUCCESS.getCode());
            }else{
                msg = "登录异常";
            }
        }catch (IncorrectCredentialsException | UnknownAccountException e){
            e.printStackTrace();
            msg = "该用户不存在或密码错误";
        }catch (LockedAccountException e){
            e.printStackTrace();
            msg = "该用户已被冻结";
        }catch (Exception e){
            e.printStackTrace();
            msg = "服务器错误";
        }

        if (ret.get("code") == null) {
            ret.put("code", Constant.RespCode.BIZ_FAIL.getCode());
            ret.put("msg", msg);
        }
        return ret;
    }

    @PostMapping("/logout")
    public Object logout(HttpServletRequest request, @RequestBody JSONObject json) {
        Map<String, Object> ret = new HashMap<String, Object>();

        SecurityUtils.getSubject().logout();

        ret.put("msg", "登出成功");
        ret.put("code", Constant.RespCode.SUCCESS.getCode());
        return ret;
    }

    @PostMapping("/changePassword")
    public Object changePassword(HttpServletRequest request, @RequestBody JSONObject json) {
        Map<String, Object> ret = new HashMap<String, Object>();

        TstUser user = (TstUser) SecurityUtils.getSubject().getPrincipal();
        String oldPassword = json.getString("oldPassword");
        String password = json.getString("password");

        boolean success = accountService.changePassword(user.getId(), oldPassword, password);
        int code = success ? Constant.RespCode.SUCCESS.getCode() : Constant.RespCode.BIZ_FAIL.getCode();

        ret.put("code", code);
        return ret;
    }

    @PostMapping("/forgotPassword")
    public Object forgotPassword(HttpServletRequest request, @RequestBody JSONObject json) {
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

    @PostMapping("/checkResetPassword")
    public Object checkResetPassword(HttpServletRequest request, @RequestBody JSONObject json) {
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

    @PostMapping("/resetPassword")
    public Object resetPassword(HttpServletRequest request, @RequestBody JSONObject json) {
        Map<String, Object> ret = new HashMap<String, Object>();

        String verifyCode = json.getString("vcode");
        String password = json.getString("password");

        Subject subject = SecurityUtils.getSubject();
        TstUser user = accountService.resetPassword(verifyCode, password);

        if (user != null) {
            ret.put("token", subject.getSession().getId());
            ret.put("orgId", user.getDefaultOrgId());
            ret.put("code", Constant.RespCode.SUCCESS.getCode());
        } else {
            ret.put("code", Constant.RespCode.BIZ_FAIL.getCode());
            ret.put("msg", "重置密码失败");
        }

        return ret;
    }
}
