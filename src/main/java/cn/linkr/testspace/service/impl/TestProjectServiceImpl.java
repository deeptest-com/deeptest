package cn.linkr.testspace.service.impl;

import java.util.HashMap;
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
import cn.linkr.testspace.entity.TestProject;
import cn.linkr.testspace.entity.EvtEvent.EventStatus;
import cn.linkr.testspace.service.GuestService;
import cn.linkr.testspace.service.TestCaseService;
import cn.linkr.testspace.service.TestProjectService;
import cn.linkr.testspace.util.BeanUtilEx;
import cn.linkr.testspace.util.StringUtil;
import cn.linkr.testspace.vo.GuestVo;
import cn.linkr.testspace.vo.Page;
import cn.linkr.testspace.vo.SessionVo;
import cn.linkr.testspace.vo.TestCaseTreeVo;
import cn.linkr.testspace.vo.TestCaseVo;
import cn.linkr.testspace.vo.TestProjectVo;

@Service
public class TestProjectServiceImpl extends BaseServiceImpl implements TestProjectService {

	@Override
	public Page list(String isActive, String keywords, int currentPage, int itemsPerPage) {
        DetachedCriteria dc = DetachedCriteria.forClass(TestProject.class);
        
        if (isActive != null) {
			dc.add(Restrictions.eq("isActive", Boolean.valueOf(isActive)));
		}
        if (StringUtil.isNotEmpty(keywords)) {
        	dc.add(Restrictions.like("name", "%" + keywords + "%"));
        }
        
        dc.add(Restrictions.eq("deleted", Boolean.FALSE));
        dc.add(Restrictions.eq("disabled", Boolean.FALSE));
        dc.addOrder(Order.asc("id"));
        Page page = findPage(dc, currentPage * itemsPerPage, itemsPerPage);
        
        return page;
	}
	
	@Override
	public TestProject delete(Long vo, Long clientId) {
		
		return null;
	}

	@Override
	public TestProject save(Long id, String value, Integer type, Long pid,
			Long id2) {
		
		return null;
	}

	@Override
	public List<TestProjectVo> genVos(List<TestProject> pos) {
        List<TestProjectVo> vos = new LinkedList<TestProjectVo>();

        for (TestProject po: pos) {
        	TestProjectVo vo = genVo(po);
        	vos.add(vo);
        }
		return vos;
	}
	
	@Override
	public TestProjectVo genVo(TestProject po) {
		TestProjectVo vo = new TestProjectVo();
		BeanUtilEx.copyProperties(vo, po);
		return vo;
	}

}

