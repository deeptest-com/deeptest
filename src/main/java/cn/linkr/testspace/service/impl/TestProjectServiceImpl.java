package cn.linkr.testspace.service.impl;

import java.util.Date;
import java.util.HashMap;
import java.util.HashSet;
import java.util.Iterator;
import java.util.LinkedHashSet;
import java.util.LinkedList;
import java.util.List;
import java.util.Map;

import org.apache.commons.logging.Log;
import org.apache.commons.logging.LogFactory;
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
public class TestProjectServiceImpl extends BaseServiceImpl implements
		TestProjectService {

	@Autowired
	private ProjectDao projectDao;

	@Override
	public List<TestProject> list(String isActive, String keywords, Long companyId) {
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
	public TestProjectVo genVos(List<TestProject> pos, Map<String, Integer> ret) {
		TestProjectVo root = null;
		int maxLevel = 0;
		
		Map<Long, TestProjectVo> nodeMap = new HashMap<Long, TestProjectVo>();
        for (TestProject po : pos) {
        	Long id = po.getId();
        	Long pid = po.getParentId();
        	
        	TestProjectVo newNode = genVo(po);
        	
        	nodeMap.put(id, newNode);
        	
        	if (id == 0) {
        		root = newNode;
        		continue;
        	}
        	
        	LinkedList<TestProjectVo> children = nodeMap.get(pid).getChildren();
        	children.add(newNode);
        	
        	if (po.getLevel() > maxLevel) {
				maxLevel = po.getLevel();
			}
        }
        
        for (Long key : nodeMap.keySet()) {
        	TestProjectVo vo = nodeMap.get(key);
        	
			if (vo.getId() != 0 && vo.getChildren().size() > 0) {
				TestProjectVo firstChild = vo.getChildren().get(0);
				
				firstChild.setIsFirstChild(true);
				
				int count = 0;
				count = this.countDescendantsNumb(vo, count);
				firstChild.setParentDescendantNumber(count);
				firstChild.setBrotherNumb(count);
			}
        }
		
        ret.put("maxLevel", maxLevel);
        return root;
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
	public int countDescendantsNumb(TestProjectVo vo, int count) {
		Iterator<TestProjectVo> it = vo.getChildren().iterator();
		while (it.hasNext()) {
			TestProjectVo child = it.next();
			count++;

			count = this.countDescendantsNumb(child, count);
		}
		return count;
	}
}
