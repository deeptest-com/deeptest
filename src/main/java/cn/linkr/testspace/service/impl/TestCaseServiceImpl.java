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
import cn.linkr.testspace.service.GuestService;
import cn.linkr.testspace.service.TestCaseService;
import cn.linkr.testspace.util.BeanUtilEx;
import cn.linkr.testspace.util.StringUtil;
import cn.linkr.testspace.vo.GuestVo;
import cn.linkr.testspace.vo.Page;
import cn.linkr.testspace.vo.SessionVo;
import cn.linkr.testspace.vo.TestCaseTreeVo;
import cn.linkr.testspace.vo.TestCaseVo;

@Service
public class TestCaseServiceImpl extends BaseServiceImpl implements TestCaseService {

	@Override
	public List<TestCase> query(Long projectId, Long moduleId, String keywords) {
        DetachedCriteria dc = DetachedCriteria.forClass(TestCase.class);
        
        if (projectId != null) {
        	dc.add(Restrictions.eq("projectId", projectId));
        }
        if (moduleId != null) {
        	dc.add(Restrictions.eq("moduleId", moduleId));
        }
        if (StringUtil.isNotEmpty(keywords)) {
        	dc.add(Restrictions.like("title", "%" + keywords + "%"));
        }
        
        dc.add(Restrictions.eq("deleted", Boolean.FALSE));
        dc.add(Restrictions.eq("disabled", Boolean.FALSE));
        dc.addOrder(Order.asc("id"));
        List<TestCase> ls = findAllByCriteria(dc);
        
        return ls;
	}
	
	@Override
	public TestCaseTreeVo buildTree(List<TestCase> ls) {
		TestCaseTreeVo root = null;
		int i = 0;
		Map<Long, TestCaseTreeVo> nodeMap = new HashMap<Long, TestCaseTreeVo>();
        for (TestCase po : ls) {
        	Long id = po.getId();
        	String title = po.getTitle();
        	Integer type = po.getType();
        	Long pid = po.getPid();
        	
        	TestCaseTreeVo newNode = new TestCaseTreeVo(id, title, type, pid);
        	nodeMap.put(id, newNode);
        	
        	if (type == 0) {
        		root = newNode;
        		continue;
        	}
        	
        	nodeMap.get(pid).getChildren().add(newNode);
        }
		
        return root;
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

	@Override
	public TestCase delete(Long vo, Long clientId) {
		// TODO Auto-generated method stub
		return null;
	}

	@Override
	public TestCase move(Long id, Long pid, Long prePid, Long id2) {
		// TODO Auto-generated method stub
		return null;
	}

	@Override
	public TestCase create(Long id, String value, Integer type, Long pid,
			Long id2) {
		// TODO Auto-generated method stub
		return null;
	}

	@Override
	public TestCase rename(Long id, String value, Long id2) {
		// TODO Auto-generated method stub
		return null;
	}

}

