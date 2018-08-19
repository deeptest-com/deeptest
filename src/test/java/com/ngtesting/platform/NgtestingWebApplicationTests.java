package com.ngtesting.platform;

import com.ngtesting.platform.utils.FileUtil;
import org.junit.Test;
import org.junit.runner.RunWith;
import org.springframework.context.annotation.ComponentScan;
import org.springframework.test.context.junit4.SpringRunner;

@RunWith(SpringRunner.class)
@ComponentScan(basePackages = "me.ngtesting.plaform")
public class NgtestingWebApplicationTests {

	@Test
	public void contextLoads() {
		FileUtil.CreateDirIfNeeded("/work/ngtesting-data/");
	}

}
