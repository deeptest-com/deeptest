package com.ngtesting.platform.service.impl;

import com.alibaba.fastjson.JSONObject;
import com.ngtesting.platform.entity.TestSuite;
import com.ngtesting.platform.service.SuiteService;
import com.ngtesting.platform.config.Constant.TreeNodeType;
import com.ngtesting.platform.vo.TestSuiteTreeVo;
import com.ngtesting.platform.vo.TestSuiteVo;
import org.hibernate.criterion.DetachedCriteria;
import org.hibernate.criterion.Order;
import org.hibernate.criterion.Restrictions;
import org.springframework.stereotype.Service;

import java.util.HashMap;
import java.util.LinkedList;
import java.util.List;
import java.util.Map;

@Service
public class SuiteServiceImpl extends BaseServiceImpl implements SuiteService {

	@Override
	public List<TestSuite> query(Long projectId) {
        DetachedCriteria dc = DetachedCriteria.forClass(TestSuite.class);

        if (projectId != null) {
        	dc.add(Restrictions.eq("projectId", projectId));
        }

        dc.add(Restrictions.eq("deleted", Boolean.FALSE));
        dc.add(Restrictions.eq("disabled", Boolean.FALSE));
        dc.addOrder(Order.asc("path"));
        dc.addOrder(Order.asc("id"));
        List<TestSuite> ls = findAllByCriteria(dc);

        return ls;
	}

	@Override
	public TestSuiteVo getById(Long caseId) {
		TestSuite po = (TestSuite) get(TestSuite.class, caseId);
		TestSuiteVo vo = genVo(po);

		return vo;
	}

	@Override
	public TestSuiteTreeVo buildTree(List<TestSuite> ls) {
		TestSuiteTreeVo root = null;

		Map<Long, TestSuiteTreeVo> nodeMap = new HashMap<Long, TestSuiteTreeVo>();
        for (TestSuite po : ls) {
        	Long id = po.getId();
        	String title = po.getTitle();
        	TreeNodeType type = po.getType();
        	Long pid = po.getParentId();

        	TestSuiteTreeVo newNode = new TestSuiteTreeVo(id, title, type.toString(), pid);
        	nodeMap.put(id, newNode);

        	if (type.equals(TreeNodeType.root)) {
        		root = newNode;
        		continue;
        	}

        	nodeMap.get(pid).getChildren().add(newNode);
        }

        return root;
	}

	@Override
	public List<TestSuiteVo> genVos(List<TestSuite> pos) {
        List<TestSuiteVo> vos = new LinkedList<TestSuiteVo>();

        for (TestSuite po: pos) {
        	TestSuiteVo vo = genVo(po);
        	vos.add(vo);
        }
		return vos;
	}

	@Override
	public TestSuiteVo genVo(TestSuite po) {
		TestSuiteVo vo = new TestSuiteVo(po.getId(), po.getTitle(), po.getPriority(), po.getEstimate(), po.getObjective(),
				po.getDescr(), po.getPath(), po.getType().toString());

		return vo;
	}

	@Override
	public TestSuite create(Long id, String title, String type, Long pid, Long userId) {
		TestSuite parent = (TestSuite) get(TestSuite.class, pid);

		TestSuite testSuite = new TestSuite();
		testSuite.setTitle(title);
		testSuite.setType(TreeNodeType.valueOf(type));
		testSuite.setParentId(pid);
		testSuite.setProjectId(parent.getProjectId());
		testSuite.setUserId(userId);

		testSuite.setPath(parent.getPath() + parent.getId() + "/");

		testSuite.setOrderInParent(getChildMaxOrderNumb(parent) + 1);

		saveOrUpdate(testSuite);

		return testSuite;
	}

	@Override
	public TestSuite rename(Long id, String value, Long id2) {
		// TODO Auto-generated method stub
		return null;
	}

	@Override
	public TestSuite move(Long id, Long pid, Long prePid, Long id2) {


		return null;
	}

	@Override
	public TestSuite save(JSONObject json) {
		TestSuite testSuite = (TestSuite) get(TestSuite.class, json.getLong("id"));

//		long i = 0;
//		while (i < 10000) {
//			create(i, "性能测试", "leaf", Long.valueOf(2), Long.valueOf(-1));
//			i++;
//		}

		return testSuite;
	}

	@Override
	public TestSuite delete(Long vo, Long clientId) {
		// TODO Auto-generated method stub
		return null;
	}

	private Integer getChildMaxOrderNumb(TestSuite parent) {
		String hql = "select max(orderInParent) from TestSuite where parentId = " + parent.getId();
		Integer maxOrder = (Integer) getByHQL(hql);

		if (maxOrder == null) {
			maxOrder = 0;
		}

		return maxOrder;
	}

}

