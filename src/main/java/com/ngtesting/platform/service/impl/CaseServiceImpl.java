package com.ngtesting.platform.service.impl;

import com.alibaba.fastjson.JSON;
import com.alibaba.fastjson.JSONObject;
import com.ngtesting.platform.entity.TestCase;
import com.ngtesting.platform.entity.TestCaseStep;
import com.ngtesting.platform.service.CaseService;
import com.ngtesting.platform.service.CaseStepService;
import com.ngtesting.platform.service.CustomFieldService;
import com.ngtesting.platform.util.BeanUtilEx;
import com.ngtesting.platform.vo.TestCaseStepVo;
import com.ngtesting.platform.vo.TestCaseVo;
import org.hibernate.criterion.DetachedCriteria;
import org.hibernate.criterion.Order;
import org.hibernate.criterion.Restrictions;
import org.springframework.beans.factory.annotation.Autowired;
import org.springframework.stereotype.Service;

import java.util.LinkedList;
import java.util.List;

@Service
public class CaseServiceImpl extends BaseServiceImpl implements CaseService {

	@Autowired
	CustomFieldService customFieldService;

    @Autowired
    CaseStepService caseStepService;

	@Override
	public List<TestCase> query(Long projectId) {
        DetachedCriteria dc = DetachedCriteria.forClass(TestCase.class);

        if (projectId != null) {
        	dc.add(Restrictions.eq("projectId", projectId));
        }
        
        dc.add(Restrictions.eq("deleted", Boolean.FALSE));
        dc.add(Restrictions.eq("disabled", Boolean.FALSE));

		dc.addOrder(Order.asc("pId"));
        dc.addOrder(Order.asc("ordr"));

        List<TestCase> ls = findAllByCriteria(dc);
        
        return ls;
	}

	@Override
	public TestCaseVo getById(Long caseId) {
		TestCase po = (TestCase) get(TestCase.class, caseId);
		TestCaseVo vo = genVo(po);

		return vo;
	}

	@Override
	public TestCase create(Long id, String name, String type, Long pid, Long userId) {
		TestCase parent = (TestCase) get(TestCase.class, pid);
		
		TestCase testCase = new TestCase();
		testCase.setName(name);
		testCase.setpId(pid);
		testCase.setProjectId(parent.getProjectId());
		testCase.setUserId(userId);
		
		testCase.setOrdr(getChildMaxOrderNumb(parent.getId()) + 1);
		
		saveOrUpdate(testCase);
		
		return testCase;
	}

	@Override
	public TestCase rename(JSONObject json, Long userId) {
        TestCase testCase = JSON.parseObject(JSON.toJSONString(json), TestCase.class);
        testCase.setUserId(userId);

        if (testCase.getOrdr() == null) {
            testCase.setOrdr(getChildMaxOrderNumb(testCase.getpId()));
        }

        saveOrUpdate(testCase);
        return testCase;
	}
	
	@Override
	public TestCase movePers(JSONObject json, Long userId) {
        Long srcId = json.getLong("srcId");
        Long targetId = json.getLong("targetId");
        String moveType = json.getString("moveType");
        Boolean isCopy = json.getBoolean("isCopy");

        TestCase src = (TestCase) get(TestCase.class, srcId);;
        TestCase target = (TestCase) get(TestCase.class, targetId);

        TestCase srcCase;
        if (isCopy) {
            srcCase = new TestCase();
            BeanUtilEx.copyProperties(srcCase, src);
            srcCase.setId(null);
        } else {
            srcCase = src;
        }

        if ("inner".equals(moveType)) {
            srcCase.setpId(target.getId());
        } else if ("prev".equals(moveType)) {
            String hql = "update TestCase c set c.ordr = c.ordr+1 where c.ordr >= ?";
            getDao().queryHql(hql, target.getOrdr());
            srcCase.setOrdr(target.getOrdr());
        } else if ("next".equals(moveType)) {
            String hql = "update TestCase c set c.ordr = c.ordr+1 where c.ordr > ?";
            getDao().queryHql(hql, target.getOrdr());
            srcCase.setOrdr(target.getOrdr() + 1);
        }

        saveOrUpdate(srcCase);
		return srcCase;
	}

	@Override
	public TestCase save(JSONObject json, Long userId) {
        TestCaseVo testCaseVo = JSON.parseObject(JSON.toJSONString(json), TestCaseVo.class);

        TestCase testCasePo = new TestCase();
        if (testCaseVo.getId() > 0) {
            testCasePo = (TestCase)get(TestCase.class, testCaseVo.getId());
            copyProperties(testCasePo, testCaseVo);
        } else {
            copyProperties(testCasePo, testCaseVo);
            testCasePo.setId(null);
            testCasePo.setOrdr(getChildMaxOrderNumb(testCasePo.getpId()));
        }

        testCasePo.setUserId(userId);
        saveOrUpdate(testCasePo);

		return testCasePo;
	}

    @Override
	public TestCase saveField(JSONObject json) {
		Long id = json.getLong("id");
		String prop = json.getString("prop");
		String value = json.getString("value");

		TestCase testCase = (TestCase) get(TestCase.class, id);

		if ("title".equals(prop)) {
			testCase.setName(value);
		} else if ("objective".equals(prop)) {
            testCase.setObjective(value);
        } else if ("prop01".equals(prop)) {
            testCase.setProp01(value);
        } else if ("prop02".equals(prop)) {
            testCase.setProp02(value);
        } else if ("prop03".equals(prop)) {
            testCase.setProp03(value);
        } else if ("prop04".equals(prop)) {
            testCase.setProp04(value);
        } else if ("prop05".equals(prop)) {
            testCase.setProp05(value);
        } else if ("prop06".equals(prop)) {
            testCase.setProp06(value);
        } else if ("prop07".equals(prop)) {
            testCase.setProp07(value);
        } else if ("prop08".equals(prop)) {
            testCase.setProp08(value);
        } else if ("prop09".equals(prop)) {
            testCase.setProp09(value);
        } else if ("prop10".equals(prop)) {
            testCase.setProp10(value);
        } else if ("prop11".equals(prop)) {
            testCase.setProp11(value);
        } else if ("prop12".equals(prop)) {
            testCase.setProp12(value);
        } else if ("prop13".equals(prop)) {
            testCase.setProp13(value);
        } else if ("prop14".equals(prop)) {
            testCase.setProp14(value);
        } else if ("prop15".equals(prop)) {
            testCase.setProp15(value);
        } else if ("prop16".equals(prop)) {
            testCase.setProp16(value);
        } else if ("prop17".equals(prop)) {
            testCase.setProp17(value);
        } else if ("prop18".equals(prop)) {
            testCase.setProp18(value);
        } else if ("prop19".equals(prop)) {
            testCase.setProp19(value);
        } else if ("prop20".equals(prop)) {
            testCase.setProp20(value);
        }
		saveOrUpdate(testCase);

		return testCase;
	}

	@Override
	public TestCase saveCustomizedField(JSONObject json) {
		Long id = json.getLong("id");
		String prop = json.getString("prop");
		String value = json.getString("value");

		TestCase testCase = (TestCase) get(TestCase.class, id);
        // TODO:
		saveOrUpdate(testCase);

		return testCase;
	}

	@Override
	public TestCase delete(Long id, Long userId) {
        TestCase testCase = (TestCase) get(TestCase.class, id);
        testCase.setDeleted(true);
        testCase.setUserId(userId);
        saveOrUpdate(testCase);

        return testCase;
	}
	
	private Integer getChildMaxOrderNumb(Long parentId) {
		String hql = "select max(ordr) from TestCase where pId = " + parentId;
		Integer maxOrder = (Integer) getByHQL(hql);
		
		if (maxOrder == null) {
			maxOrder = 0;
		}
        
		return maxOrder + 1;
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
    public List<TestCaseVo> genVos(List<TestCase> pos, boolean withSteps) {
        List<TestCaseVo> vos = new LinkedList<TestCaseVo>();

        for (TestCase po: pos) {
            TestCaseVo vo = genVo(po, withSteps);
            vos.add(vo);
        }
        return vos;
    }

    @Override
    public TestCaseVo genVo(TestCase po) {
	    return genVo(po, true);
    }

    @Override
    public TestCaseVo genVo(TestCase po, boolean withSteps) {
        TestCaseVo vo = new TestCaseVo();

        BeanUtilEx.copyProperties(vo, po);
        vo.setEstimate(po.getEstimate());

        vo.setSteps(new LinkedList<TestCaseStepVo>());

        if (withSteps) {
            List<TestCaseStep> steps = po.getSteps();
            for (TestCaseStep step : steps) {
                TestCaseStepVo stepVo = new TestCaseStepVo(
                        step.getId(), step.getOpt(), step.getExpect(), step.getOrdr(), step.getTestCaseId());

                vo.getSteps().add(stepVo);
            }
        }

//		List<TestCaseProp> props = po.getProps();
//		for (TestCaseProp propPo : props) {
//
//			TestCasePropVo propVo = new TestCasePropVo(propPo.getId(), propPo.getCode(),
//                    propPo.getLabel(), propPo.getValue(), propPo.getFieldId());
//
////			CustomFieldVo fieldVo = new CustomFieldVo();
////			BeanUtilEx.copyProperties(fieldVo, propPo.getField());
////			propVo.setField(fieldVo);
//
//			vo.getProps().add(propVo);
//		}

        return vo;
    }

    @Override
    public void copyProperties(TestCase testCasePo, TestCaseVo testCaseVo) {
        testCasePo.setId(testCaseVo.getId());
        testCasePo.setName(testCaseVo.getName());
        testCasePo.setPriority(testCaseVo.getPriority());
        testCasePo.setType(TestCase.CaseType.valueOf(testCaseVo.getType()));
        testCasePo.setEstimate(testCaseVo.getEstimate());

        testCasePo.setObjective(testCaseVo.getObjective());

        testCasePo.setDescr(testCaseVo.getDescr());
        testCasePo.setOrdr(testCaseVo.getOrdr());

        testCasePo.setpId(testCaseVo.getpId());

        testCasePo.setProp01(testCaseVo.getProp01());
        testCasePo.setProp02(testCaseVo.getProp02());
        testCasePo.setProp03(testCaseVo.getProp03());
        testCasePo.setProp04(testCaseVo.getProp04());
        testCasePo.setProp05(testCaseVo.getProp05());

        testCasePo.setProp06(testCaseVo.getProp06());
        testCasePo.setProp07(testCaseVo.getProp07());
        testCasePo.setProp08(testCaseVo.getProp08());
        testCasePo.setProp09(testCaseVo.getProp09());
        testCasePo.setProp10(testCaseVo.getProp10());

        testCasePo.setProp11(testCaseVo.getProp11());
        testCasePo.setProp12(testCaseVo.getProp12());
        testCasePo.setProp13(testCaseVo.getProp13());
        testCasePo.setProp14(testCaseVo.getProp14());
        testCasePo.setProp15(testCaseVo.getProp15());

        testCasePo.setProp16(testCaseVo.getProp16());
        testCasePo.setProp17(testCaseVo.getProp17());
        testCasePo.setProp18(testCaseVo.getProp18());
        testCasePo.setProp19(testCaseVo.getProp19());
        testCasePo.setProp20(testCaseVo.getProp20());
    }

}

