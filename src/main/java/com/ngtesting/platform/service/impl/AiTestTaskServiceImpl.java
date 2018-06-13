package com.ngtesting.platform.service.impl;

import com.alibaba.fastjson.JSON;
import com.alibaba.fastjson.JSONObject;
import com.ngtesting.platform.config.Constant;
import com.ngtesting.platform.config.PropertyConfig;
import com.ngtesting.platform.entity.AiTestTask;
import com.ngtesting.platform.service.AiTestTaskService;
import com.ngtesting.platform.util.BeanUtilEx;
import com.ngtesting.platform.util.DateUtils;
import com.ngtesting.platform.util.FileUtils;
import com.ngtesting.platform.vo.AiRunMlf;
import com.ngtesting.platform.vo.AiTestTaskVo;
import com.ngtesting.platform.vo.UserVo;
import org.hibernate.criterion.DetachedCriteria;
import org.hibernate.criterion.Order;
import org.hibernate.criterion.Restrictions;
import org.springframework.stereotype.Service;
import org.zeroturnaround.zip.ZipUtil;

import java.io.File;
import java.util.Date;
import java.util.LinkedList;
import java.util.List;
import java.util.UUID;

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
	public AiTestTaskVo getById(Long taskId) {
		AiTestTask po = (AiTestTask) get(AiTestTask.class, taskId);
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
        AiTestTask po = new AiTestTask();
        String action = "";
        if (id != null && id > 0) {
            po = (AiTestTask)get(AiTestTask.class, id);

            po.setUpdateTime(new Date());
            action = "rename";
        } else {
            po.setLeaf(true);
            po.setId(null);
            po.setpId(pId);
            po.setOrdr(getChildMaxOrderNumb(po.getpId()));

            po.setCreateTime(new Date());
            action = "create";
        }
        po.setName(name);
        po.setProjectId(projectId);

        saveOrUpdate(po);

        return po;
	}

	@Override
	public AiTestTaskVo movePers(JSONObject json, UserVo user) {
        Long srcId = json.getLong("srcId");
        Long targetId = json.getLong("targetId");
        String moveType = json.getString("moveType");
        Boolean isCopy = json.getBoolean("isCopy");

        AiTestTask src = (AiTestTask) get(AiTestTask.class, srcId);;
        AiTestTask target = (AiTestTask) get(AiTestTask.class, targetId);

        AiTestTask task;
        String action;
        if (isCopy) {
            task = new AiTestTask();
            BeanUtilEx.copyProperties(task, src);

            task.setId(null);
            action = "copy";
        } else {
            task = src;
            action = "move";
        }

        if ("inner".equals(moveType)) {
            task.setpId(target.getId());
        } else if ("prev".equals(moveType)) {
            String hql = "update AiTestTask c set c.ordr = c.ordr+1 where c.ordr >= ? and c.pId=? and id!=?";
            getDao().queryHql(hql, target.getOrdr(), target.getpId(), task.getId());

            task.setpId(target.getpId());
            task.setOrdr(target.getOrdr());
        } else if ("next".equals(moveType)) {
            String hql = "update AiTestTask c set c.ordr = c.ordr+1 where c.ordr > ? and c.pId=? and id!=?";
            getDao().queryHql(hql, target.getOrdr(), target.getpId(), task.getId());

            task.setpId(target.getpId());
            task.setOrdr(target.getOrdr() + 1);
        }

        saveOrUpdate(task);
        boolean isParent = false;
        if (isCopy) {
            isParent = cloneChildrenPers(task, src);
        }

        AiTestTaskVo vo = new AiTestTaskVo();
        if (isCopy && isParent) {
            loadNodeTree(vo, task);
        } else {
            vo = genVo(task);
        }

        return vo;
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
        AiTestTaskVo testTaskVo = JSON.parseObject(JSON.toJSONString(json), AiTestTaskVo.class);

        String action = "";
        AiTestTask testTaskPo;
        if (testTaskVo.getId() > 0) {
            testTaskPo = (AiTestTask)get(AiTestTask.class, testTaskVo.getId());
            copyProperties(testTaskPo, testTaskVo);

            testTaskPo.setUpdateTime(new Date());

            action = "update";
        } else {
            testTaskPo = new AiTestTask();
            copyProperties(testTaskPo, testTaskVo);
            testTaskPo.setId(null);
            testTaskPo.setLeaf(true);
            testTaskPo.setOrdr(getChildMaxOrderNumb(testTaskPo.getpId()));

            testTaskPo.setCreateById(user.getId());
            testTaskPo.setCreateTime(new Date());

            action = "create";
        }

        // 解压文件
        String zipPath = Constant.WORK_DIR + testTaskVo.getTestsetPath();
        String destDir = PropertyConfig.getConfig("res.upload.dir");
        String dateDist = DateUtils.GetDateNoSeparator();
        destDir = destDir + dateDist + "/" + UUID.randomUUID().toString();

        File file = new File(destDir);
        System.out.println(destDir);
        if (!file.exists()) {
            boolean ret = file.mkdirs();
        }
        if (zipPath.endsWith(".zip")) {
            ZipUtil.unpack(new File(zipPath), file);
        }

        List<AiRunMlf> mlfs = FileUtils.ListMlf(destDir, testTaskVo.getTestType());
        testTaskPo.setMlfs(JSON.toJSONString(mlfs));
        saveOrUpdate(testTaskPo);

		return testTaskPo;
	}

	@Override
	public AiTestTask delete(Long id, UserVo user) {
        AiTestTask task = (AiTestTask) get(AiTestTask.class, id);

        getDao().querySql("{call remove_aitask_and_its_children(?)}", id);

        return task;
	}

    @Override
    public void updateParentIfNeededPers(Long pid) {
        getDao().querySql("{call update_aitask_parent_if_needed(?)}", pid);
    }

    @Override
    public boolean cloneChildrenPers(AiTestTask task, AiTestTask src) {
	    boolean isParent = false;

        List<AiTestTask> children = getChildren(src.getId());
        for(AiTestTask child : children) {
            AiTestTask clonedChild = new AiTestTask();
            BeanUtilEx.copyProperties(clonedChild, child);

            clonedChild.setId(null);
            clonedChild.setpId(task.getId());

            saveOrUpdate(clonedChild);
            cloneChildrenPers(clonedChild, child);
        }

        return children.size() > 0;
    }

    @Override
    public List<AiTestTask> getChildren(Long taskId) {
        DetachedCriteria dc = DetachedCriteria.forClass(AiTestTask.class);
        dc.add(Restrictions.eq("pId", taskId));

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
    public void copyProperties(AiTestTask po, AiTestTaskVo vo) {
        BeanUtilEx.copyProperties(po, vo);
    }

}

