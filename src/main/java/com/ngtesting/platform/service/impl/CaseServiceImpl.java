package com.ngtesting.platform.service.impl;

import com.alibaba.fastjson.JSON;
import com.alibaba.fastjson.JSONObject;
import com.ngtesting.platform.dao.*;
import com.ngtesting.platform.model.*;
import com.ngtesting.platform.service.intf.*;
import com.ngtesting.platform.utils.BeanUtilEx;
import com.ngtesting.platform.utils.CustomFieldUtil;
import com.ngtesting.platform.utils.MsgUtil;
import com.ngtesting.platform.utils.StringUtil;
import org.springframework.beans.factory.annotation.Autowired;
import org.springframework.stereotype.Service;
import org.springframework.transaction.annotation.Transactional;

import java.util.Date;
import java.util.LinkedList;
import java.util.List;

@Service
public class CaseServiceImpl extends BaseServiceImpl implements CaseService {

    @Autowired
    CaseDao caseDao;
    @Autowired
    CaseStepDao caseStepDao;

    @Autowired
    TestSuiteDao testSuiteDao;
    @Autowired
    TestTaskDao testTaskDao;
    @Autowired
    CaseCommentsService caseCommentsService;

    public static List<String> ExtPropList;

    @Autowired
    CasePriorityService casePriorityService;

    @Autowired
    CaseTypeService caseTypeService;

    @Autowired
    CaseHistoryService caseHistoryService;
    @Autowired
    CustomFieldDao customFieldDao;
    @Autowired
    AuthDao authDao;

	@Override
	public List<TstCase> query(Integer projectId) {
        List<TstCase> ls = caseDao.query(projectId);

        return ls;
	}

    @Override
    public List<TstCase> queryForSuiteSelection(Integer id, Integer suiteId) {
        List<TstCase> pos = caseDao.query(id);
        List<Integer> selectIds = testSuiteDao.listCaseIds(suiteId);
        genVos(pos, selectIds);

        return pos;
    }

    @Override
    public List<TstCase> queryForTaskSelection(Integer id, Integer taskId) {
        List<TstCase> pos = caseDao.query(id);
        List<Integer> selectIds = testTaskDao.listCaseIds(taskId);
        genVos(pos, selectIds);

        return pos;
    }

    @Override
	public TstCase getDetail(Integer caseId, Integer prjId) {
		TstCase po = caseDao.getDetail(caseId, prjId);

		return po;
	}

    @Override
    @Transactional
    public TstCase rename(JSONObject json, TstUser user) {
        Integer id = json.getInteger("id");
        String name = json.getString("name");
        Boolean isParent = json.getBoolean("isParent");
        Integer pId = json.getInteger("pId");

        Integer projectId = user.getDefaultPrjId();

        return rename(id, name, isParent, pId, projectId, user);
    }

	@Override
    @Transactional
	public TstCase rename(Integer id, String name, Boolean isParent, Integer pId, Integer projectId, TstUser user) {
        TstCase po = new TstCase();
        MsgUtil.MsgAction action;

        boolean isNew;
        if (id != null && id > 0) {
            isNew = false;
            action = MsgUtil.MsgAction.rename;
            po = caseDao.get(id, projectId);
            if(po == null) {
                return null;
            }

            po.setUpdateById(user.getId());
        } else {
            isNew = true;
            action = MsgUtil.MsgAction.create;

            po.setIsParent(isParent);
            po.setId(null);
            po.setpId(pId);
            po.setOrdr(getChildMaxOrderNumb(po.getpId()));

            po.setProjectId(projectId);
            po.setCreateById(user.getId());
        }
        po.setName(name);
        po.setReviewResult(null);

        if (isNew) {
            caseDao.renameNew(po);
            caseDao.setDefaultVal(po.getId(), user.getDefaultOrgId());
//            caseDao.updateParentIfNeeded(po.getpId());
        } else {
            caseDao.renameUpdate(po);
        }

        caseHistoryService.saveHistory(user, action, po.getId(),null);

        TstCase ret = caseDao.getDetail(po.getId(), projectId);
        return ret;
	}

	@Override
    @Transactional
	public TstCase move(JSONObject json, TstUser user) {
        Integer projectId = user.getDefaultPrjId();

        Integer srcId = json.getInteger("srcId");

        Integer targetId = json.getInteger("targetId");
        String moveType = json.getString("moveType");
        Boolean isCopy = json.getBoolean("isCopy");

        TstCase src = caseDao.get(srcId, projectId);
        TstCase target = caseDao.get(targetId, projectId);

        if (src == null || target == null) {
            return null;
        }

        Integer srcParentId = src.getpId();

        TstCase testCase;
        MsgUtil.MsgAction action;

        if (isCopy) {
            action = MsgUtil.MsgAction.copy;

            testCase = new TstCase();
            BeanUtilEx.copyProperties(src, testCase);
            testCase.setId(null);

            testCase.setCreateById(user.getId());
        } else {
            action = MsgUtil.MsgAction.move;
            testCase = src;
            testCase.setUpdateById(user.getId());
        }

        if ("inner".equals(moveType)) {
            testCase.setpId(targetId);
        } else if ("prev".equals(moveType)) {
            caseDao.addOrderForTargetAndNextCases(testCase.getId(), target.getOrdr(), target.getpId());

            testCase.setpId(target.getpId());
            testCase.setOrdr(target.getOrdr());
        } else if ("next".equals(moveType)) {
            caseDao.addOrderForNextCases(testCase.getId(), target.getOrdr(), target.getpId());

            testCase.setpId(target.getpId());
            testCase.setOrdr(target.getOrdr() + 1);
        }

        boolean isParent = false;
        if (isCopy) {
            caseDao.moveCopy(testCase);
            isParent = cloneStepsAndChildrenPers(testCase, src);
        } else {
            caseDao.moveUpdate(testCase);
        }

//        if (!isCopy) {
//            caseDao.updateParentIfNeeded(srcParentId);
//        }
//        if ("inner".equals(moveType)) {
//            caseDao.updateParentIfNeeded(targetId);
//        }

        caseHistoryService.saveHistory(user, action, testCase.getId(),null);

        TstCase ret = caseDao.getDetail(testCase.getId(), projectId);
        if (isCopy && isParent) {
            loadNodeTree(ret);
        }

        return ret;
	}

    @Override
    @Transactional
	public TstCase update(JSONObject json, TstUser user) {
        Integer projectId = user.getDefaultPrjId();

        TstCase vo = JSON.parseObject(JSON.toJSONString(json), TstCase.class);

        json.put("updateById", user.getId());

        List<CustomField> fields = customFieldDao.listForCase(user.getDefaultOrgId());
        JSONObject jsonb = new JSONObject();
        List<String> props = genExtPropList();
        for (CustomField field : fields) {
            jsonb.put(field.getColCode(), json.get(field.getColCode()));
        }

        Integer count = caseDao.update(vo, jsonb.toJSONString(), projectId);
        if (count == 0) {
            return null;
        }

        caseHistoryService.saveHistory(user, MsgUtil.MsgAction.update, json.getInteger("id"),null);

        TstCase ret = caseDao.getDetail(json.getInteger("id"), projectId);
		return ret;
	}

	@Override
    @Transactional
	public Integer delete(Integer id, TstUser user) {
        Integer projectId = user.getDefaultPrjId();

        Integer count = caseDao.delete(id, projectId);
        if (count == 0) {
            return count;
        }

        TstCase testCase = caseDao.get(id, null);
//        caseDao.updateParentIfNeeded(testCase.getpId());

        caseHistoryService.saveHistory(user, MsgUtil.MsgAction.delete, testCase.getId(),null);

        return count;
	}

    @Override
    @Transactional
    public TstCase changeContentType(Integer id, String contentType, TstUser user) {
        Integer projectId = user.getDefaultPrjId();

        Integer count = caseDao.changeContentType(id, contentType, projectId, user.getId());
        if (count == 0) {
            return null;
        }

        TstCase testCase = caseDao.getDetail(id, projectId);
        caseHistoryService.saveHistory(user, MsgUtil.MsgAction.update, testCase.getId(), "内容类型");
        return testCase;
    }

    @Override
    @Transactional
    public TstCase reviewResult(Integer id, Boolean result, Integer nextId, TstUser user) {
        Integer projectId = user.getDefaultPrjId();

        Integer count = caseDao.reviewResult(id, result, projectId, user.getId());
        if (count == 0) {
            return null;
        }

        TstCaseComments vo = new TstCaseComments(id, result?"评审通过": "评审失败");
        caseCommentsService.save(vo, user);

        TstCase testCase;
        if (nextId != null) {
            testCase = caseDao.getDetail(nextId, projectId);
        } else {
            testCase = caseDao.getDetail(id, projectId);
        }
        return testCase;
    }

    @Override
    @Transactional
    public TstCase saveField(JSONObject json, TstUser user) {
        Integer projectId = user.getDefaultPrjId();
        Integer caseProjectId = json.getInteger("caseProjectId");

        if (caseProjectId != null && !authDao.userNotInProject(user.getId(), projectId)) {
            projectId = caseProjectId;
        }

        Integer id = json.getInteger("id");
        String code = json.getString("code");
        String label = json.getString("label");
        Boolean buildIn = json.getBoolean("buildIn");
        String type = json.getString("type");

        Object value = CustomFieldUtil.GetFieldVal(type, json);

        Integer count;
        if (buildIn) {
            count = caseDao.updateProp(id, code, value, projectId);
        } else {
            count = caseDao.updateExtProp(id, code, value, projectId);
        }

        if (count == 0) {
            return null;
        }

        TstCase testCase = caseDao.getDetail(id, projectId);
        caseHistoryService.saveHistory(user, MsgUtil.MsgAction.update, testCase.getId(),label);

        return testCase;
    }

    @Override
    @Transactional
    public void createSample(Integer projectId, TstUser user) {
        TstCase root = new TstCase("测试用例", null, projectId, null, null, user.getId(), true, 1);
        caseDao.createSample(root);
        caseDao.setDefaultVal(root.getId(), user.getDefaultOrgId());

        TstCase testCase = new TstCase("新特性", root.getId(), projectId, null, null, user.getId(), true, 1);
        caseDao.createSample(testCase);

        TstCaseType caseType = caseTypeService.getDefault(user.getDefaultOrgId());
        TstCasePriority casePriority = casePriorityService.getDefault(user.getDefaultOrgId());

        TstCase testCase2 = new TstCase("新用例", testCase.getId(), projectId, caseType.getId(), casePriority.getId(),
                user.getId(), false, 1);
        caseDao.createSample(testCase2);
        caseDao.setDefaultVal(testCase2.getpId(), user.getDefaultOrgId());

        caseHistoryService.saveHistory(user, MsgUtil.MsgAction.create, testCase2.getId(),null);

        TstCaseStep step = new TstCaseStep("操作步骤1", "期待结果1", 1, testCase2.getId());
        caseStepDao.save(step);
        step = new TstCaseStep("操作步骤2", "期待结果2", 2, testCase2.getId());
        caseStepDao.save(step);
        step = new TstCaseStep("操作步骤3", "期待结果3", 3, testCase2.getId());
        caseStepDao.save(step);
    }

    @Override
    public void loadNodeTree(TstCase po) {
        List<TstCase> children = caseDao.getChildren(po.getId());
        for (TstCase childPo : children) {
            po.getChildren().add(childPo);

            loadNodeTree(childPo);
        }
    }

    @Override
    @Transactional
    public boolean cloneStepsAndChildrenPers(TstCase testCase, TstCase src) {
        List<TstCaseStep> steps = caseStepDao.query(src.getId());

        for (TstCaseStep step : steps) {
            TstCaseStep step1 = new TstCaseStep();
            BeanUtilEx.copyProperties(step, step1);

            step1.setId(null);
            step1.setCaseId(testCase.getId());
            caseStepDao.save(step1);
        }

        List<TstCase> children = caseDao.getChildren(src.getId());
        for(TstCase child : children) {
            TstCase clonedChild = new TstCase();
            BeanUtilEx.copyProperties(child, clonedChild);

            clonedChild.setId(null);
            clonedChild.setpId(testCase.getId());

            testCase.setCreateTime(new Date());
            testCase.setUpdateTime(null);

            caseDao.moveCopy(clonedChild);
            cloneStepsAndChildrenPers(clonedChild, child);
        }

        return children.size() > 0;
    }

    @Override
    public Integer getChildMaxOrderNumb(Integer parentId) {
		Integer maxOrder = caseDao.getChildMaxOrderNumb(parentId);

		if (maxOrder == null) {
			maxOrder = 0;
		}

		return maxOrder + 1;
	}

    @Override
    public void genVos(List<TstCase> pos, List<Integer> selectIds) {
        for (TstCase po: pos) {
            genVo(po, selectIds);
        }
    }

    @Override
    public void genVo(TstCase po, List<Integer> selectIds) {
        if (selectIds != null && selectIds.contains(po.getId())) {
            po.setChecked(true);
        }
    }

    @Override
    public List<String> genExtPropList() {
        if (CaseServiceImpl.ExtPropList == null) {
            CaseServiceImpl.ExtPropList = new LinkedList<>();
            for (Integer i = 1; i <= 20; i++) {
                String str = StringUtil.formatString(i, 2);
                CaseServiceImpl.ExtPropList.add(str);
            }
        }

        return CaseServiceImpl.ExtPropList;
    }

}