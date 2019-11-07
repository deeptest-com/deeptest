package com.ngtesting.platform.test;

import org.testng.Assert;
import org.testng.annotations.Test;

public class Demo {
    @Test
    public void f() {
        System.out.println("This is a test.");
        Assert.assertTrue(true);
    }
}
