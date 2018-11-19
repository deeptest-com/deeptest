package com.ngtesting.platform;

import com.alibaba.fastjson.JSON;
import com.ngtesting.platform.config.AppLaunch;
import com.ngtesting.platform.dao.CaseDao;
import com.ngtesting.platform.model.TstCase;
import com.ngtesting.platform.service.CaseExportService;
import com.ngtesting.platform.utils.FileUtil;
import org.junit.Test;
import org.junit.runner.RunWith;
import org.springframework.boot.autoconfigure.EnableAutoConfiguration;
import org.springframework.boot.test.context.SpringBootTest;
import org.springframework.context.annotation.ComponentScan;
import org.springframework.test.context.junit4.SpringRunner;

import javax.annotation.Resource;
import java.util.List;

@RunWith(SpringRunner.class)
@ComponentScan(basePackages = "me.ngtesting.plaform")
@SpringBootTest(classes = AppLaunch.class)
@EnableAutoConfiguration
public class NgtestingWebApplicationTests {

	@Test
	public void contextLoads() {
		FileUtil.CreateDirIfNeeded("/work/ngtesting-data/");
	}

	@Resource
	private CaseExportService caseExportService;

	@Resource
	private CaseDao caseDao;

	@Test
	public void test() {
//		List<TstCase> caseList=caseDao.queryCaseStepInfoByPrj(360);

//		List<TstCase> caseList=caseService.query(360);
//		System.out.println(JSON.toJSONString(caseList));
		List<TstCase> caseList2=caseDao.queryCaseWithStepInfoByPrj(461);
		System.out.println(JSON.toJSONString(caseList2));
	}

	@Test
	public void test2() {
		caseExportService.export(360);
	}

}
