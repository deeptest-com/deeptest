package com.ngtesting.platform.contrl;

import com.ngtesting.platform.model.TstUser;
import com.ngtesting.platform.service.intf.AccountService;
import com.ngtesting.platform.service.intf.UserService;
import org.springframework.beans.factory.annotation.Autowired;
import org.springframework.stereotype.Controller;
import org.springframework.web.bind.annotation.GetMapping;
import org.springframework.web.bind.annotation.PostMapping;
import org.springframework.web.bind.annotation.RequestMapping;
import org.springframework.web.bind.annotation.ResponseBody;

@Controller
@RequestMapping(value = "/account")
public class AccountController {
    @Autowired
    private AccountService accountService;
    @Autowired
    private UserService userService;

    @ResponseBody
    @PostMapping("/register")
    public TstUser register(TstUser user){
        TstUser po = accountService.register(user);
        return po;
    }

    @ResponseBody
    @GetMapping("/login")
    public Object login(TstUser user){
        return accountService.register(user);
    }

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

    @ResponseBody
    @GetMapping("/checkResetPassword")
    public Object checkResetPassword(TstUser user){
        return accountService.register(user);
    }

    @ResponseBody
    @GetMapping("/resetPassword")
    public Object resetPassword(TstUser user){
        return accountService.register(user);
    }

    @ResponseBody
    @GetMapping("/forgotPassword")
    public Object forgotPassword(TstUser user){
        return accountService.register(user);
    }
}
