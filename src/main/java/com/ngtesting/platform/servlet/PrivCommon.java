package com.ngtesting.platform.servlet;

import org.springframework.core.annotation.AliasFor;

import java.lang.annotation.*;

@Target(ElementType.METHOD)
@Retention(RetentionPolicy.RUNTIME)
@Documented
public @interface PrivCommon {

    @AliasFor("check")
    String check() default "true";

}