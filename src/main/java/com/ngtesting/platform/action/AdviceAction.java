package com.ngtesting.platform.action;

import org.apache.shiro.authz.UnauthorizedException;
import org.springframework.web.bind.annotation.ControllerAdvice;
import org.springframework.web.bind.annotation.ExceptionHandler;
import org.springframework.web.bind.annotation.RestController;

import java.util.Map;

@ControllerAdvice
@RestController
public class AdviceAction extends BaseAction {

    @ExceptionHandler(value = UnauthorizedException.class)
    public Map authorErrorHandler(UnauthorizedException ex) {
//        ex.printStackTrace();
        return authorFail();
    }

    @ExceptionHandler(value = Exception.class)
    public Map errorHandler(Exception ex) {
        ex.printStackTrace();
        return bizFail();
    }

}
