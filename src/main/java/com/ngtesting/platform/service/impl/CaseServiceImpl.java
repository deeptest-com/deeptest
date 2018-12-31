package com.ngtesting.platform.service.impl;

import com.alibaba.fastjson.JSON;
import com.alibaba.fastjson.JSONObject;
import com.ngtesting.platform.config.Constant;
import com.ngtesting.platform.dao.CaseDao;
import com.ngtesting.platform.dao.CaseStepDao;
import com.ngtesting.platform.dao.TestSuiteDao;
import com.ngtesting.platform.dao.TestTaskDao;
import com.ngtesting.platform.model.TstCase;
import com.ngtesting.platform.model.TstCaseComments;
import com.ngtesting.platform.model.TstCaseStep;
import com.ngtesting.platform.model.TstUser;
import com.ngtesting.platform.service.intf.CaseCommentsService;
import com.ngtesting.platform.service.intf.CaseHistoryService;
import com.ngtesting.platform.service.intf.CaseService;
import com.ngtesting.platform.utils.BeanUtilEx;
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
    CaseHistoryService caseHistoryService;

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
        Integer pId = json.getInteger("pId");

        Integer projectId = user.getDefaultPrjId();

        return rename(id, name, pId, projectId, user);
    }

	@Override
    @Transactional
	public TstCase rename(Integer id, String name, Integer pId, Integer projectId, TstUser user) {
        TstCase po = new TstCase();
        Constant.EntityAct action;

        boolean isNew;
        if (id != null && id > 0) {
            isNew = false;
            action = Constant.EntityAct.rename;
            po = caseDao.get(id, projectId);
            if(po == null) {
                return null;
            }

            po.setUpdateById(user.getId());
        } else {
            isNew = true;
            action = Constant.EntityAct.create;

            po.setLeaf(true);
            po.setId(null);
            po.setpId(pId);
            po.setOrdr(getChildMaxOrderNumb(po.getpId()));

            po.setProjectId(projectId);
            po.setCreateById(user.getId());
            action = Constant.EntityAct.create;
        }
        po.setName(name);
        po.setReviewResult(null);

        if (isNew) {
            caseDao.renameNew(po);
            caseDao.setDefaultVal(po.getId(), user.getDefaultOrgId());
            caseDao.updateParentIfNeeded(po.getpId());
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
        Constant.EntityAct action;

        if (isCopy) {
            action = Constant.EntityAct.copy;

            testCase = new TstCase();
            BeanUtilEx.copyProperties(src, testCase);
            testCase.setId(null);

            testCase.setCreateById(user.getId());
        } else {
            action = Constant.EntityAct.move;
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

        if (!isCopy) {
            caseDao.updateParentIfNeeded(srcParentId);
        }
        if ("inner".equals(moveType)) {
            caseDao.updateParentIfNeeded(targetId);
        }

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

        TstCase testCaseVo = JSON.parseObject(JSON.toJSONString(json), TstCase.class);

        testCaseVo.setUpdateById(user.getId());
        Integer count = caseDao.update(testCaseVo, genExtPropList(), projectId);
        if (count == 0) {
            return null;
        }

        caseHistoryService.saveHistory(user, Constant.EntityAct.update, testCaseVo.getId(),null);

        TstCase ret = caseDao.getDetail(testCaseVo.getId(), projectId);
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
        caseDao.updateParentIfNeeded(testCase.getpId());

        caseHistoryService.saveHistory(user, Constant.EntityAct.delete, testCase.getId(),null);

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
        caseHistoryService.saveHistory(user, Constant.EntityAct.update, testCase.getId(), "内容类型");
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

        Integer id = json.getInteger("id");
        String code = json.getString("code");
        String value = json.getString("value");
        String label = json.getString("label");

        Integer count = caseDao.updateProp(id, code, value, projectId, user.getId());
        if (count == 0) {
            return null;
        }

        TstCase testCase = caseDao.getDetail(id, projectId);
        caseHistoryService.saveHistory(user, Constant.EntityAct.update, testCase.getId(),label);

        return testCase;
    }

    @Override
    @Transactional
    public void createSample(Integer projectId, TstUser user) {
        TstCase root = new TstCase();
        root.setName("测试用例");
        root.setLeaf(false);
        root.setProjectId(projectId);
        root.setCreateById(user.getId());
        root.setCreateTime(new Date());
        root.setOrdr(0);

        caseDao.create(root);
        caseDao.setDefaultVal(root.getId(), user.getDefaultOrgId());

        TstCase testCase = new TstCase();
        testCase.setName("新特性");
        testCase.setpId(root.getId());
        testCase.setProjectId(projectId);
        testCase.setCreateById(user.getId());
        testCase.setCreateTime(new Date());
        testCase.setLeaf(false);
        testCase.setOrdr(0);
        caseDao.create(testCase);
        caseDao.setDefaultVal(testCase.getId(), user.getDefaultOrgId());
        caseHistoryService.saveHistory(user, Constant.EntityAct.create, testCase.getId(),null);

        TstCase testCase2 = new TstCase();
        testCase2.setName("新用例");
        testCase2.setpId(testCase.getId());
        testCase2.setProjectId(projectId);
        testCase2.setCreateById(user.getId());
        testCase2.setCreateTime(new Date());
        testCase2.setLeaf(true);
        testCase2.setOrdr(0);
        caseDao.create(testCase2);
        caseDao.setDefaultVal(testCase2.getpId(), user.getDefaultOrgId());

        caseHistoryService.saveHistory(user, Constant.EntityAct.create, testCase2.getId(),null);
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