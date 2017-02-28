package cn.linkr.testspace.service.impl;

import java.util.LinkedList;
import java.util.List;
import java.util.Map;

import org.hibernate.criterion.DetachedCriteria;
import org.hibernate.criterion.Order;
import org.hibernate.criterion.Restrictions;
import org.springframework.stereotype.Service;

import com.alibaba.fastjson.JSONObject;

import cn.linkr.testspace.entity.EvtGuest;
import cn.linkr.testspace.entity.EvtScheduleItem;
import cn.linkr.testspace.entity.EvtSession;
import cn.linkr.testspace.entity.TestCase;
import cn.linkr.testspace.service.GuestService;
import cn.linkr.testspace.service.TestCaseService;
import cn.linkr.testspace.util.BeanUtilEx;
import cn.linkr.testspace.vo.GuestVo;
import cn.linkr.testspace.vo.Page;
import cn.linkr.testspace.vo.SessionVo;
import cn.linkr.testspace.vo.TestCaseVo;

@Service
public class TestCaseServiceImpl extends BaseServiceImpl implements TestCaseService {

	@Override
	public List<TestCaseVo> query(String keywords, String status) {
        List ls = getDao().findObjectByProcedure("query_test_case", TestCaseVo.class, null);
        
        return ls;
	}

	@Override
	public List<TestCaseVo> genVos(List<TestCase> pos) {
        List<TestCaseVo> vos = new LinkedList<TestCaseVo>();

        for (TestCase po: pos) {
        	TestCaseVo vo = genVo(po);
        	vos.add(vo);
        }
		return vos;
	}

	@Override
	public TestCaseVo genVo(TestCase po) {
		TestCaseVo vo = new TestCaseVo();
		BeanUtilEx.copyProperties(vo, po);
		return vo;
	}

}

