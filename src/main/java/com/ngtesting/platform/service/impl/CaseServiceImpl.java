package com.ngtesting.platform.service.impl;

import com.alibaba.fastjson.JSONObject;
import com.ngtesting.platform.config.Constant;
import com.ngtesting.platform.model.TstCase;
import com.ngtesting.platform.model.TstCaseHistory;
import com.ngtesting.platform.model.TstUser;
import com.ngtesting.platform.service.CaseAttachmentService;
import com.ngtesting.platform.service.CaseCommentsService;
import com.ngtesting.platform.service.CaseService;
import org.springframework.beans.factory.annotation.Autowired;
import org.springframework.stereotype.Service;

import java.util.LinkedList;
import java.util.List;

@Service
public class CaseServiceImpl extends BaseServiceImpl implements CaseService {
    @Autowired
    CaseCommentsService caseCommentsService;
    @Autowired
    CaseAttachmentService caseAttachmentService;

	@Override
	public List<TstCase> query(Integer projectId) {
//        DetachedCriteria dc = DetachedCriteria.forClass(TstCase.class);
//
//        if (projectId != null) {
//        	dc.add(Restrictions.eq("projectId", projectId));
//        }
//
//        dc.add(Restrictions.eq("deleted", Boolean.FALSE));
//        dc.add(Restrictions.eq("disabled", Boolean.FALSE));
//
//		dc.addOrder(Order.asc("pId"));
//        dc.addOrder(Order.asc("ordr"));
//
//        List<TstCase> ls = findAllByCriteria(dc);
//
//        return ls;

        return null;
	}

    @Override
    public List<TstCase> queryForSuiteSelection(Integer projectId, Integer caseProjectId, Integer suiteId) {
//        Integer id = caseProjectId == null? projectId: caseProjectId;
//
//        List<Integer> selectIds = new LinkedList<>();
//	    if (suiteId != null) {
//            TstSuite suite = (TstSuite)get(TstSuite.class, suiteId);
//            for (TstCaseInSuite testcase : suite.getTestcases()) {
//                selectIds.add(testcase.getCaseId());
//            }
//        }
//
//        List<TstCase> pos = query(id);
//        List<TstCase> vos = genVos(pos, selectIds, false);
//
//        return vos;

        return null;
    }
    @Override
    public List<TstCase> queryForRunSelection(Integer projectId, Integer caseProjectId, Integer runId) {
//        Integer id = caseProjectId == null? projectId: caseProjectId;
//
//        TstTask run = (TstTask)get(TstTask.class, runId);
//
//        List<Long> selectIds = new LinkedList<>();
//        for (TstCaseInRun testcase : run.getTestcases()) {
//            selectIds.add(testcase.getCaseId());
//        }
//
//        List<TstCase> pos = query(id);
//        List<TstCase> vos = genVos(pos, selectIds, false);
//
//        return vos;

        return null;
    }

    @Override
	public TstCase getById(Integer caseId) {
//		TstCase po = (TstCase) get(TstCase.class, caseId);
//		TstCase vo = genVo(po, true);
//
//		return vo;

        return null;
	}

    @Override
    public TstCase renamePers(JSONObject json, TstUser user) {
        Integer id = json.getInteger("id");
        String name = json.getString("name");
        Integer pId = json.getInteger("pId");
        Integer projectId = json.getInteger("projectId");

        return renamePers(id, name, pId, projectId, user);
    }

	@Override
	public TstCase renamePers(Integer id, String name, Integer pId, Integer projectId, TstUser user) {
        TstCase testCasePo = new TstCase();
//        Constant.CaseAct action;
//        if (id != null && id > 0) {
//            testCasePo = (TstCase)get(TstCase.class, id);
//
//            testCasePo.setUpdateById(user.getId());
//            testCasePo.setUpdateTime(new Date());
//            action = Constant.CaseAct.rename;
//        } else {
//            testCasePo.setLeaf(true);
//            testCasePo.setId(null);
//            testCasePo.setpId(pId);
//            testCasePo.setType("functional");
//            testCasePo.setPriority("medium");
//            testCasePo.setContent("");
//            testCasePo.setOrdr(getChildMaxOrderNumb(testCasePo.getpId()));
//
//            testCasePo.setCreateById(user.getId());
//            testCasePo.setCreateTime(new Date());
//            action = Constant.CaseAct.create;
//        }
//        testCasePo.setName(name);
//        testCasePo.setReviewResult(null);
//        testCasePo.setProjectId(projectId);
//
//        saveOrUpdate(testCasePo);
//
//        updateParentIfNeededPers(testCasePo.getpId());
//
//        saveHistory(user, action, testCasePo,null);

        return testCasePo;
	}

	@Override
	public TstCase movePers(JSONObject json, TstUser user) {
//        Integer srcId = json.getInteger("srcId");
//
//        Integer targetId = json.getInteger("targetId");
//        String moveType = json.getString("moveType");
//        Boolean isCopy = json.getBoolean("isCopy");
//
//        TstCase src = (TstCase) get(TstCase.class, srcId);;
//        TstCase target = (TstCase) get(TstCase.class, targetId);
//
//        Long parentId = src.getpId();
//
//        TstCase testCase;
//        Constant.CaseAct action;
//        if (isCopy) {
//            testCase = new TstCase();
//            BeanUtilEx.copyProperties(testCase, src);
//            testCase.setCreateTime(new Date());
//            testCase.setUpdateTime(null);
//
//            // 不能用旧的
//            testCase.setSteps(new LinkedList());
//            testCase.setHistories(new LinkedList());
//            testCase.setComments(new LinkedList());
//            testCase.setAttachments(new LinkedList());
//
//            testCase.setId(null);
//            action = Constant.CaseAct.copy;
//        } else {
//            testCase = src;
//            action = Constant.CaseAct.move;
//        }
//
//        if ("inner".equals(moveType)) {
//            testCase.setpId(target.getId());
//        } else if ("prev".equals(moveType)) {
//            String hql = "update TstCase c set c.ordr = c.ordr+1 where c.ordr >= ? and c.pId=? and id!=?";
//            getDao().queryHql(hql, target.getOrdr(), target.getpId(), testCase.getId());
//
//            testCase.setpId(target.getpId());
//            testCase.setOrdr(target.getOrdr());
//        } else if ("next".equals(moveType)) {
//            String hql = "update TstCase c set c.ordr = c.ordr+1 where c.ordr > ? and c.pId=? and id!=?";
//            getDao().queryHql(hql, target.getOrdr(), target.getpId(), testCase.getId());
//
//            testCase.setpId(target.getpId());
//            testCase.setOrdr(target.getOrdr() + 1);
//        }
//
//        saveOrUpdate(testCase);
//        boolean isParent = false;
//        if (isCopy) {
//            isParent = cloneStepsAndChildrenPers(testCase, src);
//        }
//
//        getDao().flush();
//        updateParentIfNeededPers(parentId);
//        updateParentIfNeededPers(targetId);
//
//        TstCase caseVo = new TstCase();
//        if (isCopy && isParent) {
//            loadNodeTree(caseVo, testCase);
//        } else {
//            caseVo = genVo(testCase);
//        }
//
//        saveHistory(user, action, testCase,null);
//        return caseVo;

        return null;
	}

    @Override
    public void loadNodeTree(TstCase vo, TstCase po) {
//        BeanUtilEx.copyProperties(vo, po);
//        vo.setEstimate(po.getEstimate());
//
//        List<TstCase> children = getChildren(po.getId());
//        for (TstCase childPo : children) {
//            TstCase childVo = new TstCase();
//            vo.getChildren().add(childVo);
//
//            loadNodeTree(childVo, childPo);
//        }
    }

    @Override
    public void createRoot(Integer projectId, TstUser user) {
//        TstCase root = new TstCase();
//        root.setName("测试用例");
//        root.setType(null);
//        root.setpId(null);
//        root.setLeaf(false);
//        root.setProjectId(projectId);
//
//        root.setCreateById(user.getId());
//        root.setCreateTime(new Date());
//
//        root.setOrdr(0);
//
//        saveOrUpdate(root);
//
//        TstCase testCase = new TstCase();
//        testCase.setName("新特性");
//        testCase.setType("functional");
//        testCase.setPriority("medium");
//        testCase.setpId(root.getId());
//        testCase.setProjectId(projectId);
//        testCase.setCreateById(user.getId());
//        testCase.setCreateTime(new Date());
//        testCase.setLeaf(false);
//        testCase.setOrdr(0);
//        saveOrUpdate(testCase);
//        saveHistory(user, Constant.CaseAct.create, testCase,null);
//
//        TstCase testCase2 = new TstCase();
//        testCase2.setName("新用例");
//        testCase2.setType("functional");
//        testCase2.setPriority("medium");
//        testCase2.setpId(testCase.getId());
//        testCase2.setProjectId(projectId);
//        testCase2.setCreateById(user.getId());
//        testCase2.setCreateTime(new Date());
//        testCase2.setLeaf(true);
//        testCase2.setOrdr(0);
//        saveOrUpdate(testCase2);
//        saveHistory(user, Constant.CaseAct.create, testCase2,null);
    }

    @Override
	public TstCase save(JSONObject json, TstUser user) {
//        TstCase testCaseVo = JSON.parseObject(JSON.toJSONString(json), TstCase.class);
//
//        Constant.CaseAct action;
//
//        TstCase testCasePo;
//        if (testCaseVo.getId() > 0) {
//            testCasePo = (TstCase)get(TstCase.class, testCaseVo.getId());
//            copyProperties(testCasePo, testCaseVo);
//
//            testCasePo.setUpdateById(user.getId());
//            testCasePo.setUpdateTime(new Date());
//
//            action = Constant.CaseAct.update;
//        } else {
//            testCasePo = new TstCase();
//            copyProperties(testCasePo, testCaseVo);
//            testCasePo.setId(null);
//            testCasePo.setLeaf(true);
//            testCasePo.setOrdr(getChildMaxOrderNumb(testCasePo.getpId()));
//
//            testCasePo.setCreateById(user.getId());
//            testCasePo.setCreateTime(new Date());
//
//            action = Constant.CaseAct.create;
//        }
//
//        testCasePo.setReviewResult(null);
//        saveOrUpdate(testCasePo);
//
//        saveHistory(user, action, testCasePo,null);

		return null;
	}

    @Override
	public TstCase saveField(JSONObject json, TstUser user) {
//		Long id = json.getLong("id");
//		String prop = json.getString("prop");
//		String value = json.getString("value");
//		String label = json.getString("label");
//
//		TstCase testCase = (TstCase) get(TstCase.class, id);
//
//		if ("name".equals(prop)) {
//			testCase.setName(value);
//		} else if ("objective".equals(prop)) {
//            testCase.setObjective(value);
//        } else if ("content".equals(prop)) {
//            testCase.setContent(value);
//        } else if ("priority".equals(prop)) {
//            testCase.setPriority(value);
//        } else if ("estimate".equals(prop)) {
//            testCase.setEstimate(Integer.valueOf(value));
//        } else if ("type".equals(prop)) {
//            testCase.setType(value);
//        } else if ("prop01".equals(prop)) {
//            testCase.setProp01(value);
//        } else if ("prop02".equals(prop)) {
//            testCase.setProp02(value);
//        } else if ("prop03".equals(prop)) {
//            testCase.setProp03(value);
//        } else if ("prop04".equals(prop)) {
//            testCase.setProp04(value);
//        } else if ("prop05".equals(prop)) {
//            testCase.setProp05(value);
//        } else if ("prop06".equals(prop)) {
//            testCase.setProp06(value);
//        } else if ("prop07".equals(prop)) {
//            testCase.setProp07(value);
//        } else if ("prop08".equals(prop)) {
//            testCase.setProp08(value);
//        } else if ("prop09".equals(prop)) {
//            testCase.setProp09(value);
//        } else if ("prop10".equals(prop)) {
//            testCase.setProp10(value);
//        } else if ("prop11".equals(prop)) {
//            testCase.setProp11(value);
//        } else if ("prop12".equals(prop)) {
//            testCase.setProp12(value);
//        } else if ("prop13".equals(prop)) {
//            testCase.setProp13(value);
//        } else if ("prop14".equals(prop)) {
//            testCase.setProp14(value);
//        } else if ("prop15".equals(prop)) {
//            testCase.setProp15(value);
//        } else if ("prop16".equals(prop)) {
//            testCase.setProp16(value);
//        } else if ("prop17".equals(prop)) {
//            testCase.setProp17(value);
//        } else if ("prop18".equals(prop)) {
//            testCase.setProp18(value);
//        } else if ("prop19".equals(prop)) {
//            testCase.setProp19(value);
//        } else if ("prop20".equals(prop)) {
//            testCase.setProp20(value);
//        }
//        testCase.setReviewResult(null);
//		saveOrUpdate(testCase);
//
//        saveHistory(user, Constant.CaseAct.update, testCase,label);
//
//		return testCase;

        return null;
	}

    @Override
    public void saveHistory(TstUser user, Constant.CaseAct act, TstCase testCase, String field) {
//	    String action = act.msg;
//
//        String msg = "用户" + StringUtil.highlightDict(user.getName()) + action;
//        if (StringUtils.isNotEmpty(field)) {
//            msg += " " + field;
//        } else {
////            msg += "信息";
//        }
//        TstCaseHistory his = new TstCaseHistory();
//        his.setTitle(msg);
//        his.setTstCaseId(testCase.getId());
//        saveOrUpdate(his);
    }

	@Override
	public TstCase delete(Integer id, TstUser user) {
//        TstCase testCase = (TstCase) get(TstCase.class, id);
//
//        getDao().querySql("{call remove_case_and_its_children(?)}", id);
//
//        getDao().flush();
//        updateParentIfNeededPers(testCase.getpId());
//        saveHistory(user, Constant.CaseAct.delete, testCase,null);
//
//        return testCase;

        return null;
	}

    @Override
    public void updateParentIfNeededPers(Integer pid) {
//        getDao().querySql("{call update_case_parent_if_needed(?)}", pid);
    }

    @Override
    public boolean cloneStepsAndChildrenPers(TstCase testCase, TstCase src) {
//	    boolean isParent = false;
//
//        for (TstCaseStep step : src.getSteps()) {
//            TstCaseStep step1 = new TstCaseStep(testCase.getId(), step.getOpt(), step.getExpect(), step.getOrdr());
//            saveOrUpdate(step1);
//            testCase.getSteps().add(step1);
//        }
//
//        List<TstCase> children = getChildren(src.getId());
//        for(TstCase child : children) {
//            TstCase clonedChild = new TstCase();
//            BeanUtilEx.copyProperties(clonedChild, child);
//            // 不能用以前的
//            clonedChild.setComments(new LinkedList());
//            clonedChild.setSteps(new LinkedList());
//            clonedChild.setHistories(new LinkedList());
//            clonedChild.setAttachments(new LinkedList());
//
//            clonedChild.setId(null);
//            clonedChild.setpId(testCase.getId());
//
//            saveOrUpdate(clonedChild);
//            cloneStepsAndChildrenPers(clonedChild, child);
//        }
//
//        return children.size() > 0;

        return true;
    }

    @Override
    public List<TstCase> getChildren(Integer caseId) {
//        DetachedCriteria dc = DetachedCriteria.forClass(TstCase.class);
//        dc.add(Restrictions.eq("pId", caseId));
//
//        dc.add(Restrictions.eq("deleted", Boolean.FALSE));
//        dc.add(Restrictions.eq("disabled", Boolean.FALSE));
//
//        dc.addOrder(Order.asc("pId"));
//        dc.addOrder(Order.asc("ordr"));
//
//        List<TstCase> children = findAllByCriteria(dc);
//        return children;

        return null;
    }

	private Integer getChildMaxOrderNumb(Long parentId) {
//		String hql = "select max(ordr) from TstCase where pId = " + parentId;
//		Integer maxOrder = (Integer) getByHQL(hql);
//
//		if (maxOrder == null) {
//			maxOrder = 0;
//		}
//
//		return maxOrder + 1;

        return 1;
	}

    @Override
    public List<TstCase> genVos(List<TstCase> pos) { return genVos(pos, false); }
    @Override
    public List<TstCase> genVos(List<TstCase> pos, boolean withSteps) { return genVos(pos, null,false); }
    @Override
    public List<TstCase> genVos(List<TstCase> pos, List<Integer> selectIds, boolean withSteps) {
        List<TstCase> vos = new LinkedList<TstCase>();

        for (TstCase po: pos) {
            TstCase vo = genVo(po, selectIds, withSteps);
            vos.add(vo);
        }
        return vos;
    }

    @Override
    public TstCase genVo(TstCase po) {
        return genVo(po, false);
    }

    @Override
    public TstCase genVo(TstCase po, boolean withSteps) { return genVo(po, null, withSteps);}

    @Override
    public TstCase genVo(TstCase po, List<Integer> selectIds, boolean withSteps) {
//        TstCase vo = new TstCase();
//
//        BeanUtilEx.copyProperties(vo, po);
//        vo.setEstimate(po.getEstimate());
//
//        if (selectIds != null && selectIds.contains(po.getId())) {
//            vo.setChecked(true);
//        }
//
//        vo.setSteps(new LinkedList<TstCaseStep>());
//        vo.setComments(new LinkedList<TstCaseComments>());
//        vo.setHistories(new LinkedList<TstCaseHistory>());
//        vo.setAttachments(new LinkedList<TstCaseAttachment>());
//
//        if (withSteps) {
//            List<TstCaseStep> steps = po.getSteps();
//            for (TstCaseStep step : steps) {
//                TstCaseStep stepVo = new TstCaseStep(
//                        step.getId(), step.getOpt(), step.getExpect(), step.getOrdr(), step.getTstCaseId());
//
//                vo.getSteps().add(stepVo);
//            }
//
//            List<TstCaseComments> comments = po.getComments();
//            Iterator<TstCaseComments> iterator  = comments.iterator();
//            while (iterator.hasNext()) {
//                TstCaseComments comment = iterator.next();
//                TstCaseComments commentVo = caseCommentsService.genVo(comment);
//                vo.getComments().add(commentVo);
//            }
//
//            // 用例历史
//            List<TstCaseHistory> histories = findHistories(po.getId());
//            for (TstCaseHistory his : histories) {
//                TstCaseHistory historyVo = new TstCaseHistory(
//                        his.getId(), his.getTitle(), his.getDescr(), his.getTstCaseId(), his.getCreateTime());
//
//                vo.getHistories().add(historyVo);
//            }
//
//            List<TstCaseAttachment> attachments = po.getAttachments();
//            Iterator<TstCaseAttachment> iteratorAttach  = attachments.iterator();
//            while (iteratorAttach.hasNext()) {
//                TstCaseAttachment attachment = iteratorAttach.next();
//                TstCaseAttachment attachVo = caseAttachmentService.genVo(attachment);
//                vo.getAttachments().add(attachVo);
//            }
//        }
//
//        return vo;

        return null;
    }

    @Override
    public List<TstCaseHistory> findHistories(Integer testCaseId) {
//        DetachedCriteria dc = DetachedCriteria.forClass(TstCaseHistory.class);
//        dc.add(Restrictions.eq("testCaseId", testCaseId));
//
//        dc.add(Restrictions.eq("deleted", Boolean.FALSE));
//        dc.add(Restrictions.eq("disabled", Boolean.FALSE));
//
//        dc.addOrder(Order.desc("createTime"));
//
//        List<TstCaseHistory> ls = findAllByCriteria(dc);
//        return ls;

        return null;
    }

    @Override
    public void copyProperties(TstCase testCasePo, TstCase testCaseVo) {
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
    public TstCase changeContentTypePers(Integer id, String contentType) {
//        TstCase testCase = (TstCase)get(TstCase.class, id);
//        testCase.setContentType(contentType);
//        testCase.setReviewResult(null);
//        saveOrUpdate(testCase);
//
//        return testCase;

        return null;
    }

    @Override
    public TstCase reviewPassPers(Integer id, Boolean pass) {
//        TstCase testCase = (TstCase)get(TstCase.class, id);
//        testCase.setReviewResult(pass);
//        saveOrUpdate(testCase);
//
//        return testCase;

        return null;
    }

}

