package com.ngtesting.platform.service.impl;

import com.alibaba.fastjson.JSON;
import com.alibaba.fastjson.JSONObject;
import com.ngtesting.platform.entity.AiTestTask;
import com.ngtesting.platform.service.AiTestTaskService;
import com.ngtesting.platform.util.BeanUtilEx;
import com.ngtesting.platform.vo.*;
import org.hibernate.criterion.DetachedCriteria;
import org.hibernate.criterion.Order;
import org.hibernate.criterion.Restrictions;
import org.springframework.stereotype.Service;

import java.util.Date;
import java.util.LinkedList;
import java.util.List;

@Service
public class AiTestTaskServiceImpl extends BaseServiceImpl implements AiTestTaskService {

	@Override
	public List<AiTestTask> query(Long projectId) {
        DetachedCriteria dc = DetachedCriteria.forClass(AiTestTask.class);


        dc.add(Restrictions.eq("projectId", projectId));

        dc.add(Restrictions.eq("deleted", Boolean.FALSE));
        dc.add(Restrictions.eq("disabled", Boolean.FALSE));

		dc.addOrder(Order.asc("pId"));
        dc.addOrder(Order.asc("ordr"));

        List<AiTestTask> ls = findAllByCriteria(dc);

        return ls;
	}

    @Override
	public AiTestTaskVo getById(Long caseId) {
		AiTestTask po = (AiTestTask) get(AiTestTask.class, caseId);
		AiTestTaskVo vo = genVo(po);

		return vo;
	}

    @Override
    public AiTestTask renamePers(JSONObject json, UserVo user) {
        Long id = json.getLong("id");
        String name = json.getString("name");
        Long pId = json.getLong("pId");
        Long projectId = json.getLong("projectId");

        return renamePers(id, name, pId, projectId, user);
    }

	@Override
	public AiTestTask renamePers(Long id, String name, Long pId, Long projectId, UserVo user) {
        AiTestTask testCasePo = new AiTestTask();
        String action = "";
        if (id != null && id > 0) {
            testCasePo = (AiTestTask)get(AiTestTask.class, id);

            testCasePo.setUpdateTime(new Date());
            action = "rename";
        } else {
            testCasePo.setLeaf(true);
            testCasePo.setId(null);
            testCasePo.setpId(pId);
            testCasePo.setOrdr(getChildMaxOrderNumb(testCasePo.getpId()));

            testCasePo.setCreateTime(new Date());
            action = "create";
        }
        testCasePo.setName(name);
        testCasePo.setProjectId(projectId);

        saveOrUpdate(testCasePo);

        return testCasePo;
	}

	@Override
	public AiTestTaskVo movePers(JSONObject json, UserVo user) {
        Long srcId = json.getLong("srcId");
        Long targetId = json.getLong("targetId");
        String moveType = json.getString("moveType");
        Boolean isCopy = json.getBoolean("isCopy");

        AiTestTask src = (AiTestTask) get(AiTestTask.class, srcId);;
        AiTestTask target = (AiTestTask) get(AiTestTask.class, targetId);

        AiTestTask testCase;
        String action;
        if (isCopy) {
            testCase = new AiTestTask();
            BeanUtilEx.copyProperties(testCase, src);

            testCase.setId(null);
            action = "copy";
        } else {
            testCase = src;
            action = "move";
        }

        if ("inner".equals(moveType)) {
            testCase.setpId(target.getId());
        } else if ("prev".equals(moveType)) {
            String hql = "update AiTestTask c set c.ordr = c.ordr+1 where c.ordr >= ? and c.pId=? and id!=?";
            getDao().queryHql(hql, target.getOrdr(), target.getpId(), testCase.getId());

            testCase.setpId(target.getpId());
            testCase.setOrdr(target.getOrdr());
        } else if ("next".equals(moveType)) {
            String hql = "update AiTestTask c set c.ordr = c.ordr+1 where c.ordr > ? and c.pId=? and id!=?";
            getDao().queryHql(hql, target.getOrdr(), target.getpId(), testCase.getId());

            testCase.setpId(target.getpId());
            testCase.setOrdr(target.getOrdr() + 1);
        }

        saveOrUpdate(testCase);
        boolean isParent = false;
        if (isCopy) {
            isParent = cloneChildrenPers(testCase, src);
        }

        AiTestTaskVo caseVo = new AiTestTaskVo();
        if (isCopy && isParent) {
            loadNodeTree(caseVo, testCase);
        } else {
            caseVo = genVo(testCase);
        }

        return caseVo;
	}

    @Override
    public void loadNodeTree(AiTestTaskVo vo, AiTestTask po) {
        BeanUtilEx.copyProperties(vo, po);

        List<AiTestTask> children = getChildren(po.getId());
        for (AiTestTask childPo : children) {
            AiTestTaskVo childVo = new AiTestTaskVo();
            vo.getChildren().add(childVo);

            loadNodeTree(childVo, childPo);
        }
    }

    @Override
	public AiTestTask save(JSONObject json, UserVo user) {
        AiTestTaskVo testCaseVo = JSON.parseObject(JSON.toJSONString(json), AiTestTaskVo.class);

        String action = "";

        AiTestTask testCasePo;
        if (testCaseVo.getId() > 0) {
            testCasePo = (AiTestTask)get(AiTestTask.class, testCaseVo.getId());
            copyProperties(testCasePo, testCaseVo);

            testCasePo.setUpdateTime(new Date());

            action = "update";
        } else {
            testCasePo = new AiTestTask();
            copyProperties(testCasePo, testCaseVo);
            testCasePo.setId(null);
            testCasePo.setLeaf(true);
            testCasePo.setOrdr(getChildMaxOrderNumb(testCasePo.getpId()));

            testCasePo.setCreateById(user.getId());
            testCasePo.setCreateTime(new Date());

            action = "create";
        }

        saveOrUpdate(testCasePo);

		return testCasePo;
	}

	@Override
	public AiTestTask delete(Long id, UserVo user) {
        AiTestTask testCase = (AiTestTask) get(AiTestTask.class, id);

        getDao().querySql("{call delete_case_and_its_children(?)}", id);

        return testCase;
	}

    @Override
    public void updateParentIfNeededPers(Long pid) {
        getDao().querySql("{call update_case_parent_if_needed(?)}", pid);
    }

    @Override
    public boolean cloneChildrenPers(AiTestTask testCase, AiTestTask src) {
	    boolean isParent = false;

        List<AiTestTask> children = getChildren(src.getId());
        for(AiTestTask child : children) {
            AiTestTask clonedChild = new AiTestTask();
            BeanUtilEx.copyProperties(clonedChild, child);

            clonedChild.setId(null);
            clonedChild.setpId(testCase.getId());

            saveOrUpdate(clonedChild);
            cloneChildrenPers(clonedChild, child);
        }

        return children.size() > 0;
    }

    @Override
    public List<AiTestTask> getChildren(Long caseId) {
        DetachedCriteria dc = DetachedCriteria.forClass(AiTestTask.class);
        dc.add(Restrictions.eq("pId", caseId));

        dc.add(Restrictions.eq("deleted", Boolean.FALSE));
        dc.add(Restrictions.eq("disabled", Boolean.FALSE));

        dc.addOrder(Order.asc("pId"));
        dc.addOrder(Order.asc("ordr"));

        List<AiTestTask> children = findAllByCriteria(dc);
        return children;
    }

	private Integer getChildMaxOrderNumb(Long parentId) {
		String hql = "select max(ordr) from AiTestTask where pId = " + parentId;
		Integer maxOrder = (Integer) getByHQL(hql);

		if (maxOrder == null) {
			maxOrder = 0;
		}

		return maxOrder + 1;
	}

     @Override
    public List<AiTestTaskVo> genVos(List<AiTestTask> pos) {
        List<AiTestTaskVo> vos = new LinkedList<AiTestTaskVo>();

        for (AiTestTask po: pos) {
            AiTestTaskVo vo = genVo(po);
            vos.add(vo);
        }
        return vos;
    }

    @Override
    public AiTestTaskVo genVo(AiTestTask po) {
        AiTestTaskVo vo = new AiTestTaskVo();

        BeanUtilEx.copyProperties(vo, po);
        vo.setTestsetName(po.getTestset()==null?"":po.getTestset().getName());

        return vo;
    }
    @Override
    public void copyProperties(AiTestTask testCasePo, AiTestTaskVo testCaseVo) {
        BeanUtilEx.copyProperties(testCasePo, testCaseVo);
    }

}

