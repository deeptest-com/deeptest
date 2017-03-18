package cn.linkr.testspace.service.impl;

import java.util.HashMap;
import java.util.HashSet;
import java.util.Iterator;
import java.util.LinkedHashSet;
import java.util.LinkedList;
import java.util.List;
import java.util.Map;

import org.hibernate.criterion.DetachedCriteria;
import org.hibernate.criterion.Order;
import org.hibernate.criterion.Restrictions;
import org.springframework.beans.factory.annotation.Autowired;
import org.springframework.stereotype.Service;

import com.alibaba.fastjson.JSONObject;

import cn.linkr.testspace.dao.BaseDao;
import cn.linkr.testspace.dao.ProjectDao;
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
	
	   @Autowired
	    private ProjectDao projectDao;
	

	@Override
	public List list(String isActive, String keywords, Long companyId) {
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
        List ls = findAllByCriteria(dc);
        
        return ls;
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

	// need to be cache
	@Override
	public HashSet<TestProjectVo> genVos(List<TestProject> pos, Map<String, Integer> ret) {
		int maxLevel = 0;
		
        TestProjectVo root = new TestProjectVo();
        Map<Long, TestProjectVo> voMap = new HashMap<Long, TestProjectVo>();
        for (TestProject po: pos) {
        	if (po.getId() < 1) {
        		continue;
        	}
        	if (po.getLevel() > maxLevel) {
        		maxLevel = po.getLevel();
        	}
        	
        	Long id = po.getId();
        	Long pid = po.getParentId();
        	
        	TestProjectVo vo = genVo(po);
        	voMap.put(id, vo);
        	
        	if (voMap.get(pid) != null) {
        		voMap.get(pid).getChildren().add(vo);
        	} else {
        		root.getChildren().add(vo);
        	}
        	
        }
        
        HashSet<TestProjectVo> out = new LinkedHashSet<TestProjectVo>();
        this.toList(root, out);
        for (TestProjectVo vo: out) {
        	vo.setChildren(null);
        }
        ret.put("maxLevel", maxLevel);
		return out;
	}
	
	@Override
	public TestProjectVo genVo(TestProject po) {
		TestProjectVo vo = new TestProjectVo();
		BeanUtilEx.copyProperties(vo, po);
		if (po.getParentId() == null) {
			vo.setParentId(null);
		}
		return vo;
	}
	
	@Override
	public void toList(TestProjectVo root, HashSet<TestProjectVo> out) {
		Iterator<TestProjectVo> it = root.getChildren().iterator();
		while (it.hasNext()) {
			TestProjectVo vo = it.next();
			int count = 0;
			count = this.countChildrenNumb(vo, count);
			vo.setChildrenNumb(count);
			
			out.add(vo);
			this.toList(vo, out);
		}
	}
	
	@Override
	public int countChildrenNumb(TestProjectVo vo, int count) {
		Iterator<TestProjectVo> it = vo.getChildren().iterator();
		while (it.hasNext()) {
			TestProjectVo child = it.next();
			count++;
			
			count = this.countChildrenNumb(child, count);
		}
		return count;
	}
}
