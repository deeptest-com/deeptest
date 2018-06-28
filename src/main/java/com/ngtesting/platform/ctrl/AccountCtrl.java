package com.ngtesting.platform.ctrl;

import com.ngtesting.platform.config.Constant;
import com.ngtesting.platform.model.TstUser;
import com.ngtesting.platform.service.intf.AccountService;
import com.ngtesting.platform.service.intf.UserService;
import com.ngtesting.platform.utils.AuthPassport;
import org.springframework.beans.factory.annotation.Autowired;
import org.springframework.stereotype.Controller;
import org.springframework.web.bind.annotation.*;

@Controller
@RequestMapping(value = Constant.API_PATH_CLIENT + "/account")
public class AccountCtrl {
    @Autowired
    private AccountService accountService;
    @Autowired
    private UserService userService;

    @AuthPassport(validate=false)
    @ResponseBody
    @PostMapping("/register")
    public TstUser register(@RequestBody TstUser user){
        TstUser po = accountService.register(user);
        return po;
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
