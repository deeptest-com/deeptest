package com.ngtesting.platform.ctrl;

import com.ngtesting.platform.config.Constant;
import com.ngtesting.platform.model.TstUser;
import com.ngtesting.platform.model.TstVerifyCode;
import com.ngtesting.platform.service.inf.*;
import com.ngtesting.platform.utils.AuthPassport;
import org.springframework.beans.factory.annotation.Autowired;
import org.springframework.stereotype.Controller;
import org.springframework.web.bind.annotation.*;

import java.util.HashMap;
import java.util.Map;

@Controller
@RequestMapping(value = Constant.API_PATH_CLIENT + "/account")
public class AccountCtrl {
    @Autowired
    private AccountService accountService;
    @Autowired
    private OrgService orgService;
    @Autowired
    private UserService userService;

    @Autowired
    private PropService propService;
    @Autowired
    private MailService mailService;

    @AuthPassport(validate=false)
    @ResponseBody
    @PostMapping("/register")
    public Map register(@RequestBody TstUser json){
        Map<String, Object> ret = new HashMap();
        TstUser user = accountService.register(json);

        if (user != null) {
            orgService.createDefaultBasicDataPers(user);

            TstVerifyCode verifyCode = accountService.genVerifyCode(user.getId());
            String sys = propService.getSysName();

            Map<String, String> map = new HashMap<String, String>();
            map.put("name", user.getNickname());
            map.put("vcode", verifyCode.getCode());

            String url = propService.getUrlLogin();
            if (!url.startsWith("http")) {
                url = Constant.WEB_ROOT + url;
            }
            map.put("url", url);
            mailService.sendTemplateMail("[\"" + sys + "\"]注册成功", "register-success.ftl", user.getEmail(), map);
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
    @GetMapping("/login")
    public Object login(TstUser user){
        return accountService.register(user);
    }

    @AuthPassport(validate=false)
    @ResponseBody
    @GetMapping("/loginWithVerifyCode")
    public Object loginWithVerifyCode(TstUser user){
        return accountService.register(user);
    }

    @ResponseBody
    @GetMapping("/logout")
    public Object logout(TstUser user){
        return accountService.register(user);
    }

    @ResponseBody
    @GetMapping("/changePassword")
    public Object changePassword(TstUser user){
        return accountService.register(user);
    }

    @AuthPassport(validate=false)
    @ResponseBody
    @GetMapping("/checkResetPassword")
    public Object checkResetPassword(TstUser user){
        return accountService.register(user);
    }

    @AuthPassport(validate=false)
    @ResponseBody
    @GetMapping("/resetPassword")
    public Object resetPassword(TstUser user){
        return accountService.register(user);
    }

    @AuthPassport(validate=false)
    @ResponseBody
    @GetMapping("/forgotPassword")
    public Object forgotPassword(TstUser user){
        return accountService.register(user);
    }
}
