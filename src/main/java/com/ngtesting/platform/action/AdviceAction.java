package com.ngtesting.platform.action;

import com.ngtesting.platform.exception.AuthException;
import org.springframework.web.bind.annotation.ControllerAdvice;
import org.springframework.web.bind.annotation.ExceptionHandler;
import org.springframework.web.bind.annotation.ResponseBody;

import java.util.Map;

@ControllerAdvice
public class AdviceAction extends BaseAction {

    @ResponseBody
    @ExceptionHandler(value = Exception.class)
    public Map errorHandler(Exception ex) {
        return bizFail();
    }

    @ResponseBody
    @ExceptionHandler(value = AuthException.class)
    public Map authErrorHandler(AuthException ex) {
        return authFail();
    }

}
