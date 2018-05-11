package com.ngtesting.platform.service.impl;

import com.alibaba.fastjson.JSON;
import com.alibaba.fastjson.JSONObject;
import com.ngtesting.platform.config.Constant;
import com.ngtesting.platform.entity.*;
import com.ngtesting.platform.service.CaseAttachmentService;
import com.ngtesting.platform.service.CaseCommentsService;
import com.ngtesting.platform.service.CaseService;
import com.ngtesting.platform.util.BeanUtilEx;
import com.ngtesting.platform.util.StringUtil;
import com.ngtesting.platform.vo.*;
import org.apache.commons.lang.StringUtils;
import org.hibernate.criterion.DetachedCriteria;
import org.hibernate.criterion.Order;
import org.hibernate.criterion.Restrictions;
import org.springframework.beans.factory.annotation.Autowired;
import org.springframework.stereotype.Service;

import java.util.Date;
import java.util.Iterator;
import java.util.LinkedList;
import java.util.List;

@Service
public class CaseServiceImpl extends BaseServiceImpl implements CaseService {
    @Autowired
    CaseCommentsService caseCommentsService;
    @Autowired
    CaseAttachmentService caseAttachmentService;

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
    public List<TestCaseVo> queryForSuiteSelection(Long projectId, Long caseProjectId, Long suiteId) {
        Long id = caseProjectId == null? projectId: caseProjectId;

        List<Long> selectIds = new LinkedList<>();
	    if (suiteId != null) {
            TestSuite suite = (TestSuite)get(TestSuite.class, suiteId);
            for (TestCaseInSuite testcase : suite.getTestcases()) {
                selectIds.add(testcase.getCaseId());
            }
        }

        List<TestCase> pos = query(id);
        List<TestCaseVo> vos = genVos(pos, selectIds, false);

        return vos;
    }
    @Override
    public List<TestCaseVo> queryForRunSelection(Long projectId, Long caseProjectId, Long runId) {
	    Long id = caseProjectId == null? projectId: caseProjectId;

        TestRun run = (TestRun)get(TestRun.class, runId);

        List<Long> selectIds = new LinkedList<>();
        for (TestCaseInRun testcase : run.getTestcases()) {
            selectIds.add(testcase.getCaseId());
        }

        List<TestCase> pos = query(id);
        List<TestCaseVo> vos = genVos(pos, selectIds, false);

        return vos;
    }

    @Override
	public TestCaseVo getById(Long caseId) {
		TestCase po = (TestCase) get(TestCase.class, caseId);
		TestCaseVo vo = genVo(po, true);

		return vo;
	}

    @Override
    public TestCase renamePers(JSONObject json, UserVo user) {
        Long id = json.getLong("id");
        String name = json.getString("name");
        Long pId = json.getLong("pId");
        Long projectId = json.getLong("projectId");

        return renamePers(id, name, pId, projectId, user);
    }

	@Override
	public TestCase renamePers(Long id, String name, Long pId, Long projectId, UserVo user) {
        TestCase testCasePo = new TestCase();
        Constant.CaseAct action;
        if (id != null && id > 0) {
            testCasePo = (TestCase)get(TestCase.class, id);

            testCasePo.setUpdateById(user.getId());
            testCasePo.setUpdateTime(new Date());
            action = Constant.CaseAct.rename;
        } else {
            testCasePo.setLeaf(true);
            testCasePo.setId(null);
            testCasePo.setpId(pId);
            testCasePo.setType("functional");
            testCasePo.setPriority("medium");
            testCasePo.setContent("");
            testCasePo.setOrdr(getChildMaxOrderNumb(testCasePo.getpId()));

            testCasePo.setCreateById(user.getId());
            testCasePo.setCreateTime(new Date());
            action = Constant.CaseAct.create;
        }
        testCasePo.setName(name);
        testCasePo.setReviewResult(null);
        testCasePo.setProjectId(projectId);

        saveOrUpdate(testCasePo);

        saveHistory(user, action, testCasePo,null);

        return testCasePo;
	}

	@Override
	public TestCaseVo movePers(JSONObject json, UserVo user) {
        Long srcId = json.getLong("srcId");
        Long targetId = json.getLong("targetId");
        String moveType = json.getString("moveType");
        Boolean isCopy = json.getBoolean("isCopy");

        TestCase src = (TestCase) get(TestCase.class, srcId);;
        TestCase target = (TestCase) get(TestCase.class, targetId);

        TestCase testCase;
        Constant.CaseAct action;
        if (isCopy) {
            testCase = new TestCase();
            BeanUtilEx.copyProperties(testCase, src);

            // 不能用旧的
            testCase.setSteps(new LinkedList());
            testCase.setHistories(new LinkedList());
            testCase.setComments(new LinkedList());
            testCase.setAttachments(new LinkedList());

            testCase.setId(null);
            action = Constant.CaseAct.copy;
        } else {
            testCase = src;
            action = Constant.CaseAct.move;
        }

        if ("inner".equals(moveType)) {
            testCase.setpId(target.getId());
        } else if ("prev".equals(moveType)) {
            String hql = "update TestCase c set c.ordr = c.ordr+1 where c.ordr >= ? and c.pId=? and id!=?";
            getDao().queryHql(hql, target.getOrdr(), target.getpId(), testCase.getId());

            testCase.setpId(target.getpId());
            testCase.setOrdr(target.getOrdr());
        } else if ("next".equals(moveType)) {
            String hql = "update TestCase c set c.ordr = c.ordr+1 where c.ordr > ? and c.pId=? and id!=?";
            getDao().queryHql(hql, target.getOrdr(), target.getpId(), testCase.getId());

            testCase.setpId(target.getpId());
            testCase.setOrdr(target.getOrdr() + 1);
        }

        saveOrUpdate(testCase);
        boolean isParent = false;
        if (isCopy) {
            isParent = cloneStepsAndChildrenPers(testCase, src);
        }

        TestCaseVo caseVo = new TestCaseVo();
        if (isCopy && isParent) {
            loadNodeTree(caseVo, testCase);
        } else {
            caseVo = genVo(testCase);
        }

        saveHistory(user, action, testCase,null);
        return caseVo;
	}

    @Override
    public void loadNodeTree(TestCaseVo vo, TestCase po) {
        BeanUtilEx.copyProperties(vo, po);
        vo.setEstimate(po.getEstimate());

        List<TestCase> children = getChildren(po.getId());
        for (TestCase childPo : children) {
            TestCaseVo childVo = new TestCaseVo();
            vo.getChildren().add(childVo);

            loadNodeTree(childVo, childPo);
        }
    }

    @Override
    public void createRoot(Long projectId, UserVo user) {
        TestCase root = new TestCase();
        root.setName("测试用例");
        root.setType(null);
        root.setpId(null);
        root.setLeaf(false);
        root.setProjectId(projectId);

        root.setCreateById(user.getId());
        root.setCreateTime(new Date());

        root.setOrdr(0);

        saveOrUpdate(root);

        TestCase testCase = new TestCase();
        testCase.setName("新特性");
        testCase.setType("functional");
        testCase.setPriority("medium");
        testCase.setpId(root.getId());
        testCase.setProjectId(projectId);
        testCase.setCreateById(user.getId());
        testCase.setCreateTime(new Date());
        testCase.setLeaf(false);
        testCase.setOrdr(0);
        saveOrUpdate(testCase);
        saveHistory(user, Constant.CaseAct.create, testCase,null);

        TestCase testCase2 = new TestCase();
        testCase2.setName("新用例");
        testCase2.setType("functional");
        testCase2.setPriority("medium");
        testCase2.setpId(testCase.getId());
        testCase2.setProjectId(projectId);
        testCase2.setCreateById(user.getId());
        testCase2.setCreateTime(new Date());
        testCase2.setLeaf(true);
        testCase2.setOrdr(0);
        saveOrUpdate(testCase2);
        saveHistory(user, Constant.CaseAct.create, testCase2,null);
    }

    @Override
	public TestCase save(JSONObject json, UserVo user) {
        TestCaseVo testCaseVo = JSON.parseObject(JSON.toJSONString(json), TestCaseVo.class);

        Constant.CaseAct action;

        TestCase testCasePo;
        if (testCaseVo.getId() > 0) {
            testCasePo = (TestCase)get(TestCase.class, testCaseVo.getId());
            copyProperties(testCasePo, testCaseVo);

            testCasePo.setUpdateById(user.getId());
            testCasePo.setUpdateTime(new Date());

            action = Constant.CaseAct.update;
        } else {
            testCasePo = new TestCase();
            copyProperties(testCasePo, testCaseVo);
            testCasePo.setId(null);
            testCasePo.setLeaf(true);
            testCasePo.setOrdr(getChildMaxOrderNumb(testCasePo.getpId()));

            testCasePo.setCreateById(user.getId());
            testCasePo.setCreateTime(new Date());

            action = Constant.CaseAct.create;
        }

        testCasePo.setReviewResult(null);
        saveOrUpdate(testCasePo);

        saveHistory(user, action, testCasePo,null);

		return testCasePo;
	}

    @Override
	public TestCase saveField(JSONObject json, UserVo user) {
		Long id = json.getLong("id");
		String prop = json.getString("prop");
		String value = json.getString("value");
		String label = json.getString("label");

		TestCase testCase = (TestCase) get(TestCase.class, id);

		if ("name".equals(prop)) {
			testCase.setName(value);
		} else if ("objective".equals(prop)) {
            testCase.setObjective(value);
        } else if ("content".equals(prop)) {
            testCase.setContent(value);
        } else if ("priority".equals(prop)) {
            testCase.setPriority(value);
        } else if ("estimate".equals(prop)) {
            testCase.setEstimate(Integer.valueOf(value));
        } else if ("type".equals(prop)) {
            testCase.setType(value);
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
        testCase.setReviewResult(null);
		saveOrUpdate(testCase);

        saveHistory(user, Constant.CaseAct.update, testCase,label);

		return testCase;
	}

    @Override
    public void saveHistory(UserVo user, Constant.CaseAct act, TestCase testCase, String field) {
	    String action = act.msg;

        String msg = "用户" + StringUtil.highlightDict(user.getName()) + action;
        if (StringUtils.isNotEmpty(field)) {
            msg += " " + field;
        } else {
//            msg += "信息";
        }
        TestCaseHistory his = new TestCaseHistory();
        his.setTitle(msg);
        his.setTestCaseId(testCase.getId());
        saveOrUpdate(his);
    }

	@Override
	public TestCase delete(Long id, UserVo user) {
        TestCase testCase = (TestCase) get(TestCase.class, id);

        getDao().querySql("{call delete_case_and_its_children(?)}", id);
        saveHistory(user, Constant.CaseAct.delete, testCase,null);

        return testCase;
	}

    @Override
    public void updateParentIfNeededPers(Long pid) {
        getDao().querySql("{call update_case_parent_if_needed(?)}", pid);
    }

    @Override
    public boolean cloneStepsAndChildrenPers(TestCase testCase, TestCase src) {
	    boolean isParent = false;

        for (TestCaseStep step : src.getSteps()) {
            TestCaseStep step1 = new TestCaseStep(testCase.getId(), step.getOpt(), step.getExpect(), step.getOrdr());
            saveOrUpdate(step1);
            testCase.getSteps().add(step1);
        }

        List<TestCase> children = getChildren(src.getId());
        for(TestCase child : children) {
            TestCase clonedChild = new TestCase();
            BeanUtilEx.copyProperties(clonedChild, child);
            // 不能用以前的
            clonedChild.setComments(new LinkedList());
            clonedChild.setSteps(new LinkedList());
            clonedChild.setHistories(new LinkedList());
            clonedChild.setAttachments(new LinkedList());

            clonedChild.setId(null);
            clonedChild.setpId(testCase.getId());

            saveOrUpdate(clonedChild);
            cloneStepsAndChildrenPers(clonedChild, child);
        }

        return children.size() > 0;
    }

    @Override
    public List<TestCase> getChildren(Long caseId) {
        DetachedCriteria dc = DetachedCriteria.forClass(TestCase.class);
        dc.add(Restrictions.eq("pId", caseId));

        dc.add(Restrictions.eq("deleted", Boolean.FALSE));
        dc.add(Restrictions.eq("disabled", Boolean.FALSE));

        dc.addOrder(Order.asc("pId"));
        dc.addOrder(Order.asc("ordr"));

        List<TestCase> children = findAllByCriteria(dc);
        return children;
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
    public List<TestCaseVo> genVos(List<TestCase> pos) { return genVos(pos, false); }
    @Override
    public List<TestCaseVo> genVos(List<TestCase> pos, boolean withSteps) { return genVos(pos, null,false); }
    @Override
    public List<TestCaseVo> genVos(List<TestCase> pos, List<Long> selectIds, boolean withSteps) {
        List<TestCaseVo> vos = new LinkedList<TestCaseVo>();

        for (TestCase po: pos) {
            TestCaseVo vo = genVo(po, selectIds, withSteps);
            vos.add(vo);
        }
        return vos;
    }

    @Override
    public TestCaseVo genVo(TestCase po) {
        return genVo(po, false);
    }

    @Override
    public TestCaseVo genVo(TestCase po, boolean withSteps) { return genVo(po, null, withSteps);}

    @Override
    public TestCaseVo genVo(TestCase po, List<Long> selectIds, boolean withSteps) {
        TestCaseVo vo = new TestCaseVo();

        BeanUtilEx.copyProperties(vo, po);
        vo.setEstimate(po.getEstimate());

        if (selectIds != null && selectIds.contains(po.getId())) {
            vo.setChecked(true);
        }

        vo.setSteps(new LinkedList<TestCaseStepVo>());
        vo.setComments(new LinkedList<TestCaseCommentsVo>());
        vo.setHistories(new LinkedList<TestCaseHistoryVo>());
        vo.setAttachments(new LinkedList<TestCaseAttachmentVo>());

        if (withSteps) {
            List<TestCaseStep> steps = po.getSteps();
            for (TestCaseStep step : steps) {
                TestCaseStepVo stepVo = new TestCaseStepVo(
                        step.getId(), step.getOpt(), step.getExpect(), step.getOrdr(), step.getTestCaseId());

                vo.getSteps().add(stepVo);
            }

            List<TestCaseComments> comments = po.getComments();
            Iterator<TestCaseComments> iterator  = comments.iterator();
            while (iterator.hasNext()) {
                TestCaseComments comment = iterator.next();
                TestCaseCommentsVo commentVo = caseCommentsService.genVo(comment);
                vo.getComments().add(commentVo);
            }

            // 用例历史
            List<TestCaseHistory> histories = findHistories(po.getId());
            for (TestCaseHistory his : histories) {
                TestCaseHistoryVo historyVo = new TestCaseHistoryVo(
                        his.getId(), his.getTitle(), his.getDescr(), his.getTestCaseId(), his.getCreateTime());

                vo.getHistories().add(historyVo);
            }

            List<TestCaseAttachment> attachments = po.getAttachments();
            Iterator<TestCaseAttachment> iteratorAttach  = attachments.iterator();
            while (iteratorAttach.hasNext()) {
                TestCaseAttachment attachment = iteratorAttach.next();
                TestCaseAttachmentVo attachVo = caseAttachmentService.genVo(attachment);
                vo.getAttachments().add(attachVo);
            }
        }

        return vo;
    }

    @Override
    public List<TestCaseHistory> findHistories(Long testCaseId) {
        DetachedCriteria dc = DetachedCriteria.forClass(TestCaseHistory.class);
        dc.add(Restrictions.eq("testCaseId", testCaseId));

        dc.add(Restrictions.eq("deleted", Boolean.FALSE));
        dc.add(Restrictions.eq("disabled", Boolean.FALSE));

        dc.addOrder(Order.desc("createTime"));

        List<TestCaseHistory> ls = findAllByCriteria(dc);
        return ls;
    }

    @Override
    public void copyProperties(TestCase testCasePo, TestCaseVo testCaseVo) {
        testCasePo.setId(testCaseVo.getId());
        testCasePo.setName(testCaseVo.getName());
        testCasePo.setPriority(testCaseVo.getPriority());
        testCasePo.setType(testCaseVo.getType());
        testCasePo.setEstimate(testCaseVo.getEstimate());

        testCasePo.setContentType(testCaseVo.getContentType());
        testCasePo.setContent(testCaseVo.getContent());
        testCasePo.setObjective(testCaseVo.getObjective());

        testCasePo.setOrdr(testCaseVo.getOrdr());

        testCasePo.setpId(testCaseVo.getpId());
        testCasePo.setProjectId(testCaseVo.getProjectId());

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

    @Override
    public TestCase changeContentTypePers(Long id, String contentType) {
        TestCase testCase = (TestCase)get(TestCase.class, id);
        testCase.setContentType(contentType);
        testCase.setReviewResult(null);
        saveOrUpdate(testCase);

        return testCase;
    }

    @Override
    public TestCase reviewPassPers(Long id, Boolean pass) {
        TestCase testCase = (TestCase)get(TestCase.class, id);
        testCase.setReviewResult(pass);
        saveOrUpdate(testCase);

        return testCase;
    }

}

