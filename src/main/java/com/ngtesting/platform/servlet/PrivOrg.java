package com.ngtesting.platform.servlet;

import org.springframework.core.annotation.AliasFor;

import java.lang.annotation.*;

@Target(ElementType.METHOD)
@Retention(RetentionPolicy.RUNTIME)
@Documented
public @interface PrivOrg {

    @AliasFor("perms")
    String[] perms() default {};

    @AliasFor("src")
    String src() default "";

    @AliasFor("key")
    String key() default "orgId";

    @AliasFor("opt")
    String opt() default "or";

}